package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Person struct {
	Id          primitive.ObjectID `json:"_id" bson:"_id"`
	IdN         int                `json:"id" bson:"id"`
	Name        string             `json:"name" bson:"name"`
	FullName    string             `json:"fullName" bson:"fullName"`
	Email       string             `json:"email" bson:"email"`
	Gender      string             `json:"gender" bson:"gender"`
	Status      string             `json:"status" bson:"status"`
	Age         int                `json:"age" bson:"age"`
	HasChildren bool               `json:"hasChildren"`
	Engaged     bool               `json:"engaged"`
	CreatedAt   time.Time          `json:"createdAt"`
	CountryId   string             `json:"countryId" bson:"countryId"`
}

type Car struct {
	Id           primitive.ObjectID `json:"_id" bson:"_id"`
	CarId        int                `json:"carId" bson:"carId"`
	MaxSpeed     int                `json:"maxSpeed" bson:"maxSpeed"`
	Model        string             `json:"model"`
	Color        string             `json:"color"`
	Age          int                `json:"age" bson:"age"`
	Origin       string             `json:"origin"`
	CountryCodes string             `json:"countryCodes"`
	CountryId    string             `json:"countryId" bson:"countryId"`
	PersonId     string             `json:"personId" bson:"personId"`
}

type Cat struct {
	Id           primitive.ObjectID `json:"_id" bson:"_id"`
	CatId        int                `json:"id" bson:"id"`
	Name         int                `json:"name" bson:"name"`
	Breed        string             `json:"breed"`
	Color        string             `json:"color"`
	Age          int                `json:"age" bson:"age"`
	Adaptability int                `json:"adaptability" bson:"adaptability"`
	Intelligence int                `json:"intelligence" bson:"intelligence"`
	Hairless     int                `json:"hairless" bson:"hairless"`
	Registry     string             `json:"registry"`
	DogFriendly  bool               `json:"dogFriendly"`
	PersonId     string             `json:"personId" bson:"personId"`
}
