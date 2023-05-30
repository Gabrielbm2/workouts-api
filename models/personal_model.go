package models

import (
	"CadastroGo/db"
	"context"
	"fmt"
	_ "fmt"
)

type Personal struct {
	Nome      string `json:"nome"`
	Sobrenome string `json:"sobrenome"`
	Email     string `json:"email"`
	CPF       string `json:"cpf"`
	Senha     string `json:"senha"`
}

func CreatePersonalTable() error {
	sql := `
		CREATE TABLE IF NOT EXISTS personal (
			id SERIAL PRIMARY KEY,
			nome VARCHAR(255) NOT NULL,
			sobrenome VARCHAR(255) NOT NULL,
			email VARCHAR(255) NOT NULL,
			cpf VARCHAR(11) NOT NULL,
			senha VARCHAR(255) NOT NULL
		);
	`
	fmt.Println("Criando Tabela do Personal")
	_, err := db.GetDatabase().Exec(context.Background(), sql)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("erro ao criar tabela do Personal: %s", err)
	} else {
		fmt.Println("Tabela do Personal criada")
	}
	return nil
}
