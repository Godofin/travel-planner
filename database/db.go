package database

import (
	"log"

	"github.com/Godofin/travel-planner/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	stringDeConexao := "host=dpg-cpkcrcfsc6pc73er6gi0-a.oregon-postgres.render.com user=wonder password=ZLzj1XIKzR3TK0D0iMkA1PBKb2H0nRnY dbname=travelplannerwonder port=5432 sslmode=require"
	DB, err = gorm.Open(postgres.Open(stringDeConexao), &gorm.Config{})
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}
	DB.AutoMigrate(&models.User{})
}
