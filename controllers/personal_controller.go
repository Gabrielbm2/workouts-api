package controllers

import (
	"encoding/json"
	"net/http"

	"example/CadastroGo/models"
	"example/CadastroGo/services"
)

func RegisterPersonal(w http.ResponseWriter, r *http.Request) {
	var personal models.Personal
	err := json.NewDecoder(r.Body).Decode(&personal)
	if err != nil {
		http.Error(w, "Erro ao ler os dados da requisição", http.StatusBadRequest)
		return
	}

	// Chame a função RegisterPersonal do pacote services para criar a conta de Personal
	newPersonal, err := services.RegisterPersonal(personal.Nome, personal.Sobrenome, personal.Email, personal.CPF, personal.Senha)
	if err != nil {
		http.Error(w, "Erro ao criar conta de Personal", http.StatusInternalServerError)
		return
	}

	// Retorne a resposta adequada ao cliente
	json.NewEncoder(w).Encode(newPersonal)
}
