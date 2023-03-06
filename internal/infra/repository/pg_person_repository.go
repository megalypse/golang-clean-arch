package repository

import (
	"log"
	"time"

	"github.com/megalypse/golang-clean-arch/internal/config"
	"github.com/megalypse/golang-clean-arch/internal/domain/models"
)

type PgPersonRepository struct{}

func (PgPersonRepository) CreatePerson(person models.Person) models.Person {
	db := config.GetPgDbConnection()
	statement := `
	INSERT INTO people (fullname, age, created_at, updated_at, deleted_at)
	VALUES ($1, $2, $3, $4, $5)
	`
	result, err := db.Exec(statement, person.Fullname, person.Age, time.Now())
	if err != nil {
		panic(err)
	}

	log.Println(result)
	return models.Person{}
}

func (PgPersonRepository) GetPersonById(id int) models.Person {
	db := config.GetPgDbConnection()
	stmt := `
	SELECT * FROM people p
	WHERE p.id = $1
	`

	result, err := db.Exec(stmt, id)
	if err != nil {
		panic(err)
	}

	log.Println(result)
	return models.Person{}
}
