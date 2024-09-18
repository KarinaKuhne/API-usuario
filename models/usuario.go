package models

type Usuario struct {
	ID                 int    `json:"id"`
	Nome               string `json:"nome"`
	DataNascimento     string `json:"data_nascimento"`
	Email              string `json:"email"`
	Senha              string `json:"senha"`
	CidadeResidenciaID int    `json:"cidade_residencia_id"`
	Tipo               string `json:"tipo"`
}
