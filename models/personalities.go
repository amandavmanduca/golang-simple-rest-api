package models

import (
	"api-rest/database"
	"strconv"
)

type Personality struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Details string `json:"details"`
}

func GetAll() []Personality {
	var p []Personality
	database.DB.Find(&p)
	return p
}

func GetById(id string) Personality {
	var p Personality
	var numeric_id, err = strconv.Atoi(id)
	if err == nil {
		database.DB.First(&p, numeric_id)
		return p
	}
	panic("Not Found")
}

func DeleteOne(id string) Personality {
	var p Personality = GetById(id)
	database.DB.Delete(&p, p.Id)
	return p
}

func UpdateOne(id string, new Personality) Personality {
	var found Personality = GetById(id)
	updated := Personality{
		Id:      found.Id,
		Name:    new.Name,
		Details: new.Details,
	}
	database.DB.Save(&updated)
	return updated
}

func Create(p Personality) Personality {
	database.DB.Create(&p)
	return p
}
