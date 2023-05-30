package routes

import (
	"CadastroGo/controllers"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func LoadRoutes(routes *chi.Mux) {
	routes.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("pong"))
	})

	routes.Post("/gym/personal/register", controllers.RegisterPersonal)
	routes.Post("/gym/personal/login", controllers.LoginPersonal)

	routes.Post("/gym/user/register", controllers.RegisterUser)
	routes.Post("/gym/user/login", controllers.LoginUser)

}
