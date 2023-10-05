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
	stringDeConexao := "host=localhost user=postgres password=senha dbname=postgres port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}
	DB.AutoMigrate(&models.User{}, &models.Itinerary{}, &models.ItineraryDetails{}, &models.Accommodation{}, &models.Activity{}, &models.Destination{})
}
