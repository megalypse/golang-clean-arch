package pgrepository

import (
	"database/sql"
	"fmt"
	"log"
	"math"
	"time"

	"github.com/megalypse/golang-clean-arch/internal/config"
	"github.com/megalypse/golang-clean-arch/internal/domain/models"
)

// TODO: better connection management per request
type PgPersonRepository struct{}

func (PgPersonRepository) CreatePerson(person models.Person) int64 {
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
	var newRowId int64
	err := db.QueryRow(
		sttmt, person.Fullname, person.Age, person.Email, time.Now(), nil, nil,
	).Scan(&newRowId)
	if err != nil {
		panic(err)
	}

	return newRowId
}

func (PgPersonRepository) GetPersonById(id int64) *models.Person {
	db := config.GetPgDbConnection()
	defer db.Close()

	sttmt := `
	SELECT * FROM people p
	WHERE p.id = $1
	`

	row, err := db.Query(sttmt, id)
	if err != nil {
		log.Fatal(err.Error())
	}
	person, _ := getPeopleFromRows(row)
	return &person[0]
}

func (PgPersonRepository) Update(updatedPerson *models.Person) int64 {
	db := config.GetPgDbConnection()
	defer db.Close()

	now := time.Now()

	query := `
	UPDATE people
	SET fullname = $1,
	age = $2,
	email = $3
	`

	argCounter := 3
	queryArgs := []any{
		updatedPerson.Fullname,
		updatedPerson.Age,
		updatedPerson.Email,
	}

	if updatedPerson.DeletedAt != nil {
		argCounter++
		query += setArg(",\ndeleted_at = $", argCounter)
		queryArgs = append(queryArgs, now)
	}

	argCounter++
	query += setArg(",\nupdated_at = $", argCounter)
	queryArgs = append(queryArgs, now)

	argCounter++
	query += setArg("\nWHERE id = $", argCounter)
	queryArgs = append(queryArgs, updatedPerson.Id)

	log.Println(query)
	_, err := db.Query(query, queryArgs...)
	if err != nil {
		log.Fatal(err.Error())
	}

	return *updatedPerson.Id
}

func (PgPersonRepository) Filter(baseFilter *models.BaseFilter[models.Person]) models.Paginated[models.Person] {
	filters := baseFilter.Filters
	db := config.GetPgDbConnection()
	defer db.Close()

	query := `
	SELECT COUNT(*) OVER() as count, * FROM people p
	WHERE p.id IS NOT NULL
	`

	queryArgs := []any{}
	var argCounter int

	if filters.Id != nil && *filters.Id != 0 {
		argCounter++

		query += setArg("\nAND p.id = $", argCounter)
		queryArgs = append(queryArgs, filters.Id)
	}

	if filters.Fullname != nil && *filters.Fullname != "" {
		argCounter++

		query += setArg("\nAND p.fullname ILIKE $", argCounter)
		queryArgs = append(queryArgs, "%"+(*filters.Fullname)+"%")
	}

	if filters.Age != nil && *filters.Age != 0 {
		argCounter++

		query += setArg("\nAND p.age = $", argCounter)
		queryArgs = append(queryArgs, filters.Age)
	}

	if filters.Email != nil && *filters.Email != "" {
		argCounter++

		query += setArg("\nAND p.email = $", 4)
		queryArgs = append(queryArgs, filters.Email)
	}

	if filters.CreatedAt != nil {
		argCounter++

		query += setArg("\nAND p.created_at = $", argCounter)
		queryArgs = append(queryArgs, filters.CreatedAt)
	}

	if filters.UpdatedAt != nil {
		argCounter++

		query += setArg("\nAND p.updated_at = $", argCounter)
		queryArgs = append(queryArgs, filters.UpdatedAt)
	}

	if filters.DeletedAt != nil {
		argCounter++

		query += setArg("\nAND p.deleted_at = $", argCounter)
		queryArgs = append(queryArgs, filters.DeletedAt)
	}

	limit := 20
	if baseFilter.Limit != 0 {
		limit = baseFilter.Limit
	}

	offset := 1
	if baseFilter.Page > 1 {
		offset = baseFilter.Page
	}

	argCounter++
	query += setArg("\nLIMIT $", argCounter)
	queryArgs = append(queryArgs, limit)

	argCounter++
	query += setArg("\nOFFSET $", argCounter)
	queryArgs = append(queryArgs, (offset-1)*limit)

	rows, err := db.Query(query, queryArgs...)

	if err != nil {
		panic(err.Error())
	}

	people, rawRows := getPeopleFromRows(rows)

	if len(rawRows) > 0 {
		count := rawRows[0]["count"].(int64)
		parsedLimit := int64(limit)
		paginated := models.Paginated[models.Person]{
			Content:    people,
			Total:      count,
			Page:       int64(offset),
			Limit:      parsedLimit,
			TotalPages: int64(math.Ceil(float64(count) / float64(parsedLimit))),
		}

		return paginated
	}

	return models.Paginated[models.Person]{
		Content: []models.Person{},
	}
}

func setArg(sttmt string, arg int) string {
	return sttmt + fmt.Sprint(arg)
}

func getPeopleFromRows(row *sql.Rows) ([]models.Person, []map[string]any) {
	cols, _ := row.Columns()
	results := make([]models.Person, 0)
	rawMaps := make([]map[string]any, 0)

	for row.Next() {
		columns := make([]any, len(cols))
		columnPointers := make([]any, len(columns))

		for i := range columns {
			columnPointers[i] = &columns[i]
		}

		if err := row.Scan(columnPointers...); err != nil {
			log.Fatal(err.Error())
		}

		m := make(map[string]any)
		for i, colName := range cols {
			val := columnPointers[i].(*any)
			m[colName] = *val
		}

		rawMaps = append(rawMaps, m)
		pgPerson := PgPerson{}
		pgPerson.Map(m)

		person := pgPerson.ToPerson()
		results = append(results, person)
	}

	return results, rawMaps
}
