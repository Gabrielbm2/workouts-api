package controllers

import (
	"CadastroGo/models"
	"CadastroGo/services"
	"encoding/json"
	"fmt"
	"net/http"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Erro ao ler os dados da requisição", http.StatusBadRequest)
		return
	}

	personalID := "ID_DO_PERSONAL"

	_, err = services.RegisterUser(personalID, user.Nome, user.Sobrenome, user.Email, user.Cpf, user.Senha)
	if err != nil {
		http.Error(w, "Erro ao registrar o usuário", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintf(w, "Registro do usuário concluído com sucesso")
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Erro ao ler os dados da requisição", http.StatusBadRequest)
		return
	}

	authenticatedUser, err := services.LoginUser(user.Identifier, user.Senha)
	if err != nil {
		http.Error(w, "Email/CPF ou senha incorretos", http.StatusUnauthorized)
		return
	}

	jsonResponse, err := json.Marshal(authenticatedUser)
	if err != nil {
		http.Error(w, "Erro ao gerar a resposta", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
