package survey

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ok struct {
	Data *Survey `json:"data"`
}

type clienterr struct {
	Title  string `json:"title"`
	Status int    `json:"status"` // TODO: make this a union
}

type bad struct {
	Errors []clienterr `json:"errors"`
}

func create(c *gin.Context) {
	db, _ := c.MustGet("db").(*sql.DB)

	var err error
	var survey *Survey
	err = c.BindJSON(&survey)

	if err != nil {
		c.JSON(http.StatusBadRequest, bad{
			Errors: []clienterr{{Title: "Bad survey", Status: http.StatusBadRequest}},
		})
		fmt.Println("survey create req error", err)
		return
	}
	survey, err = insertSurvey(db, survey)

	if err == nil {
		c.JSON(http.StatusCreated, ok{
			Data: survey,
		})
	} else {
		c.JSON(http.StatusInternalServerError, err)
		fmt.Println("survey create db error", err)
	}
}

func Mount(router *gin.Engine) {
	router.POST("/api/v1/survey", create)
}
