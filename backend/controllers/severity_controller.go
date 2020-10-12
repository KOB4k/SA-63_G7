package controllers

import (
	"fmt"
	"context"
	"strconv"

	"github.com/KOB4k/app/ent"
	"github.com/KOB4k/app/ent/severity"
	"github.com/gin-gonic/gin"
)

type SeverityController struct {
	client *ent.Client
	router gin.IRouter
}

func (ctl *SeverityController) GetSeverity(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	u, err := ctl.client.Severity.
		Query().
		Where(severity.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}

func (ctl *SeverityController) ListSeverity(c *gin.Context) {
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

	severitys, err := ctl.client.Severity.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, severitys)
}

func NewSeverityController(router gin.IRouter, client *ent.Client) *SeverityController {
	uc := &SeverityController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

func (ctl *SeverityController) register() {
	severitys := ctl.router.Group("/severitys")

	severitys.GET("", ctl.ListSeverity)

	severitys.GET(":id", ctl.GetSeverity)
}


func (ctl *SeverityController) CreateSeverity(c *gin.Context) {
    obj := ent.Severity{}
    if err := c.ShouldBind(&obj); err != nil {
        c.JSON(400, gin.H{
            "error": "Severity binding failed",
        })
        return
    }

    u, err := ctl.client.Severity.
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

func (ctl *SeverityController) DeleteSeverity(c *gin.Context) {

    id, err := strconv.ParseInt(c.Param("id"), 10, 64)

    if err != nil {

        c.JSON(400, gin.H{

            "error": err.Error(),
        })

        return

    }

    err = ctl.client.Severity.
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