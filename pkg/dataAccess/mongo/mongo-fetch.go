package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"npm/pkg/models"
)

type Collection struct {
	Users *mongo.Collection
}

//func (db *Collection) CreateMongoCollection(nameOfCollection string) {
//	db.LeadCollection = config.Db.Collection(nameOfCollection) //creating the collection
//}

func (db *Collection) FindOne(filter interface{}) *mongo.SingleResult {
	return db.Users.FindOne(context.Background(), filter)
}

func (db *Collection) FindAll(filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error) {
	return db.Users.Find(context.Background(), filter)
}

func (db *Collection) InsertOne(reqBody models.User) (*mongo.InsertOneResult, error) {
	return db.Users.InsertOne(context.Background(), reqBody)
}

func (db *Collection) UpdateOne(filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
	return db.Users.UpdateOne(context.Background(), filter, update)
}

func (db *Collection) TotalDocument() (int, error) {
	if db.Users == nil {
		fmt.Println("the collection is not initialise")
	}
	len, err := db.Users.EstimatedDocumentCount(context.Background())
	ans := (int)(len)
	return ans, err
}
