package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/bineyond/Go-000/ßWeek02/service"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.GET("user/:id", func(c *gin.Context) {

		name, err := service.GetUserByID(c.Param("id"))
		if err != nil {
			log.Printf("error occurred: %v\n", errors.Unwrap(err))
			log.Printf("stack trace:\n%+v", err)
		}
		c.String(http.StatusOK, "hello %s", name)
	})
	router.Run(":8080")
}
