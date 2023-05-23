package routes

import (
	"CadastroGo/controllers"
	"github.com/go-chi/chi/v5"
	"net/http"
)

//O package "routes" define as rotas (endpoints) da aplicação, mapeando cada uma para a função correspondente em "controllers" e incluindo também uma rota para servir arquivos estáticos.

func LoadRoutes(routes *chi.Mux) {
	routes.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	routes.Post("/gym/personal/register", controllers.RegisterPersonal)
	routes.Post("/gym/personal/login", controllers.LoginPersonal)

	routes.Post("/gym/user/register", controllers.RegisterUser)
	routes.Post("/gym/user/login", controllers.LoginUser)

}
