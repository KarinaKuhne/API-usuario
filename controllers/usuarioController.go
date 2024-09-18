package controllers

import (
	"database/sql"
	"github.com/KarinaKuhne/API-usuario/controllers"
	"github.com/KarinaKuhne/API-usuario/models"
	"github.com/KarinaKuhne/API-usuario/routes"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Função para criar um usuário
func CreateUser(c *gin.Context, db *sql.DB) {
	var usuario models.Usuario
	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `INSERT INTO Usuario (nome, data_nascimento, email, senha, cidade_residencia_id, tipo)
	          VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	err := db.QueryRow(query, usuario.Nome, usuario.DataNascimento, usuario.Email, usuario.Senha, usuario.CidadeResidenciaID, usuario.Tipo).Scan(&usuario.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar usuário"})
		return
	}

	c.JSON(http.StatusOK, usuario)
}

// Função para buscar todos os usuários
func GetAllUsers(c *gin.Context, db *sql.DB) {
	rows, err := db.Query("SELECT id, nome, data_nascimento, email, cidade_residencia_id, tipo FROM Usuario")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao obter usuários"})
		return
	}
	defer rows.Close()

	var usuarios []models.Usuario
	for rows.Next() {
		var usuario models.Usuario
		if err := rows.Scan(&usuario.ID, &usuario.Nome, &usuario.DataNascimento, &usuario.Email, &usuario.CidadeResidenciaID, &usuario.Tipo); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao processar usuários"})
			return
		}
		usuarios = append(usuarios, usuario)
	}

	c.JSON(http.StatusOK, usuarios)
}
