package routes

import (
	"database/sql"
	"github.com/KarinaKuhne/API-usuario/controllers"
	"github.com/KarinaKuhne/API-usuario/models"
	"github.com/KarinaKuhne/API-usuario/routes"
	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.Engine, db *sql.DB) {
	r.POST("/usuarios", func(c *gin.Context) {
		controllers.CreateUser(c, db)
	})
	r.GET("/usuarios", func(c *gin.Context) {
		controllers.GetAllUsers(c, db)
	})
}
