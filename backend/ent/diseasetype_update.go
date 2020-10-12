// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/KOB4k/app/ent/disease"
	"github.com/KOB4k/app/ent/diseasetype"
	"github.com/KOB4k/app/ent/predicate"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// DiseaseTypeUpdate is the builder for updating DiseaseType entities.
type DiseaseTypeUpdate struct {
	config
	hooks      []Hook
	mutation   *DiseaseTypeMutation
	predicates []predicate.DiseaseType
}

// Where adds a new predicate for the builder.
func (dtu *DiseaseTypeUpdate) Where(ps ...predicate.DiseaseType) *DiseaseTypeUpdate {
	dtu.predicates = append(dtu.predicates, ps...)
	return dtu
}

// SetName sets the name field.
func (dtu *DiseaseTypeUpdate) SetName(s string) *DiseaseTypeUpdate {
	dtu.mutation.SetName(s)
	return dtu
}

// AddDiseaseIDs adds the disease edge to Disease by ids.
func (dtu *DiseaseTypeUpdate) AddDiseaseIDs(ids ...int) *DiseaseTypeUpdate {
	dtu.mutation.AddDiseaseIDs(ids...)
	return dtu
}

// AddDisease adds the disease edges to Disease.
func (dtu *DiseaseTypeUpdate) AddDisease(d ...*Disease) *DiseaseTypeUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return dtu.AddDiseaseIDs(ids...)
}

// Mutation returns the DiseaseTypeMutation object of the builder.
func (dtu *DiseaseTypeUpdate) Mutation() *DiseaseTypeMutation {
	return dtu.mutation
}

// RemoveDiseaseIDs removes the disease edge to Disease by ids.
func (dtu *DiseaseTypeUpdate) RemoveDiseaseIDs(ids ...int) *DiseaseTypeUpdate {
	dtu.mutation.RemoveDiseaseIDs(ids...)
	return dtu
}

// RemoveDisease removes disease edges to Disease.
func (dtu *DiseaseTypeUpdate) RemoveDisease(d ...*Disease) *DiseaseTypeUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return dtu.RemoveDiseaseIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (dtu *DiseaseTypeUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := dtu.mutation.Name(); ok {
		if err := diseasetype.NameValidator(v); err != nil {
			return 0, &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(dtu.hooks) == 0 {
		affected, err = dtu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DiseaseTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			dtu.mutation = mutation
			affected, err = dtu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(dtu.hooks) - 1; i >= 0; i-- {
			mut = dtu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dtu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (dtu *DiseaseTypeUpdate) SaveX(ctx context.Context) int {
	affected, err := dtu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (dtu *DiseaseTypeUpdate) Exec(ctx context.Context) error {
	_, err := dtu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dtu *DiseaseTypeUpdate) ExecX(ctx context.Context) {
	if err := dtu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (dtu *DiseaseTypeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   diseasetype.Table,
			Columns: diseasetype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: diseasetype.FieldID,
			},
		},
	}
	if ps := dtu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := dtu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: diseasetype.FieldName,
		})
	}
	if nodes := dtu.mutation.RemovedDiseaseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   diseasetype.DiseaseTable,
			Columns: []string{diseasetype.DiseaseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: disease.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dtu.mutation.DiseaseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   diseasetype.DiseaseTable,
			Columns: []string{diseasetype.DiseaseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: disease.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, dtu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{diseasetype.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// DiseaseTypeUpdateOne is the builder for updating a single DiseaseType entity.
type DiseaseTypeUpdateOne struct {
	config
	hooks    []Hook
	mutation *DiseaseTypeMutation
}

// SetName sets the name field.
func (dtuo *DiseaseTypeUpdateOne) SetName(s string) *DiseaseTypeUpdateOne {
	dtuo.mutation.SetName(s)
	return dtuo
}

// AddDiseaseIDs adds the disease edge to Disease by ids.
func (dtuo *DiseaseTypeUpdateOne) AddDiseaseIDs(ids ...int) *DiseaseTypeUpdateOne {
	dtuo.mutation.AddDiseaseIDs(ids...)
	return dtuo
}

// AddDisease adds the disease edges to Disease.
func (dtuo *DiseaseTypeUpdateOne) AddDisease(d ...*Disease) *DiseaseTypeUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return dtuo.AddDiseaseIDs(ids...)
}

// Mutation returns the DiseaseTypeMutation object of the builder.
func (dtuo *DiseaseTypeUpdateOne) Mutation() *DiseaseTypeMutation {
	return dtuo.mutation
}

// RemoveDiseaseIDs removes the disease edge to Disease by ids.
func (dtuo *DiseaseTypeUpdateOne) RemoveDiseaseIDs(ids ...int) *DiseaseTypeUpdateOne {
	dtuo.mutation.RemoveDiseaseIDs(ids...)
	return dtuo
}

// RemoveDisease removes disease edges to Disease.
func (dtuo *DiseaseTypeUpdateOne) RemoveDisease(d ...*Disease) *DiseaseTypeUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return dtuo.RemoveDiseaseIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (dtuo *DiseaseTypeUpdateOne) Save(ctx context.Context) (*DiseaseType, error) {
	if v, ok := dtuo.mutation.Name(); ok {
		if err := diseasetype.NameValidator(v); err != nil {
			return nil, &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}

	var (
		err  error
		node *DiseaseType
	)
	if len(dtuo.hooks) == 0 {
		node, err = dtuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DiseaseTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			dtuo.mutation = mutation
			node, err = dtuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(dtuo.hooks) - 1; i >= 0; i-- {
			mut = dtuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dtuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (dtuo *DiseaseTypeUpdateOne) SaveX(ctx context.Context) *DiseaseType {
	dt, err := dtuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return dt
}

// Exec executes the query on the entity.
func (dtuo *DiseaseTypeUpdateOne) Exec(ctx context.Context) error {
	_, err := dtuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dtuo *DiseaseTypeUpdateOne) ExecX(ctx context.Context) {
	if err := dtuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (dtuo *DiseaseTypeUpdateOne) sqlSave(ctx context.Context) (dt *DiseaseType, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   diseasetype.Table,
			Columns: diseasetype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: diseasetype.FieldID,
			},
		},
	}
	id, ok := dtuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing DiseaseType.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := dtuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: diseasetype.FieldName,
		})
	}
	if nodes := dtuo.mutation.RemovedDiseaseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   diseasetype.DiseaseTable,
			Columns: []string{diseasetype.DiseaseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: disease.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dtuo.mutation.DiseaseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   diseasetype.DiseaseTable,
			Columns: []string{diseasetype.DiseaseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: disease.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	dt = &DiseaseType{config: dtuo.config}
	_spec.Assign = dt.assignValues
	_spec.ScanValues = dt.scanValues()
	if err = sqlgraph.UpdateNode(ctx, dtuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{diseasetype.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return dt, nil
}
