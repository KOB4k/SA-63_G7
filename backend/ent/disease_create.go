// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/KOB4k/app/ent/disease"
	"github.com/KOB4k/app/ent/diseasetype"
	"github.com/KOB4k/app/ent/employee"
	"github.com/KOB4k/app/ent/severity"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// DiseaseCreate is the builder for creating a Disease entity.
type DiseaseCreate struct {
	config
	mutation *DiseaseMutation
	hooks    []Hook
}

// SetName sets the Name field.
func (dc *DiseaseCreate) SetName(s string) *DiseaseCreate {
	dc.mutation.SetName(s)
	return dc
}

// SetSymptom sets the Symptom field.
func (dc *DiseaseCreate) SetSymptom(s string) *DiseaseCreate {
	dc.mutation.SetSymptom(s)
	return dc
}

// SetContagion sets the Contagion field.
func (dc *DiseaseCreate) SetContagion(s string) *DiseaseCreate {
	dc.mutation.SetContagion(s)
	return dc
}

// SetEmployeeID sets the employee edge to Employee by id.
func (dc *DiseaseCreate) SetEmployeeID(id int) *DiseaseCreate {
	dc.mutation.SetEmployeeID(id)
	return dc
}

// SetNillableEmployeeID sets the employee edge to Employee by id if the given value is not nil.
func (dc *DiseaseCreate) SetNillableEmployeeID(id *int) *DiseaseCreate {
	if id != nil {
		dc = dc.SetEmployeeID(*id)
	}
	return dc
}

// SetEmployee sets the employee edge to Employee.
func (dc *DiseaseCreate) SetEmployee(e *Employee) *DiseaseCreate {
	return dc.SetEmployeeID(e.ID)
}

// SetServerityID sets the serverity edge to Severity by id.
func (dc *DiseaseCreate) SetServerityID(id int) *DiseaseCreate {
	dc.mutation.SetServerityID(id)
	return dc
}

// SetNillableServerityID sets the serverity edge to Severity by id if the given value is not nil.
func (dc *DiseaseCreate) SetNillableServerityID(id *int) *DiseaseCreate {
	if id != nil {
		dc = dc.SetServerityID(*id)
	}
	return dc
}

// SetServerity sets the serverity edge to Severity.
func (dc *DiseaseCreate) SetServerity(s *Severity) *DiseaseCreate {
	return dc.SetServerityID(s.ID)
}

// SetDiseasetypeID sets the diseasetype edge to Diseasetype by id.
func (dc *DiseaseCreate) SetDiseasetypeID(id int) *DiseaseCreate {
	dc.mutation.SetDiseasetypeID(id)
	return dc
}

// SetNillableDiseasetypeID sets the diseasetype edge to Diseasetype by id if the given value is not nil.
func (dc *DiseaseCreate) SetNillableDiseasetypeID(id *int) *DiseaseCreate {
	if id != nil {
		dc = dc.SetDiseasetypeID(*id)
	}
	return dc
}

// SetDiseasetype sets the diseasetype edge to Diseasetype.
func (dc *DiseaseCreate) SetDiseasetype(d *Diseasetype) *DiseaseCreate {
	return dc.SetDiseasetypeID(d.ID)
}

// Mutation returns the DiseaseMutation object of the builder.
func (dc *DiseaseCreate) Mutation() *DiseaseMutation {
	return dc.mutation
}

// Save creates the Disease in the database.
func (dc *DiseaseCreate) Save(ctx context.Context) (*Disease, error) {
	if _, ok := dc.mutation.Name(); !ok {
		return nil, &ValidationError{Name: "Name", err: errors.New("ent: missing required field \"Name\"")}
	}
	if _, ok := dc.mutation.Symptom(); !ok {
		return nil, &ValidationError{Name: "Symptom", err: errors.New("ent: missing required field \"Symptom\"")}
	}
	if _, ok := dc.mutation.Contagion(); !ok {
		return nil, &ValidationError{Name: "Contagion", err: errors.New("ent: missing required field \"Contagion\"")}
	}
	var (
		err  error
		node *Disease
	)
	if len(dc.hooks) == 0 {
		node, err = dc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DiseaseMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			dc.mutation = mutation
			node, err = dc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(dc.hooks) - 1; i >= 0; i-- {
			mut = dc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DiseaseCreate) SaveX(ctx context.Context) *Disease {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (dc *DiseaseCreate) sqlSave(ctx context.Context) (*Disease, error) {
	d, _spec := dc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	d.ID = int(id)
	return d, nil
}

func (dc *DiseaseCreate) createSpec() (*Disease, *sqlgraph.CreateSpec) {
	var (
		d     = &Disease{config: dc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: disease.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: disease.FieldID,
			},
		}
	)
	if value, ok := dc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: disease.FieldName,
		})
		d.Name = value
	}
	if value, ok := dc.mutation.Symptom(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: disease.FieldSymptom,
		})
		d.Symptom = value
	}
	if value, ok := dc.mutation.Contagion(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: disease.FieldContagion,
		})
		d.Contagion = value
	}
	if nodes := dc.mutation.EmployeeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   disease.EmployeeTable,
			Columns: []string{disease.EmployeeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: employee.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dc.mutation.ServerityIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   disease.ServerityTable,
			Columns: []string{disease.ServerityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: severity.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dc.mutation.DiseasetypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   disease.DiseasetypeTable,
			Columns: []string{disease.DiseasetypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: diseasetype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return d, _spec
}
