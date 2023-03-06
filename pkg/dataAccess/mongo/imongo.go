package mongodb

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"npm/pkg/models"
)

type DbMethods interface {
	FindOne(interface{}) *mongo.SingleResult
	FindAll(interface{}, *options.FindOptions) (*mongo.Cursor, error)
	InsertOne(user models.User) (*mongo.InsertOneResult, error)
	UpdateOne() error
	TotalDocument() (int64, error)
}
