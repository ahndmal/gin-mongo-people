package main

import (
	"github.com/AndriiMaliuta/gin-mongo-people/repo"
	"github.com/gin-gonic/gin"
)

const (
//mongoUrl = os.Getenv("MONGO_URL")
// context
)

func getPersons(c *gin.Context) {
	mongoRepo := repo.MongoRepo{}
	//jsonP, err := json.Marshal(people)
	c.PureJSON(200, mongoRepo.GetPersons())
	//c.JSON(200, jsonP)
}

func main() {

	r := gin.Default()
	r.GET("/persons", getPersons)
	r.GET("/persons/:id", getPersonId)
	r.Run(":8082")
}

func getPersonId(c *gin.Context) {
	id := c.Param("id")
	mongoRepo := repo.MongoRepo{}
	c.JSON(200, gin.H{"people": mongoRepo.GetPersonById(id)})
}
