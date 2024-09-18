package controllers

import (
	"net/http"

	"github.com/KarinaKuhne/API-usuario/database"
	"github.com/KarinaKuhne/API-usuario/models"
	"github.com/gin-gonic/gin"
)

func GetAllUsuarios(c *gin.Context) {
	var usuarios []models.Usuario
	database.DB.Find(&usuarios)
	c.JSON(200, usuarios)
}

func CreateUsuario(c *gin.Context) {
	var usuario models.Usuario
	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error()})
		return
	}
	database.DB.Create(&usuario)
	c.JSON(http.StatusOK, usuario)
}

func GetUsuariobyID(c *gin.Context) {
	var usuario models.Usuario
	id := c.Params.ByName("id")
	database.DB.First(&usuario, id)
	if usuario.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not Found": "Viajante não encontrado"})
		return
	}
	c.JSON(http.StatusOK, usuario)
}

func DeletarUsuario(c *gin.Context) {
	var usuario models.Usuario
	id := c.Params.ByName("id")
	database.DB.Delete(&usuario, id)
	c.JSON(http.StatusOK, gin.H{"data": "Viajante deletado com sucesso"})
}

func EditarUsuario(c *gin.Context) {
	var usuario models.Usuario
	id := c.Params.ByName("id")
	database.DB.First(&usuario, id)
	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error()})
		return
	}
	database.DB.Save(&usuario)
	c.JSON(http.StatusOK, usuario)
}

func RotaNaoEncontrada(c *gin.Context) {
	c.HTML(http.StatusNotFound, "./assets/notFound.html", nil)
}

//Criar func de rota 404 ok criar asset com o html

//Adicionar validações jwt
