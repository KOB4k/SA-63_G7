package main

import (
	"context"
	"log"

	"github.com/KOB4k/app/controllers"
	_ "github.com/KOB4k/app/docs"
	"github.com/KOB4k/app/ent"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

// Employees  defines the struct for the employees
type Employees struct {
	Employee []Employee
}

// Employee  defines the struct for the employee
type Employee struct {
	UserID string
}

// Diseasetypes  defines the struct for the diseasetypes
type Diseasetypes struct {
	Diseasetype []Diseasetype
}

// Diseasetype  defines the struct for the diseasetype
type Diseasetype struct {
	Name string
}

// Severitys  defines the struct for the severitys
type Severitys struct {
	Severity []Severity
}

// Severity  defines the struct for the severity
type Severity struct {
	Name string
}

// @title SUT SA Example API
// @version 1.0
// @description This is a sample server for SUT SE 2563
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	client, err := ent.Open("sqlite3", "file:contagion.db?&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("fail to open sqlite3: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	v1 := router.Group("/api/v1")
	controllers.NewDiseaseController(v1, client)
	controllers.NewDiseasetypeController(v1, client)
	controllers.NewEmployeeController(v1, client)
	controllers.NewSeverityController(v1, client)

	// Set Diseasetypes Data
	diseasetypes := Diseasetypes{
		Diseasetype: []Diseasetype{
			Diseasetype{"โรคติดต่อ"},
			Diseasetype{"โรคติดต่อต้องแจ้งความ"},
			Diseasetype{"โรคติดต่ออันตราย"},
		},
	}

	for _, dt := range diseasetypes.Diseasetype {
		client.Diseasetype.
			Create().
			SetName(dt.Name).
			Save(context.Background())
	}

	// Set Employees Data
	employees := Employees{
		Employee: []Employee{
			Employee{"D12345"},
			Employee{"D54231"},
		},
	}

	for _, e := range employees.Employee {
		client.Employee.
			Create().
			SetUserID(e.UserID).
			Save(context.Background())
	}

	// Set Severity Data
	severitys := Severitys{
		Severity: []Severity{
			Severity{"เริ่มต้น"},
			Severity{"รุนแรง"},
			Severity{"รุนแรงมาก"},
		},
	}

	for _, s := range severitys.Severity {
		client.Severity.
			Create().
			SetName(s.Name).
			Save(context.Background())
	}

	//router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
