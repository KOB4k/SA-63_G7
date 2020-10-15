package controllers

import (
	"context"
	"strconv"

	"github.com/KOB4k/app/ent"
	"github.com/KOB4k/app/ent/diseasetype"
	"github.com/KOB4k/app/ent/employee"
	"github.com/KOB4k/app/ent/severity"
	"github.com/gin-gonic/gin"
)

// DiseaseController defines the struct for the disease controller
type DiseaseController struct {
	client *ent.Client
	router gin.IRouter
}

// Disease defines the struct for the disease controller
type Disease struct {
	Name        string
	Symptom     string
	Contagion   string
	Diseasetype int
	Employee    int
	Severity    int
}

// CreateDisease handles POST requests for adding disease  entities
// @Summary Create disease
// @Description Create disease
// @ID create-disease
// @Accept   json
// @Produce  json
// @Param disease body ent.Disease true "Disease entity"
// @Success 200 {object} ent.Disease
// @Failure 400 {object} gin.H.
// @Failure 500 {object} gin.H
// @Router /diseases [post]
func (ctl *DiseaseController) CreateDisease(c *gin.Context) {
	obj := Disease{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "disease binding failed",
		})
		return
	}

	dt, err := ctl.client.Diseasetype.
		Query().
		Where(diseasetype.IDEQ(int(obj.Diseasetype))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "diseasetype not found",
		})
		return
	}
	e, err := ctl.client.Employee.
		Query().
		Where(employee.IDEQ(int(obj.Employee))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "employee not found",
		})
		return
	}
	s, err := ctl.client.Severity.
		Query().
		Where(severity.IDEQ(int(obj.Severity))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "severity not found",
		})
		return
	}

	d, err := ctl.client.Disease.
		Create().
		SetName(obj.Name).
		SetSymptom(obj.Symptom).
		SetContagion(obj.Contagion).
		SetEmployee(e).
		SetDiseasetype(dt).
		SetServerity(s).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, d)
}

// ListDisease handles request to get a list of disease entities
// @Summary List disease entities
// @Description list disease entities
// @ID list-disease
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Disease
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /diseases [get]
func (ctl *DiseaseController) ListDisease(c *gin.Context) {
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
	diseases, err := ctl.client.Disease.
		Query().
		WithDiseasetype().
		WithServerity().
		WithEmployee().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, diseases)
}

// NewDiseaseController creates and registers handles for the disease controller
func NewDiseaseController(router gin.IRouter, client *ent.Client) *DiseaseController {
	dc := &DiseaseController{
		client: client,
		router: router,
	}
	dc.register()
	return dc
}

func (ctl *DiseaseController) register() {
	diseases := ctl.router.Group("/diseases")
	diseases.POST("", ctl.CreateDisease)
	diseases.GET("", ctl.ListDisease)

}
