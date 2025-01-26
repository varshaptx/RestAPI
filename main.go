package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type person struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender string `json:"gender"`
}

func getPerson(context *gin.Context) {

	p := person{
		Name:   "Bob Jones",
		Age:    26,
		Gender: "Male",
	}

	context.JSON(http.StatusOK, p)

}

func main() {
	router := gin.Default()
	router.GET("/person", getPerson)
	router.Run("localhost:8090")
}
