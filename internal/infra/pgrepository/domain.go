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
	pp.Id = rows["id"].(int64)
	pp.Fullname = rows["fullname"].(string)
	pp.Age = rows["age"].(int64)
	pp.Email = rows["email"].(string)

	createdAt, _ := rows["created_at"].(time.Time)
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
