package mongodb

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"

	"npm/pkg/models"
)

// functions in mongofetch
type DbMethods interface {
	FindOne(interface{}) *mongo.SingleResult
	FindAll(context.Context, interface{}, ...*options.FindOptions) (*mongo.Cursor, error)
	InsertOne(user models.User) (*mongo.InsertOneResult, error)
	UpdateOne(interface{}, interface{}) (*mongo.UpdateResult, error)
	TotalDocument() (int64, error)
}
