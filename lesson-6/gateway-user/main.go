package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/metadata"

	"github.com/azomio/courses/lesson4/pkg/grpc/user"
	"github.com/azomio/courses/lesson4/pkg/jwt"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

var UserCli user.UserClient

const ServiceName = "gateway-user"

func main() {
	f, err := os.Create(
		fmt.Sprintf("logs/%s.log", ServiceName),
	)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)
	log.SetFormatter(&log.JSONFormatter{})

	consulAddr := flag.String("consul_addr", "localhost:8600", "Consul address")
	flag.Parse()

	if err := loadConfig(*consulAddr); err != nil {
		log.Fatal(err)
	}
	r := mux.NewRouter()
	r.HandleFunc("/", requestIDMiddleware(MainHandler))

	r.HandleFunc("/login", requestIDMiddleware(LoginFormHandler)).Methods(http.MethodGet)
	r.HandleFunc("/login", requestIDMiddleware(LoginHandler)).Methods(http.MethodPost)
	r.HandleFunc("/logout", requestIDMiddleware(LogoutHandler)).Methods(http.MethodPost)

	conn, err := grpc.Dial(cfg.UserGRPCAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	UserCli = user.NewUserClient(conn)

	fs := http.FileServer(http.Dir("assets"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	TT.MovieList, err = template.ParseFiles("base.html", "main.html")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Name: %s", TT.MovieList.Name())

	TT.Login, err = template.ParseFiles("base.html", "login.html")
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Service started on port " + cfg.Port)
	err = http.ListenAndServe(fmt.Sprintf(":%s", cfg.Port), r)
	log.Print(err)
}

func MainHandler(w http.ResponseWriter, r *http.Request) {

	page := MainPage{}

	rid := r.Context().Value(RequestIDContextKey).(string)

	var err error
	movies, err := getMovies()
	if err != nil {
		log.Printf("[%s] Get movie error: %v", rid, err)
		page.MoviesError = "Не удалось загрузить список. Код ошибки: " + rid
	}
	page.Movies = movies

	user, err := getUserByToken(r)
	if err != nil {
		log.Printf("Get user error: %v", err)
	}
	page.User = user

	log.Printf("User: %+v", page.User)

	err = TT.MovieList.ExecuteTemplate(w, "base", page)
	if err != nil {
		log.Printf("Render error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

type LoginPage struct {
	User  User
	Error string
}

func LoginFormHandler(w http.ResponseWriter, r *http.Request) {
	page := &LoginPage{}

	var err error
	page.User, err = getUserByToken(r)
	if err != nil {
		log.Printf("No user: %v", err)
		// В случае не валидного токена показываем страницу логина
		TT.Login.ExecuteTemplate(w, "base", page)
		return
	}

	TT.Login.ExecuteTemplate(w, "base", page)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	page := &LoginPage{}

	r.ParseForm()
	email := r.PostFormValue("email")
	pwd := r.PostFormValue("pwd")

	rid := r.Context().Value(RequestIDContextKey).(string)
	md := metadata.Pairs("X-Request-ID", rid)
	ctxRpc := metadata.NewOutgoingContext(context.Background(), md)

	res, err := UserCli.Login(
		ctxRpc,
		&user.LoginRequest{Email: email, Pwd: pwd},
	)

	// Что-то не так с сервисом user
	if err != nil {
		log.Printf("Get user error: %v", err)
		page.Error = "Сервис авторизации недоступен. Код ошибки: " + rid
		TT.Login.ExecuteTemplate(w, "base", page)
		return
	}

	// Ошибка логина, ее можно показать пользователю
	if res.GetError() != "" {
		page.Error = res.GetError()
		TT.Login.ExecuteTemplate(w, "base", page)
		return
	}

	tok := res.GetJwt()

	// Если пользователь успешно залогинен записываем токен в cookie
	http.SetCookie(w, &http.Cookie{Name: "jwt", Value: tok})

	jwtData, err := jwt.Parse(tok)
	if err != nil {
		// В случае не валидного токена показываем страницу логина
		TT.Login.ExecuteTemplate(w, "base", page)
		return
	}

	page.User = User{Name: jwtData.Name}
	TT.Login.ExecuteTemplate(w, "base", page)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{Name: "jwt", MaxAge: -1})
	http.Redirect(w, r, "/login", http.StatusFound)
}

func getMovies() (*[]Movie, error) {
	// _ = ctx.Value(RequestIDContextKey).(string)
	// md := metadata.Pairs("X-Request-ID", rid)
	// ctxRpc := metadata.NewOutgoingContext(context.Background(), md)

	// res, err := MovieCli.MovieList(
	// 	context.Background(),
	// 	&moviegrpc.MovieListRequest{},
	// )

	mm := &[]Movie{}
	err := get(cfg.MovieAddr+"/movie", mm)
	if err != nil {
		return nil, err
	}

	return mm, nil
}

var ERR_NO_JWT = errors.New("No 'jwt' cookie")

func getUserByToken(r *http.Request) (u User, err error) {
	tok, err := r.Cookie("jwt")
	if tok == nil {
		return u, ERR_NO_JWT
	}

	jwtData, err := jwt.Parse(tok.Value)
	if err != nil {
		return u, fmt.Errorf("Can't parse toke: %w", err)
	}

	u.Name = jwtData.Name
	u.IsPaid = jwtData.IsPaid
	return u, err
}

func getUser(r *http.Request) (u User, err error) {
	ses, err := r.Cookie("session")
	if ses == nil {
		return u, err
	}

	res := &struct {
		User
		Error string
	}{}
	err = get(cfg.UserAddr+"/user?token="+ses.Value, res)
	if err != nil {
		return u, err
	}

	if res.Error != "" {
		return u, fmt.Errorf(res.Error)
	}

	return User{
		Name:   res.Name,
		IsPaid: true,
	}, err
}

func post(url string, in url.Values, out interface{}) error {
	r, err := http.DefaultClient.PostForm(url, in)
	if err != nil {
		return fmt.Errorf("make POST request error: %w", err)
	}

	return parseResponse(r, out)
}

func get(url string, out interface{}) error {
	r, err := http.DefaultClient.Get(url)
	if err != nil {
		return fmt.Errorf("make GET request error: %w", err)
	}

	return parseResponse(r, out)
}

func parseResponse(res *http.Response, out interface{}) error {
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("read response error: %w", err)
	}

	err = json.Unmarshal(body, out)
	fmt.Printf("%s", body)
	if err != nil {
		return fmt.Errorf("parse body error '%s': %w", body, err)
	}

	return nil
}
