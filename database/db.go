package database

import (
	"log"

	"github.com/KarinaKuhne/API-usuario/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDB() {
	stringDeConexao := "host=dpg-crh1mp56l47c73c1f6bg-a.virginia-postgres.render.com user=admintrv password=w6Di1sG93ZcUutxvWcWJzoLPdNTjVRDR dbname=travelerhomolog port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}

	DB.AutoMigrate(&models.Usuario{})
}
