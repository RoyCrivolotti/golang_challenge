package configuration

import (
	"golangchallenge/internal/infrastructure/configuration/authentication"
	"golangchallenge/internal/infrastructure/configuration/handlers"
	"net/http"

	"firebase.google.com/go/auth"
	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"
)

type Server struct {
	appRouter *chi.Mux
}

func NewServer(appRouter *chi.Mux) *Server {
	return &Server{
		appRouter: appRouter,
	}
}

func (server *Server) Initialize(firebaseAuth *auth.Client, authenticationMiddleware authentication.IAuthenticationMiddleware, db *gorm.DB) {
	services := handlers.NewServiceHandler(firebaseAuth, db)
	controllers := handlers.NewControllerHandler(services)

	server.appRouter.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, err := w.Write([]byte("pong"))
		if err != nil {
			http.Error(w, "Encountered an unexpected server error when trying to send a response", 500)
			return
		}
	})

	//Defining routes
	server.appRouter.With(authenticationMiddleware.Authenticate).Post("/courses/sort", controllers.CourseController.SortCourses)

	server.appRouter.Post("/user/login", controllers.AuthenticationController.SignIn)
	server.appRouter.Post("/user/signup", controllers.AuthenticationController.SignUp)
}
