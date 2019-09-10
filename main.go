package main

import (
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v9"
)

func main() {
	binding.Validator = new(defaultValidator)

	r := gin.Default()

	r.GET("/car", func(c *gin.Context) {
		var query struct {
			Name  string `form:"name" binding:"required"`
			Color string `form:"color" binding:"required,oneof=blue yellow"`
		}

		if err := c.ShouldBind(&query); err != nil {
			spew.Dump(err)
			for _, fieldErr := range err.(validator.ValidationErrors) {
				c.JSON(http.StatusBadRequest, fieldError{fieldErr}.String())
				return // exit on first error
			}
		}

		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	r.Run()
}
