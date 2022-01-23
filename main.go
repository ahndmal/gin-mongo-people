package main

import (
	"context"
	"encoding/json"
	"github.com/AndriiMaliuta/gin-mongo-people/models"
	"github.com/gin-gonic/gin"
	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"os"
)

func getPersons(c *gin.Context) {
	mongoUrl := os.Getenv("MONGO_URL")

	ctx := context.Background()
	client, err := qmgo.NewClient(ctx, &qmgo.Config{Uri: mongoUrl})
	if err != nil {
		log.Panicln(err)
	}
	db := client.Database("people")
	coll := db.Collection("person")

	people := make([]models.Person, 0)

	coll.Find(ctx, bson.M{"age": 36}).All(&people)

	jsonP, err := json.Marshal(people)
	c.JSON(200, gin.H{"people": jsonP})
}

func main() {

	r := gin.Default()
	r.GET("/", getPersons)
	r.Run(":8082")
}
