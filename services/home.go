package services

import (
	"fmt"

	"github.com/TalesPalma/api-go-rest/database"
	"github.com/TalesPalma/api-go-rest/models"
)

func Home() []models.Personalidade {
	var personalidades []models.Personalidade
	database.DB.Find(&personalidades)
	return personalidades
}

func Insert(persona models.Personalidade) {
	database.DB.Create(&persona)
	fmt.Println("Inserido dados com sucesso via post", persona)
}

func DeleteById(id int) {
	database.DB.Delete(&models.Personalidade{}, id)
	fmt.Println("Deletado com sucesso")
}

func PutByid(id int, newPersona models.Personalidade) models.Personalidade {
	var novaPersonalidade = models.Personalidade{
		Id:       id,
		Nome:     newPersona.Nome,
		Historia: newPersona.Historia,
	}
	database.DB.Updates(novaPersonalidade)
	return newPersona
}

func GetById(id int) models.Personalidade {
	var personalidade models.Personalidade
	database.DB.First(&personalidade, id)
	return personalidade
}
