package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

type DiseaseType struct {
	ent.Schema
}

//Fields of the DiseaseType
func (DiseaseType) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
	}
}

// Edges of the DiseaseType.
func (DiseaseType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("disease", Disease.Type),
	}
}
