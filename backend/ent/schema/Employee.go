package schema

import (


    "github.com/facebookincubator/ent"
    "github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"
	
)

// Employee schema
type Employee struct {
	ent.Schema
}

//Fields of the employee
func (Employee) Fields() []ent.Field {
	return []ent.Field{
		field.String("User_id"),

	}
}

// Edges of the employee.
func (Employee) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("disease", Disease.Type),

	}
}
