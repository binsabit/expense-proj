package mongoDB

import "go.mongodb.org/mongo-driver/mongo"

type ModelDB struct {
	DB *mongo.Database
}
