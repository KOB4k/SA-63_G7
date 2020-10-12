package controllers

import (
	"fmt"
	"context"
	"strconv"

	"github.com/KOB4k/app/ent"
	"github.com/KOB4k/app/ent/diseasetype"
	"github.com/gin-gonic/gin"
)

type DiseaseTypeController struct {
	client *ent.Client
	router gin.IRouter
}

func (ctl *DiseaseTypeController) GetDiseaseType(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	u, err := ctl.client.DiseaseType.
		Query().
		Where(diseasetype.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}

func (ctl *DiseaseTypeController) ListDiseaseType(c *gin.Context) {
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

	diseasetypes, err := ctl.client.DiseaseType.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, diseasetypes)
}

func NewDiseaseTypeController(router gin.IRouter, client *ent.Client) *DiseaseTypeController {
	uc := &DiseaseTypeController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

func (ctl *DiseaseTypeController) register() {
	diseasetypes := ctl.router.Group("/diseasetypes")

	diseasetypes.GET("", ctl.ListDiseaseType)

	diseasetypes.GET(":id", ctl.GetDiseaseType)
}

////

func (ctl *DiseaseTypeController) CreateDiseaseType(c *gin.Context) {
    obj := ent.DiseaseType{}
    if err := c.ShouldBind(&obj); err != nil {
        c.JSON(400, gin.H{
            "error": "DiseaseType binding failed",
        })
        return
    }

    u, err := ctl.client.DiseaseType.
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

func (ctl *DiseaseTypeController) DeleteDiseaseType(c *gin.Context) {

    id, err := strconv.ParseInt(c.Param("id"), 10, 64)

    if err != nil {

        c.JSON(400, gin.H{

            "error": err.Error(),
        })

        return

    }

    err = ctl.client.DiseaseType.
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