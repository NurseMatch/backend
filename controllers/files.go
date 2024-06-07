package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterFileEndpoints(e *gin.Engine) {
	e.POST("/files", uploadFiles)
}

func uploadFiles(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["files[]"]

	for _, file := range files {
		err := c.SaveUploadedFile(file, "./tmp/"+file.Filename)
		if err != nil {
			return
		}
	}

	c.JSON(http.StatusCreated, gin.H{
		"files uploaded": len(files),
		"message":        "Files uploaded successfully",
	})
}
