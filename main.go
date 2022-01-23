package main

import (
	"github.com/AndriiMaliuta/gin-mongo-people/repo"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	mongoRepo = repo.MongoRepo{}
)

func getPersons(c *gin.Context) {

	c.PureJSON(200, mongoRepo.GetPersons())
	//c.JSON(200, jsonP)
}

func getCars(c *gin.Context) {
	c.PureJSON(200, mongoRepo.GetCars())
}

func getPersonId(c *gin.Context) {
	id := c.Param("id")
	//c.JSON(http.StatusOK, gin.H{"people": mongoRepo.GetPersonById(id)})
	c.JSON(http.StatusOK, mongoRepo.GetPersonById(id))
}

func main() {

	r := gin.Default()
	r.GET("/persons", getPersons)
	r.GET("/persons/:id", getPersonId)
	r.GET("/cars", getCars)

	r.Run(":8082")
}
