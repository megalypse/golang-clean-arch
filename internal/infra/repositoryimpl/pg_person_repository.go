package repositoryimpl

import (
	"database/sql"
	"log"
	"time"

	"github.com/lib/pq"
	"github.com/megalypse/golang-clean-arch/internal/config"
	"github.com/megalypse/golang-clean-arch/internal/domain/models"
)

// TODO: better connection management per request
type PgPersonRepository struct{}

func (rep PgPersonRepository) CreatePerson(person models.Person) int {
	db := config.GetPgDbConnection()
	defer db.Close()

	sttmt := `
	INSERT INTO people (fullname, age, email, created_at, updated_at, deleted_at)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING id
	`

	/*
		LastInsertId() is not supported by lib/pq, so this needs to be manually done to get
		the new row ID.
	*/
	var newRowId int
	err := db.QueryRow(
		sttmt, person.Fullname, person.Age, person.Email, time.Now(), nil, nil,
	).Scan(&newRowId)
	if err != nil {
		panic(err)
	}

	return newRowId
}

func (rep PgPersonRepository) GetPersonById(id int) *models.Person {
	db := config.GetPgDbConnection()
	defer db.Close()

	sttmt := `
	SELECT * FROM people p
	WHERE p.id = $1
	`

	person := models.Person{}
	row := db.QueryRow(sttmt, id)

	var updatedAt pq.NullTime
	var deletedAt pq.NullTime

	switch err := row.Scan(
		&person.Id,
		&person.Fullname,
		&person.Age,
		&person.Email,
		&person.CreatedAt,
		&updatedAt,
		&deletedAt,
	); err {
	case sql.ErrNoRows:
		return &models.Person{}
	case nil:
		person.UpdatedAt = nil
		person.DeletedAt = nil

		if updatedAt.Valid {
			person.UpdatedAt = &updatedAt.Time
		}

		if deletedAt.Valid {
			person.DeletedAt = &deletedAt.Time
		}

		return &person
	default:
		log.Println(err.Error())
		return nil
	}
}
