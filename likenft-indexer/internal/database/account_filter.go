package database

import (
	"likenft-indexer/ent"

	"entgo.io/ent/dialect/sql"
)

type AccountPagination struct {
	Limit   *int
	Key     *int
	Reverse *bool
}

func (f *AccountPagination) HandlePagination(q *ent.AccountQuery) *ent.AccountQuery {
	if f.Reverse != nil && *f.Reverse {
		q = q.Order(sql.OrderByField("id", sql.OrderDesc()).ToFunc())
	} else {
		q = q.Order(sql.OrderByField("id", sql.OrderAsc()).ToFunc())
	}

	if f.Limit != nil {
		q = q.Limit(*f.Limit)
	} else {
		q = q.Limit(100)
	}

	if f.Key != nil {
		if f.Reverse != nil && *f.Reverse {
			q = q.Where(sql.FieldLT("id", *f.Key))
		} else {
			q = q.Where(sql.FieldGT("id", *f.Key))
		}
	}

	return q
}
