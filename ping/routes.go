package ping

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ok struct {
	Ok bool `json:"ok"`
}

type bad struct {
	ErrMsg string `json:"errMsg"`
}

func ping(c *gin.Context) {
	db, _ := c.MustGet("db").(*sql.DB)

	err := db.Ping()
	if err == nil {
		c.JSON(http.StatusOK, ok{Ok: true})
	} else {
		c.JSON(http.StatusInternalServerError, bad{ErrMsg: err.Error()})
	}
}

func Mount(router *gin.Engine) {
	router.GET("/ping", ping)
}
