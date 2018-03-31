package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"jaketrent.com/survey-api/survey"
	"log"
	"os"
)

func hasDatabase(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}

func main() {
	connStr := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal("Db unable to connect", err)
	}
	defer db.Close()

	router := gin.Default()

	router.Use(hasDatabase(db))

	survey.Mount(router)

	router.Run()
}
