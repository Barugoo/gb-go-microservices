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
	"strconv"

	"google.golang.org/grpc/metadata"

	"github.com/azomio/courses/lesson4/pkg/grpc/user"
	"github.com/azomio/courses/lesson4/pkg/jwt"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"

	movieApiV1 "gateway-user/api/movie"
	paymentApiV1 "gateway-user/api/payment"
	log "gateway-user/logger"
	"gateway-user/reqdata"
)

var UserCli user.UserClient
var MovieCli movieApiV1.MovieServiceClient
var PaymentCli paymentApiV1.PaymentServiceClient

var logger = log.NewLogger()

const ServiceName = "gateway-user"

func main() {
	ctx := context.Background()

	f, err := os.Create(
		fmt.Sprintf("/var/log/super-cinema/%s.log", ServiceName),
	)
	if err != nil {
		logger.Fatalf(ctx, "error opening file: %v", err)
	}
	defer f.Close()
	logger.SetOutput(f)

	consulAddr := flag.String("consul_addr", "consul:8500", "Consul address")
	flag.Parse()

	if err := loadConfig(*consulAddr); err != nil {
		logger.Fatal(ctx, err)
	}
	r := mux.NewRouter()
	r.HandleFunc("/", reqdata.RequestIDMiddleware(MainHandler))
	r.HandleFunc("/accept-payment", reqdata.RequestIDMiddleware(AcceptPaymentHandler))
	r.HandleFunc("/decline-payment", reqdata.RequestIDMiddleware(DeclinePaymentHandler))
	r.HandleFunc("/create-payment", reqdata.RequestIDMiddleware(CreatePaymentHandler))

	r.HandleFunc("/login", reqdata.RequestIDMiddleware(LoginFormHandler)).Methods(http.MethodGet)
	r.HandleFunc("/login", reqdata.RequestIDMiddleware(LoginHandler)).Methods(http.MethodPost)
	r.HandleFunc("/logout", reqdata.RequestIDMiddleware(LogoutHandler)).Methods(http.MethodPost)

	conn, err := grpc.Dial(cfg.UserGRPCAddr, grpc.WithInsecure())
	if err != nil {
		logger.Fatalf(ctx, "did not connect: %s", err)
	}
	UserCli = user.NewUserClient(conn)

	connMovie, err := grpc.Dial(cfg.MovieGRPCAddr, grpc.WithInsecure())
	if err != nil {
		logger.Fatalf(ctx, "did not connect: %s", err)
	}
	MovieCli = movieApiV1.NewMovieServiceClient(connMovie)

	connPayment, err := grpc.Dial(cfg.PaymentGRPCAddr, grpc.WithInsecure())
	if err != nil {
		logger.Fatalf(ctx, "did not connect: %s", err)
	}
	PaymentCli = paymentApiV1.NewPaymentServiceClient(connPayment)

	fs := http.FileServer(http.Dir("assets"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	TT.MovieList, err = template.ParseFiles("base.html", "main.html")
	if err != nil {
		logger.Fatal(ctx, err)
	}

	logger.Infof(ctx, "Name: %s", TT.MovieList.Name())

	TT.Login, err = template.ParseFiles("base.html", "login.html")
	if err != nil {
		logger.Fatal(ctx, err)
	}

	logger.Infof(ctx, "Service started on port "+cfg.Port)
	err = http.ListenAndServe(fmt.Sprintf(":%s", cfg.Port), r)
	logger.Fatal(ctx, err)
}

func CreatePaymentFormHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	user, err := getUserByToken(r)
	if err != nil {
		logger.Errorf(ctx, "Get user error: %v", err)
	}
	render.RenderTemplate(w, "payform", struct{ Uid string }{fmt.Sprintf("%d", user.ID)})
}

type Msg struct {
	Msg     string
	BackURL string
}

func CreatePaymentHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	ctx := r.Context()

	uidRaw := r.FormValue("uid")
	uid, _ := strconv.Atoi(uidRaw)

	amountRaw := r.FormValue("amount")
	amount, _ := strconv.Atoi(uidRaw)

	req := &paymentApiV1.GenerateDepositLinkRequest{
		UserId: int64(uid),
		Amount: int64(amount),
	}

	resp, err := PaymentCli.GenerateDepositLink(ctx, req)
	if err != nil {
		return
	}

	// resp.Payload

	return
}

func CreatePaymentHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	orderIDRaw := r.URL.Query().Get("order_id")
	if orderIDRaw == "" {
		logger.Errorf(ctx, "Empty order id")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	orderID, _ := strconv.Atoi(orderIDRaw)

	req := &paymentApiV1.UpdateTransactionStatusRequest{
		TransactionId: int64(orderID),
		NewStatus:     paymentApiV1.TransactionStatus_DONE,
	}

	_, err := PaymentCli.UpdateTransactionStatus(ctx, req)
	if err != nil {
		logger.Errorf(ctx, "Update transaction status")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func AcceptPaymentHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	orderIDRaw := r.URL.Query().Get("order_id")
	if orderIDRaw == "" {
		logger.Errorf(ctx, "Empty order id")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	orderID, _ := strconv.Atoi(orderIDRaw)

	req := &paymentApiV1.UpdateTransactionStatusRequest{
		TransactionId: int64(orderID),
		NewStatus:     paymentApiV1.TransactionStatus_DONE,
	}

	_, err := PaymentCli.UpdateTransactionStatus(ctx, req)
	if err != nil {
		logger.Errorf(ctx, "Update transaction status")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func DeclinePaymentHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	orderIDRaw := r.URL.Query().Get("order_id")
	if orderIDRaw == "" {
		logger.Errorf(ctx, "Empty order id")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	orderID, _ := strconv.Atoi(orderIDRaw)

	req := &paymentApiV1.UpdateTransactionStatusRequest{
		TransactionId: int64(orderID),
		NewStatus:     paymentApiV1.TransactionStatus_FAILURE,
	}

	_, err := PaymentCli.UpdateTransactionStatus(ctx, req)
	if err != nil {
		logger.Errorf(ctx, "Update transaction status")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func MainHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	page := MainPage{}

	var err error
	movies, err := getMovies(ctx)
	if err != nil {
		logger.Errorf(ctx, "Get movie error: %s", err.Error())
		page.MoviesError = "Не удалось загрузить список. Код ошибки: " + reqdata.GetRequestID(ctx)
	}
	page.Movies = movies

	user, err := getUserByToken(r)
	if err != nil {
		logger.Errorf(ctx, "Get user error: %v", err)
	}
	page.User = user

	logger.Infof(ctx, "User: %+v", page.User)

	err = TT.MovieList.ExecuteTemplate(w, "base", page)
	if err != nil {
		logger.Errorf(ctx, "Render error: %v", err)
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
		logger.Errorf(r.Context(), "No user: %v", err)
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

	md := metadata.Pairs(reqdata.RequestIDHeader, reqdata.GetRequestID(r.Context()))
	ctxRpc := metadata.NewOutgoingContext(context.Background(), md)

	res, err := UserCli.Login(
		ctxRpc,
		&user.LoginRequest{Email: email, Pwd: pwd},
	)

	// Что-то не так с сервисом user
	if err != nil {
		logger.Errorf(r.Context(), "Get user error: %v", err)
		page.Error = "Сервис авторизации недоступен. Код ошибки: " + reqdata.GetRequestID(r.Context())
		TT.Login.ExecuteTemplate(w, "base", page)
		return
	}

	// Ошибка логина, ее можно показать пользователю
	if res.GetError() != "" {
		logger.Errorf(r.Context(), "Login error: %v", err)
		page.Error = res.GetError()
		TT.Login.ExecuteTemplate(w, "base", page)
		return
	}

	logger.Infof(r.Context(), "Successfuly logged in: %s", email)

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

func getMovies(ctx context.Context) ([]*Movie, error) {
	rid := ctx.Value(reqdata.RequestIDContextKey).(string)
	md := metadata.Pairs("X-Request-ID", rid)
	ctx = metadata.NewOutgoingContext(ctx, md)

	resp, err := MovieCli.ListMovies(ctx, &movieApiV1.ListMoviesRequest{})
	if err != nil {
		return nil, err
	}

	res := make([]*Movie, 0, len(resp.Items))
	for _, m := range resp.Items {
		movie := &Movie{
			ID:       int(m.Id),
			Name:     m.Name,
			Poster:   m.Poster,
			MovieUrl: m.Url,
		}
		res = append(res, movie)
	}
	return res, nil
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
