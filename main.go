package main

import (
	"backend/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"os"
)

func main() {
	local := os.Getenv("LOCALDB") == "true"

	var db *gorm.DB
	var err error

	if local {
		db, err = connectToLocalDb()
	} else {
		db, err = connectToDb()
	}
	runMigration(db)
	err = setupApi(db)
	if err != nil {
		return
	}
}

func setupApi(db *gorm.DB) error {
	c := gin.Default()

	c.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	controllers.RegisterAssignmentEndpoints(c)
	controllers.RegisterAccountEndpoints(c)

	err := c.Run()
	return err
}
