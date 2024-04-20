package controllers

import (
	"backend/data"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
)

func RegisterAccountEndpoints(e *gin.Engine) {
	e.POST("/createUser", createUser)
	e.POST("/login", login)
}

func createUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var user data.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
		})
		return
	}

	// Check if the username or email already exists
	var existingUser data.User
	if err := db.Where("username = ?", user.Username).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": "Username or email already exists",
		})
		return
	}

	hashedPassword, err := HashPassword(user.HashedPassword)
	if err != nil {
		panic("failed to hash password")
	}
	user.HashedPassword = hashedPassword

	// Create the user
	if err := db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"user_id": user.ID,
		"message": "User created successfully",
	})
}

func login(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var user data.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
		})
		return
	}
	var unHashedPassword = user.HashedPassword

	if err := db.Where(&data.User{Username: user.Username}).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid username or password",
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(unHashedPassword)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid username or password",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"user_id": user.ID,
	})
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
