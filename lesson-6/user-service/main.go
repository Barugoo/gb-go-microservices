package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"runtime/debug"
	"strconv"

	pb "user-service/api"
	"user-service/reqdata"

	log "user-service/logger"

	render "github.com/barugoo/gb-go-microservices/lesson-2/pkg/render"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var logger = log.NewLogger()

func RecoverInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	var rid string
	defer func() {
		if r := recover(); r != nil {
			logger.Errorf(ctx, "Recover from %v, %s", rid, r, debug.Stack())
			err = status.Errorf(codes.Internal, "Internal error")
			return
		}
	}()
	md, _ := metadata.FromIncomingContext(ctx)
	if ridSli := md.Get(reqdata.RequestIDHeader); len(ridSli) > 0 {
		rid = ridSli[0]
	}

	return handler(ctx, req)
}

const ServiceName = "user-service"

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

	srv := grpc.NewServer(grpc.UnaryInterceptor(RecoverInterceptor))

	pb.RegisterUserServer(srv, &UserService{})

	listener, err := net.Listen("tcp", ":9096")
	if err != nil {
		logger.Fatalf(ctx, "failed to listen: %v", err)
	}

	logger.Infof(ctx, "Starting server on localhost: 9096")
	srv.Serve(listener)
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	email := r.Form.Get("email")
	password := r.Form.Get("password")
	name := r.Form.Get("name")

	// TO DO better validation (especially for password)
	if email == "" || password == "" || name == "" {
		render.RenderJSONErr(w, "Wrong input data", http.StatusBadRequest)
		return
	}

	existUsr := UU.GetByEmail(email)
	if existUsr != nil {
		render.RenderJSONErr(w, "Email exists", http.StatusBadRequest)
		return
	}

	newUser := &User{
		Email: email,
		Pwd:   password,
		Name:  name,
	}
	usr := UU.CreateUser(newUser)

	render.RenderJSON(w, usr)
	return
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	email := r.Form.Get("email")
	password := r.Form.Get("password")

	// TO DO better validation (especially for password)
	if email == "" || password == "" {
		logger.Errorf(r.Context(), "empty email or password %s", email)
		render.RenderJSONErr(w, "Wrong email or password", http.StatusBadRequest)
		return
	}

	usr := UU.GetByEmail(email)
	if usr == nil || usr.Pwd != password {
		logger.Errorf(r.Context(), "wrong email or password %s", email)
		render.RenderJSONErr(w, "Wrong email or password", http.StatusUnauthorized)
		return
	}

	logger.Infof(r.Context(), "got user by email %s", email)

	render.RenderJSON(w, usr)
	return
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	token := r.Form.Get("token")

	usr := UU.GetByToken(token)

	if usr == nil {
		render.RenderJSONErr(w, "Пользователь не найден", http.StatusNotFound)
		return
	}

	render.RenderJSON(w, usr)
	return
}

func userPatchHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	idStr := r.FormValue("id")
	isPaidStr := r.FormValue("is_paid")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		render.RenderJSONErr(w, "Invalid 'id': "+err.Error(), http.StatusBadRequest)
		return
	}

	usr := UU.GetByID(id)

	isPaid, err := strconv.ParseBool(isPaidStr)
	if err != nil {
		render.RenderJSONErr(w, "Invalid 'is_paid': "+err.Error(), http.StatusBadRequest)
		return
	}

	if usr == nil {
		render.RenderJSONErr(w, "Пользователь не найден. ID: "+idStr, http.StatusNotFound)
		return
	}

	usr.IsPaid = isPaid

	render.RenderJSON(w, usr)
	return
}
