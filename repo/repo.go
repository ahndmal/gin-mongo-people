package repo

import (
	"context"
	"github.com/AndriiMaliuta/gin-mongo-people/models"
	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"os"
)

type MongoRepo struct{}

func (m MongoRepo) GetPersons() []models.Person {
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

	return people
}
func (m MongoRepo) GetPersonById(id string) models.Person {
	mongoUrl := os.Getenv("MONGO_URL")

	ctx := context.Background()
	client, err := qmgo.NewClient(ctx, &qmgo.Config{Uri: mongoUrl})
	if err != nil {
		log.Panicln(err)
	}
	db := client.Database("people")
	coll := db.Collection("person")

	var person models.Person

	coll.Find(ctx, bson.M{"_id": id}).One(&person)

	return person
}
