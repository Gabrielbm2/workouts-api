package services

import (
	"CadastroGo/db"
	"CadastroGo/models"
	"context"
	"fmt"
)

func RegisterUser(personalID, nome, sobrenome, email, cpf, senha string) (*models.UserRegister, error) {
	database := db.GetDatabase()

	sqlUser := `
        INSERT INTO user (personal_id, nome, sobrenome, email, cpf, senha) VALUES ($1, $2, $3, $4, $5, $6)
    `

	newUser := &models.UserRegister{}
	_, err := database.Exec(context.Background(), sqlUser, personalID, nome, sobrenome, email, cpf, senha)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("Erro ao criar conta de usuário: %s", err)
	}

	return newUser, nil
}
