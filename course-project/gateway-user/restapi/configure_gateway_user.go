// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"log"
	"net/http"
	"os"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"google.golang.org/grpc"

	pb "github.com/barugoo/gb-go-microservices/course-project/gateway-user/cmd/proto"
	"github.com/barugoo/gb-go-microservices/course-project/gateway-user/restapi/operations"
	"github.com/barugoo/gb-go-microservices/course-project/gateway-user/restapi/operations/auth"
	"github.com/barugoo/gb-go-microservices/course-project/gateway-user/restapi/operations/movies"
	"github.com/barugoo/gb-go-microservices/course-project/gateway-user/restapi/operations/profile"
	"github.com/barugoo/gb-go-microservices/course-project/gateway-user/restapi/service"
)

//go:generate swagger generate server --target ../../gateway-user --name GatewayUser --spec ../swagger.yaml --principal interface{}

func configureFlags(api *operations.GatewayUserAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.GatewayUserAPI) http.Handler {
	log.Println("configuring API")

	authServiceAddr := os.Getenv("AUTH_SERVICE_GRPC_ADDR")
	if authServiceAddr == "" {
		log.Fatal("unable to retreive env AUTH_SERVICE_GRPC_ADDR")
	}
	conn, err := grpc.Dial(authServiceAddr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewAuthServiceClient(conn)

	s := &service.GatewayUserService{
		AuthServiceClient: c,
	}

	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Applies when the "Authorization" header is set
	if api.BearerAuthAuth == nil {
		api.BearerAuthAuth = func(token string) (interface{}, error) {
			return nil, errors.NotImplemented("api key auth (bearerAuth) Authorization from header param [Authorization] has not yet been implemented")
		}
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()
	if api.MoviesGetMovieHandler == nil {
		api.MoviesGetMovieHandler = movies.GetMovieHandlerFunc(func(params movies.GetMovieParams) middleware.Responder {
			return middleware.NotImplemented("operation movies.GetMovie has not yet been implemented")
		})
	}
	if api.ProfileGetProfileHandler == nil {
		api.ProfileGetProfileHandler = profile.GetProfileHandlerFunc(func(params profile.GetProfileParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation profile.GetProfile has not yet been implemented")
		})
	}
	if api.MoviesListMoviesHandler == nil {
		api.MoviesListMoviesHandler = movies.ListMoviesHandlerFunc(func(params movies.ListMoviesParams) middleware.Responder {
			return middleware.NotImplemented("operation movies.ListMovies has not yet been implemented")
		})
	}
	if api.ProfileListOrdersHandler == nil {
		api.ProfileListOrdersHandler = profile.ListOrdersHandlerFunc(func(params profile.ListOrdersParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation profile.ListOrders has not yet been implemented")
		})
	}
	if api.ProfileListPaymentsHandler == nil {
		api.ProfileListPaymentsHandler = profile.ListPaymentsHandlerFunc(func(params profile.ListPaymentsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation profile.ListPayments has not yet been implemented")
		})
	}
	api.AuthLoginUserHandler = auth.LoginUserHandlerFunc(s.Login)
	api.AuthRegisterUserHandler = auth.RegisterUserHandlerFunc(s.Register)

	if api.AuthRegisterUserHandler == nil {
		api.AuthRegisterUserHandler = auth.RegisterUserHandlerFunc(func(params auth.RegisterUserParams) middleware.Responder {
			return middleware.NotImplemented("operation auth.RegisterUser has not yet been implemented")
		})
	}
	if api.MoviesRentMovieHandler == nil {
		api.MoviesRentMovieHandler = movies.RentMovieHandlerFunc(func(params movies.RentMovieParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation movies.RentMovie has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
