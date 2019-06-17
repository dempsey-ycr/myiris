package mongo

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	// https://github.com/mongodb/mongo-go-driver/
)

func init(){
	mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
}