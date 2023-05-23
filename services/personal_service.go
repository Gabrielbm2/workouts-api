package services

import (
	"context"
	"example/CadastroGo/db"
	"example/CadastroGo/models"
	"fmt"
)

func RegisterPersonal(nome, sobrenome, email, cpf, senha string) (*models.Personal, error) {
	db := db.GetDatabase()

	sqlPersonal := `
        INSERT INTO personal (nome, sobrenome, email, cpf, senha) VALUES ($1, $2, $3, $4, $5)
    `
	newPersonal := &models.Personal{}
	_, err := db.Exec(context.Background(), sqlPersonal, nome, sobrenome, email, cpf, senha)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("erro ao criar conta de Personal: %s", err)
	}

	return newPersonal, nil
}
