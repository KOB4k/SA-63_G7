// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/KOB4k/app/ent/disease"
	"github.com/KOB4k/app/ent/diseasetype"
	"github.com/KOB4k/app/ent/employee"
	"github.com/KOB4k/app/ent/schema"
	"github.com/KOB4k/app/ent/severity"
)

// The init function reads all schema descriptors with runtime
// code (default values, validators or hooks) and stitches it
// to their package variables.
func init() {
	diseaseFields := schema.Disease{}.Fields()
	_ = diseaseFields
	// diseaseDescName is the schema descriptor for name field.
	diseaseDescName := diseaseFields[0].Descriptor()
	// disease.NameValidator is a validator for the "name" field. It is called by the builders before save.
	disease.NameValidator = diseaseDescName.Validators[0].(func(string) error)
	// diseaseDescSyptom is the schema descriptor for syptom field.
	diseaseDescSyptom := diseaseFields[1].Descriptor()
	// disease.SyptomValidator is a validator for the "syptom" field. It is called by the builders before save.
	disease.SyptomValidator = diseaseDescSyptom.Validators[0].(func(string) error)
	// diseaseDescContagion is the schema descriptor for contagion field.
	diseaseDescContagion := diseaseFields[2].Descriptor()
	// disease.ContagionValidator is a validator for the "contagion" field. It is called by the builders before save.
	disease.ContagionValidator = diseaseDescContagion.Validators[0].(func(string) error)
	diseasetypeFields := schema.DiseaseType{}.Fields()
	_ = diseasetypeFields
	// diseasetypeDescName is the schema descriptor for name field.
	diseasetypeDescName := diseasetypeFields[0].Descriptor()
	// diseasetype.NameValidator is a validator for the "name" field. It is called by the builders before save.
	diseasetype.NameValidator = diseasetypeDescName.Validators[0].(func(string) error)
	employeeFields := schema.Employee{}.Fields()
	_ = employeeFields
	// employeeDescName is the schema descriptor for name field.
	employeeDescName := employeeFields[0].Descriptor()
	// employee.NameValidator is a validator for the "name" field. It is called by the builders before save.
	employee.NameValidator = employeeDescName.Validators[0].(func(string) error)
	severityFields := schema.Severity{}.Fields()
	_ = severityFields
	// severityDescName is the schema descriptor for name field.
	severityDescName := severityFields[0].Descriptor()
	// severity.NameValidator is a validator for the "name" field. It is called by the builders before save.
	severity.NameValidator = severityDescName.Validators[0].(func(string) error)
}
