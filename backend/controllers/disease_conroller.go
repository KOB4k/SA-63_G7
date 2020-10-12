package controllers

import (
	"fmt"
	"context"
	"strconv"

	"github.com/KOB4k/app/ent"
	"github.com/KOB4k/app/ent/disease"
	"github.com/gin-gonic/gin"
)

type DiseaseController struct {
	client *ent.Client
	router gin.IRouter
}

func (ctl *DiseaseController) GetDisease(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	u, err := ctl.client.Disease.
		Query().
		Where(disease.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}

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
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, diseases)
}

func NewDiseaseController(router gin.IRouter, client *ent.Client) *DiseaseController {
	uc := &DiseaseController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

func (ctl *DiseaseController) register() {
	diseases := ctl.router.Group("/diseases")

	diseases.GET("", ctl.ListDisease)

	diseases.GET(":id", ctl.GetDisease)
}

////

func (ctl *DiseaseController) CreateDisease(c *gin.Context) {
    obj := ent.Disease{}
    if err := c.ShouldBind(&obj); err != nil {
        c.JSON(400, gin.H{
            "error": "disease binding failed",
        })
        return
    }

    u, err := ctl.client.Disease.
        Create().
		SetName(obj.Name).
		SetSyptom(obj.Syptom).
		SetContagion(obj.Contagion).
        Save(context.Background())
    if err != nil {
        c.JSON(400, gin.H{
            "error": "saving failed",
        })
        return
    }

    c.JSON(200, u)
}

func (ctl *DiseaseController) DeleteDisease(c *gin.Context) {

    id, err := strconv.ParseInt(c.Param("id"), 10, 64)

    if err != nil {

        c.JSON(400, gin.H{

            "error": err.Error(),
        })

        return

    }

    err = ctl.client.Disease.
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
