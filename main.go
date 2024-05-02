package main

import (
	"backend/controllers"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
	"net/http"
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
	// Create a new Gin router with default middleware
	r := gin.Default()

	r.Use(CORSMiddleware())
	// Middleware to inject the database instance into the Gin context
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	r.Use(jwtMiddleware())

	controllers.RegisterAccountEndpoints(r)

	// Run the Gin server
	err := r.Run(":8080")
	return err
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Authorization, Accept,X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func jwtMiddleware() gin.HandlerFunc {
	jwtSecret := []byte(os.Getenv("JWTSECRET"))

	return func(c *gin.Context) {
		if c.Request.URL.Path[:8] == "/account" {
			c.Next()
			return
		}
		tokenString, err := c.Cookie("token")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwtSecret, nil
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userID := uint(claims["user_id"].(float64))
			c.Set("user_id", userID)
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
	}
}
