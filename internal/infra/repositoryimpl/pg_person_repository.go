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

	row := db.QueryRow(sttmt, id)
	return getPersonFromRow(row)
}

// TODO: Add pagination
func (rep PgPersonRepository) Filter(filters models.Person) []models.Person {
	db := config.GetPgDbConnection()
	defer db.Close()

	query := `
	SELECT * FROM people p
	WHERE p.id IS NOT NULL
	`

	queryArgs := []any{}

	if filters.Id != 0 {
		query += "\nAND p.id = $1"
		queryArgs = append(queryArgs, filters.Id)
	}

	if filters.Fullname != "" {
		query += "\nAND p.fullname = $2"
		queryArgs = append(queryArgs, filters.Fullname)
	}

	if filters.Age != 0 {
		query += "\nAND p.age = $3"
		queryArgs = append(queryArgs, filters.Age)
	}

	if filters.Email != "" {
		query += "\nAND p.email = $4"
		queryArgs = append(queryArgs, filters.Email)
	}

	if filters.CreatedAt != nil {
		query += "\nAND p.created_at = $5"
		queryArgs = append(queryArgs, filters.CreatedAt)
	}

	if filters.UpdatedAt != nil {
		query += "\nAND p.updated_at = $6"
		queryArgs = append(queryArgs, filters.UpdatedAt)
	}

	if filters.DeletedAt != nil {
		query += "\nAND p.deleted_at = $7"
		queryArgs = append(queryArgs, filters.DeletedAt)
	}

	rows, err := db.Query(query, queryArgs...)

	if err != nil {
		panic(err.Error())
	}

	result := []models.Person{}
	for rows.Next() {
		person := getPersonFromRows(rows)
		result = append(result, *person)
	}

	return result
}

func getPersonFromRows(row *sql.Rows) *models.Person {
	person := models.Person{}

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

func getPersonFromRow(row *sql.Row) *models.Person {
	person := models.Person{}

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
