package pgrepository

import (
	"time"

	"github.com/megalypse/golang-clean-arch/internal/domain/models"
)

type RowMapper interface {
	Map(map[string]any)
}

type PgPerson models.Person

func (pp *PgPerson) Map(rows map[string]any) {
	id := rows["id"].(int64)
	fullname := rows["fullname"].(string)
	age := rows["age"].(int64)
	email := rows["email"].(string)
	createdAt := rows["created_at"].(time.Time)

	pp.Id = &id
	pp.Fullname = &fullname
	pp.Age = &age
	pp.Email = &email
	pp.CreatedAt = &createdAt

	updatedAt, ok := rows["updated_at"].(time.Time)
	if ok {
		pp.UpdatedAt = &updatedAt
	}

	deletedAt, ok := rows["deleted_at"].(time.Time)
	if ok {
		pp.DeletedAt = &deletedAt
	}
}

func (pp PgPerson) ToPerson() models.Person {
	return models.Person(pp)
}
