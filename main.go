package main

import (
	"backend/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	db, err := connectToDb()
	runMigration(db)
	err = setupApi(db)
	if err != nil {
		return
	}
}

func setupApi(db *gorm.DB) error {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})
	controllers.RegisterAssignmentEndpoints(r)

	err := r.Run()
	return err
}
