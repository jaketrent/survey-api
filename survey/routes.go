package survey

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type ok struct {
	Data []*Survey `json:"data"`
}

type clienterr struct {
	Title  string `json:"title"`
	Status int    `json:"status"` // TODO: make this a union
}

type bad struct {
	Errors []clienterr `json:"errors"`
}

func listSurvey(c *gin.Context) {
	log.Print("Survey listing...")
	db, _ := c.MustGet("db").(*sql.DB)
	surveys, err := findAll(db)
	if err == nil {
		log.Printf("Survey list res ok... (count: %v)", len(surveys))
		c.JSON(http.StatusOK, ok{
			Data: surveys,
		})
	} else {
		log.Printf("Survey list res bad... (msg: %s)", err.Error())
		c.JSON(http.StatusInternalServerError, bad{
			Errors: []clienterr{{Title: "list surveys error: " + err.Error(), Status: http.StatusInternalServerError}},
		})
	}
}

func changeSurvey(c *gin.Context) {
	db, _ := c.MustGet("db").(*sql.DB)

	var id int
	var err error
	var survey *Survey
	err = c.BindJSON(&survey)

	if err != nil {
		c.JSON(http.StatusBadRequest, bad{
			Errors: []clienterr{{Title: "Bad survey: " + err.Error(), Status: http.StatusBadRequest}},
		})
		return
	}

	id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, bad{
			Errors: []clienterr{{Title: "Bad id: " + c.Param("id"), Status: http.StatusBadRequest}},
		})
		return
	}
	if survey.Id != id {
		c.JSON(http.StatusBadRequest, bad{
			Errors: []clienterr{{Title: "URL id must match survey id", Status: http.StatusBadRequest}},
		})
		return
	}

	survey, err = updateSurvey(db, survey)
	if err == nil {
		c.JSON(http.StatusOK, ok{
			Data: []*Survey{survey},
		})
	} else {
		c.JSON(http.StatusInternalServerError, bad{
			Errors: []clienterr{{Title: "Update survey error: " + err.Error(), Status: http.StatusInternalServerError}},
		})
	}
}

func createSurvey(c *gin.Context) {
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
			Data: []*Survey{survey},
		})
	} else {
		c.JSON(http.StatusInternalServerError, bad{
			Errors: []clienterr{{Title: "Create survey error: " + err.Error(), Status: http.StatusInternalServerError}},
		})
		fmt.Println("survey create db error", err)
	}
}

func destroySurvey(c *gin.Context) {
	db, _ := c.MustGet("db").(*sql.DB)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, bad{
			Errors: []clienterr{{Title: "Bad id: " + c.Param("id"), Status: http.StatusBadRequest}},
		})
		return
	}

	err = deleteSurvey(db, id)
	if err == nil {
		c.Status(http.StatusNoContent)
	} else {
		c.JSON(http.StatusInternalServerError, bad{
			Errors: []clienterr{{Title: "Delete survey error: " + err.Error(), Status: http.StatusInternalServerError}},
		})
	}

}

func Mount(router *gin.Engine) {
	router.GET("/api/v1/survey", listSurvey)
	router.POST("/api/v1/survey", createSurvey)
	router.PUT("/api/v1/survey/:id", changeSurvey)
	router.DELETE("/api/v1/survey/:id", destroySurvey)
}
