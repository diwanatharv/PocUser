package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"npm/pkg/config"
	"npm/pkg/models"
)

type collection struct {
	users *mongo.Collection
}

func Newmongof() *collection {
	return &collection{
		users: config.Db.Collection("users"),
	}
}
func (db *collection) FindOne(filter interface{}) *mongo.SingleResult {

	return db.users.FindOne(context.Background(), filter)
}

func (db *collection) FindAll(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error) {
	return db.users.Find(ctx, filter)
}

func (db *collection) InsertOne(reqBody models.User) (*mongo.InsertOneResult, error) {
	return db.users.InsertOne(context.Background(), reqBody)
}

func (db *collection) UpdateOne(filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
	return db.users.UpdateOne(context.Background(), filter, update)
}

func (db *collection) TotalDocument() (int64, error) {
	if db.users == nil {
		fmt.Println("there is no document here")
	}
	len, err := db.users.EstimatedDocumentCount(context.Background())
	return len, err
}
