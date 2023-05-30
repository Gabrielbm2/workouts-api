package models

import (
	"CadastroGo/db"
	"context"
	"fmt"
	_ "fmt"
	_ "github.com/lib/pq"
)

type UserRegister struct {
	Nome      string `json:"nome"`
	Sobrenome string `json:"sobrenome"`
	Email     string `json:"email"`
	Cpf       string `json:"cpf"`
	Senha     string `json:"senha"`
}

func CreateUserTable() error {
	var sql = `
		CREATE TABLE IF NOT EXISTS user (
			id SERIAL PRIMARY KEY,
			nome VARCHAR(255),
			sobrenome VARCHAR(255),
			email VARCHAR(255),
			cpf BIGINT,
			senha VARCHAR(255)
		)
	`
	fmt.Println("Tabela do usuario criada com sucesso!")
	_, err := db.GetDatabase().Exec(context.Background(), sql)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("erro ao criar a tabela do usuario: %s", err)
	} else {
		fmt.Println(`Tabela do usuario criada`)
	}
	return nil
}
