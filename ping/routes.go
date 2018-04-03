package ping

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type ok struct {
	Ok bool `json:"ok"`
}

type bad struct {
	ErrMsg string `json:"errMsg"`
}

func ping(c *gin.Context) {
	log.Print("Pinging...")
	db, _ := c.MustGet("db").(*sql.DB)
	log.Printf("Got db (db: %o)", db)

	err := db.Ping()
	log.Print("Post ping")
	if err == nil {
		log.Print("Ping success")
		c.JSON(http.StatusOK, ok{Ok: true})
	} else {
		log.Printf("Ping error (msg: %s)", err.Error())
		c.JSON(http.StatusInternalServerError, bad{ErrMsg: err.Error()})
	}
}

func Mount(router *gin.Engine) {
	router.GET("/ping", ping)
}
