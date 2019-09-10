package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v9"
)

func main() {
	binding.Validator = new(defaultValidator)

	r := gin.Default()

	r.GET("/car", func(c *gin.Context) {
		var query struct {
			ID    string `json"-"`
			Name  string `json:"name" binding:"required"`
			Color string `json:"color" binding:"oneof=red blue`
		}

		if err := c.ShouldBind(&query); err != nil {
			for _, fieldErr := range err.(validator.ValidationErrors) {
				c.JSON(http.StatusBadRequest, fmt.Sprint(fieldErr))
				return // exit on first error
			}
		}

		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	r.Run()
}
