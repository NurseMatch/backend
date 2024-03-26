package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := connectToDb()
	runMigration(db)
	err = setupApi()
	if err != nil {
		return
	}
}

func setupApi() error {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	err := r.Run()
	return err
}
