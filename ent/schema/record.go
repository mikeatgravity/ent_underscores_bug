package schema

import (
	"enttest/lib_a_b"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Record struct {
	ent.Schema
}

func (Record) Fields() []ent.Field {
	return []ent.Field{
		field.JSON("data", lib_a_b.Record{}),
	}
}
