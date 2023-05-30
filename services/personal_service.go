package services

import (
	"CadastroGo/db"
	"CadastroGo/models"
	"context"
	"database/sql"
	"fmt"
)

func RegisterPersonal(nome, sobrenome, email, cpf, senha string) (*models.Personal, error) {
	database := db.GetDatabase()

	sqlNewPersonal := `
        INSERT INTO personal (nome, sobrenome, email, cpf, senha) VALUES ($1, $2, $3, $4, $5)
    `
	newPersonal := &models.Personal{}
	_, err := database.Exec(context.Background(), sqlNewPersonal, nome, sobrenome, email, cpf, senha)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("erro ao criar conta de Personal: %s", err)
	}

	return newPersonal, nil
}

func LoginPersonal(identifier, senha string) (*models.Personal, error) {
	database := db.GetDatabase()

	sqlPersonalLogin := `
        SELECT id, nome, sobrenome, email, cpf, senha FROM personal WHERE (email = $1 OR cpf = $1) AND senha = $2
    `

	personal := &models.Personal{}
	err := database.QueryRow(context.Background(), sqlPersonalLogin, identifier, senha).Scan(
		&personal.Nome,
		&personal.Sobrenome,
		&personal.Email,
		&personal.CPF,
		&personal.Senha,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Email/CPF ou senha incorretos")
		}
		return nil, fmt.Errorf("erro ao fazer login do personal: %s", err)
	}

	return personal, nil
}
