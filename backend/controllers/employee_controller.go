package controllers

import (
	"fmt"
	"context"
	"strconv"

	"github.com/KOB4k/app/ent"
	"github.com/KOB4k/app/ent/employee"
	"github.com/gin-gonic/gin"
)

type EmployeeController struct {
	client *ent.Client
	router gin.IRouter
}

func (ctl *EmployeeController) GetEmployee(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	u, err := ctl.client.Employee.
		Query().
		Where(employee.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}

func (ctl *EmployeeController) ListEmployee(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {
			limit = int(limit64)
		}
	}

	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {
			offset = int(offset64)
		}
	}

	employees, err := ctl.client.Employee.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, employees)
}

func NewEmployeeController(router gin.IRouter, client *ent.Client) *EmployeeController {
	uc := &EmployeeController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

func (ctl *EmployeeController) register() {
	employees := ctl.router.Group("/employees")

	employees.GET("", ctl.ListEmployee)

	employees.GET(":id", ctl.GetEmployee)
}

/////
func (ctl *EmployeeController) CreateEmployee(c *gin.Context) {
    obj := ent.Employee{}
    if err := c.ShouldBind(&obj); err != nil {
        c.JSON(400, gin.H{
            "error": "Employee binding failed",
        })
        return
    }

    u, err := ctl.client.Employee.
        Create().
		SetName(obj.Name).
        Save(context.Background())
    if err != nil {
        c.JSON(400, gin.H{
            "error": "saving failed",
        })
        return
    }

    c.JSON(200, u)
}

func (ctl *EmployeeController) DeleteEmployee(c *gin.Context) {

    id, err := strconv.ParseInt(c.Param("id"), 10, 64)

    if err != nil {

        c.JSON(400, gin.H{

            "error": err.Error(),
        })

        return

    }

    err = ctl.client.Employee.
        DeleteOneID(int(id)).
        Exec(context.Background())

    if err != nil {

        c.JSON(404, gin.H{

            "error": err.Error(),
        })

        return

    }

    c.JSON(200, gin.H{"result": fmt.Sprintf("ok deleted %v", id)})

}