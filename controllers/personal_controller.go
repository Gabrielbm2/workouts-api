package controllers

import (
	"encoding/json"
	"net/http"

	"CadastroGo/models"
	"CadastroGo/services"
)

func RegisterPersonal(w http.ResponseWriter, r *http.Request) {
	var personal models.Personal
	err := json.NewDecoder(r.Body).Decode(&personal)
	if err != nil {
		http.Error(w, "Erro ao ler os dados da requisição", http.StatusBadRequest)
		return
	}

	newPersonal, err := services.RegisterPersonal(personal.Nome, personal.Sobrenome, personal.Email, personal.CPF, personal.Senha)
	if err != nil {
		http.Error(w, "Erro ao criar conta de Personal", http.StatusInternalServerError)
		return
	}

	_ = json.NewEncoder(w).Encode(newPersonal)
}

func LoginPersonal(w http.ResponseWriter, r *http.Request) {
	type LoginRequest struct {
		Identifier string `json:"identifier"`
		Senha      string `json:"senha"`
	}

	var request LoginRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Erro ao ler os dados da requisição", http.StatusBadRequest)
		return
	}

	personal, err := services.LoginPersonal(request.Identifier, request.Senha)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	_ = json.NewEncoder(w).Encode(personal)
}
