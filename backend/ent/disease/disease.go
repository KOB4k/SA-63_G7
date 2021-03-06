// Code generated by entc, DO NOT EDIT.

package disease

const (
	// Label holds the string label denoting the disease type in the database.
	Label = "disease"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldSymptom holds the string denoting the symptom field in the database.
	FieldSymptom = "symptom"
	// FieldContagion holds the string denoting the contagion field in the database.
	FieldContagion = "contagion"

	// EdgeEmployee holds the string denoting the employee edge name in mutations.
	EdgeEmployee = "employee"
	// EdgeSeverity holds the string denoting the severity edge name in mutations.
	EdgeSeverity = "severity"
	// EdgeDiseasetype holds the string denoting the diseasetype edge name in mutations.
	EdgeDiseasetype = "diseasetype"

	// Table holds the table name of the disease in the database.
	Table = "diseases"
	// EmployeeTable is the table the holds the employee relation/edge.
	EmployeeTable = "diseases"
	// EmployeeInverseTable is the table name for the Employee entity.
	// It exists in this package in order to avoid circular dependency with the "employee" package.
	EmployeeInverseTable = "employees"
	// EmployeeColumn is the table column denoting the employee relation/edge.
	EmployeeColumn = "employee_disease"
	// SeverityTable is the table the holds the severity relation/edge.
	SeverityTable = "diseases"
	// SeverityInverseTable is the table name for the Severity entity.
	// It exists in this package in order to avoid circular dependency with the "severity" package.
	SeverityInverseTable = "severities"
	// SeverityColumn is the table column denoting the severity relation/edge.
	SeverityColumn = "severity_disease"
	// DiseasetypeTable is the table the holds the diseasetype relation/edge.
	DiseasetypeTable = "diseases"
	// DiseasetypeInverseTable is the table name for the Diseasetype entity.
	// It exists in this package in order to avoid circular dependency with the "diseasetype" package.
	DiseasetypeInverseTable = "diseasetypes"
	// DiseasetypeColumn is the table column denoting the diseasetype relation/edge.
	DiseasetypeColumn = "diseasetype_disease"
)

// Columns holds all SQL columns for disease fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldSymptom,
	FieldContagion,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Disease type.
var ForeignKeys = []string{
	"diseasetype_disease",
	"employee_disease",
	"severity_disease",
}
