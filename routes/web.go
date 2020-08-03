package routes

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/heliosmc89/gallery-app-with-go/controllers"
)

// RegisterRoutes is used to register the routes we need for the web application
func RegisterRoutes() {
	mux := chi.NewRouter()
	mux.NotFound(http.HandleFunc(controllers.Show404Page().Render))
	mux.MethodFunc(http.MethodGet, "/", controllers.ShowHomePage().Render)
	mux.MethodFunc(http.MethodGet, "/contact", controllers.ShowContactPage().Render)
	mux.MethodFunc(http.MethodGet, "/cookietest", controllers.ShowUserCookie)

	// auth routes
	mux.Route("/auth", func(mux chi.Router) {
		mux.MethodFunc(http.MethodPost, "/register", controllers.ParseRegisterForm)
		mux.MethodFunc(http.MethodPost, "/login", controllers.ParseLoginForm)
		mux.MethodFunc(http.MethodGet, "/login", controllers.ShowLoginForm().Render)
		mux.MethodFunc(MethodGet, "/register", controllers.ShowRegisterForm().Render)
	})
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("APP_PORT")), mux))
}
