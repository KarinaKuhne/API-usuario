package models

import (
	"gorm.io/gorm"
)

type Usuario struct {
	gorm.Model
	ID                 int    `json:"id"`
	Nome               string `json:"nome"`
	DataNascimento     string `json:"data_nascimento"`
	Email              string `json:"email"`
	Senha              string `json:"senha"`
	CidadeResidenciaID int    `json:"cidade_residencia_id"`
	Tipo               string `json:"tipo"`
}
