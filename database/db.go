package database

import (
	"fmt"
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
	// Assegure-se de que as variáveis de ambiente estão definidas no ambiente onde você executa o aplicativo
	host := "travel-planner-wonder.postgres.database.azure.com"
	user := "wonder"
	port := "5432"
	dbname := "postgres"
	password := "TravelPlanner_23" // A senha deve ser definida diretamente no ambiente, não no código

	// Ajuste a string de conexão para incluir sslmode=require
	stringDeConexao := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require", host, user, password, dbname, port)

	DB, err = gorm.Open(postgres.Open(stringDeConexao), &gorm.Config{})
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados: ", err)
	}
	// AutoMigrate irá criar ou atualizar as tabelas com base nos modelos fornecidos
	DB.AutoMigrate(&models.User{}, &models.Itinerary{}, &models.ItineraryDetails{}, &models.Accommodation{}, &models.Activity{}, &models.Destination{})
}
