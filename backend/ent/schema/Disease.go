package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

type Disease struct {
	ent.Schema
}

//Fields of the Disease
func (Disease) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("syptom").NotEmpty(),
		field.String("contagion").NotEmpty(),
	}
}

// Edges of the Disease.
func (Disease) Edges() []ent.Edge {
	return []ent.Edge{

		edge.From("employee", Employee.Type).Ref("disease").Unique(),
		edge.From("serverity", Severity.Type).Ref("disease").Unique(),
		edge.From("diseasetype", DiseaseType.Type).Ref("disease").Unique(),
	}
}
