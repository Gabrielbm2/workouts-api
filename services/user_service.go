package services

import (
	"CadastroGo/db"
	"CadastroGo/models"
	"context"
	"database/sql"
	"fmt"
)

func RegisterUser(personalID, nome, sobrenome, email, cpf, senha string) (*models.User, error) {
	database := db.GetDatabase()

	sqlUser := `
        INSERT INTO user (personal_id, nome, sobrenome, email, cpf, senha) VALUES ($1, $2, $3, $4, $5, $6)
    `

	newUser := &models.User{}
	_, err := database.Exec(context.Background(), sqlUser, personalID, nome, sobrenome, email, cpf, senha)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("Erro ao criar conta de usuário: %s", err)
	}

	return newUser, nil
}

func LoginUser(identifier, senha string) (*models.User, error) {
	db := db.GetDatabase()

	sqlQuery := `
		SELECT id, nome, sobrenome, email, cpf FROM user WHERE (email = $1 OR cpf = $1) AND senha = $2
	`

	user := &models.User{}
	err := db.QueryRow(context.Background(), sqlQuery, identifier, senha).Scan(
		&user.Nome,
		&user.Sobrenome,
		&user.Email,
		&user.Cpf,
		&user.Senha,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Email/CPF ou senha incorretos")
		}
		return nil, fmt.Errorf("Erro ao fazer login do usuário: %s", err)
	}

	return user, nil
}
