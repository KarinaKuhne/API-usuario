package routes

import (
	"github.com/KarinaKuhne/API-usuario/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	r := gin.Default()
	r.GET("/usuarios", controllers.GetAllUsuarios)
	r.GET("/usuarios/:id", controllers.GetUsuariobyID)
	r.POST("/usuarios", controllers.CreateUsuario)
	r.DELETE("/usuarios/:id", controllers.DeletarUsuario)
	r.PATCH("/usuarios/:id", controllers.EditarUsuario)
	r.NoRoute(controllers.RotaNaoEncontrada)
	r.Run()
}
