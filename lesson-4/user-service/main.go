package main

import (
	"database/sql"
	"log"
	"net"

	pb "user-service/api"

	_ "github.com/lib/pq"

	"google.golang.org/grpc"
)

func main() {
	db, err := sql.Open(
		"postgres",
		"postgres://postgres:5432/postgres?sslmode=disable&user=postgres&database=postgres&password=postgres",
	)
	if err != nil {
		log.Fatal(err)
	}
	for db.Ping() != nil {
		log.Println(db.Ping())
	}
	service := &UserService{
		db: &UserStorage{db: db},
	}

	srv := grpc.NewServer()

	pb.RegisterUserServer(srv, service)

	listener, err := net.Listen("tcp", ":9094")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("Starting server on localhost: 9094")
	srv.Serve(listener)
}

// func (s *UserService) registerHandler(w http.ResponseWriter, r *http.Request) {
// 	r.ParseForm()

// 	email := r.Form.Get("email")
// 	password := r.Form.Get("password")
// 	name := r.Form.Get("name")

// 	// TO DO better validation (especially for password)
// 	if email == "" || password == "" || name == "" {
// 		render.RenderJSONErr(w, "Wrong input data", http.StatusBadRequest)
// 		return
// 	}

// 	existUsr := UU.GetByEmail(email)
// 	if existUsr != nil {
// 		render.RenderJSONErr(w, "Email exists", http.StatusBadRequest)
// 		return
// 	}

// 	newUser := &User{
// 		Email: email,
// 		Pwd:   password,
// 		Name:  name,
// 	}
// 	usr := UU.CreateUser(newUser)

// 	render.RenderJSON(w, usr)
// 	return
// }

// func loginHandler(w http.ResponseWriter, r *http.Request) {
// 	r.ParseForm()

// 	email := r.Form.Get("email")
// 	password := r.Form.Get("password")

// 	// TO DO better validation (especially for password)
// 	if email == "" || password == "" {
// 		render.RenderJSONErr(w, "Wrong email or password", http.StatusBadRequest)
// 		return
// 	}

// 	usr := UU.GetByEmail(email)
// 	if usr == nil || usr.Pwd != password {
// 		render.RenderJSONErr(w, "Wrong email or password", http.StatusUnauthorized)
// 		return
// 	}

// 	render.RenderJSON(w, usr)
// 	return
// }

// func userHandler(w http.ResponseWriter, r *http.Request) {
// 	r.ParseForm()

// 	token := r.Form.Get("token")

// 	usr := UU.GetByToken(token)

// 	if usr == nil {
// 		render.RenderJSONErr(w, "Пользователь не найден", http.StatusNotFound)
// 		return
// 	}

// 	render.RenderJSON(w, usr)
// 	return
// }

// func userPatchHandler(w http.ResponseWriter, r *http.Request) {
// 	r.ParseForm()

// 	idStr := r.FormValue("id")
// 	isPaidStr := r.FormValue("is_paid")

// 	id, err := strconv.Atoi(idStr)
// 	if err != nil {
// 		render.RenderJSONErr(w, "Invalid 'id': "+err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	usr := UU.GetByID(id)

// 	isPaid, err := strconv.ParseBool(isPaidStr)
// 	if err != nil {
// 		render.RenderJSONErr(w, "Invalid 'is_paid': "+err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	if usr == nil {
// 		render.RenderJSONErr(w, "Пользователь не найден. ID: "+idStr, http.StatusNotFound)
// 		return
// 	}

// 	usr.IsPaid = isPaid

// 	render.RenderJSON(w, usr)
// 	return
// }
