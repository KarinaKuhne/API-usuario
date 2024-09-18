package main

import (
	"database/sql"
	"fmt"
	"github.com/KarinaKuhne/API-usuario/controllers"
	"github.com/KarinaKuhne/API-usuario/models"
	"github.com/KarinaKuhne/API-usuario/routes"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func setupDB() (*sql.DB, error) {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		return nil, fmt.Errorf("DATABASE_URL não foi definida")
	}

	connStr := fmt.Sprintf("%s?sslmode=require", dbURL)
	return sql.Open("postgres", connStr)
}

func main() {
	// Conectar ao banco de dados
	db, err := setupDB()
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados: ", err)
	}
	defer db.Close()

	// Criar uma nova instância do Gin
	r := gin.Default()

	// Registrar rotas
	routes.SetupUserRoutes(r, db)

	// Iniciar o servidor
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
