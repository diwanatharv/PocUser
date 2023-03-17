package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"npm/pkg/dataAccess/mongodb"
	"npm/pkg/dataAccess/redis"
	"npm/pkg/models"
)

var cache redisdb.Redis

//	func InitialiseCollection() {
//		mongoCollection = mongodb.GetCollection()
//	}
func Create(reqBody models.User) (*mongo.InsertOneResult, error) {
	// unique id of the lead will be the length of the collection

	var db = mongodb.Newmongof()
	ans, err := db.TotalDocument()
	if err != nil {
		return nil, err
	}
	// users id equal to the lenght of the reqbody
	reqBody.Id = int(ans)

	// insert it into the mongodb
	return db.InsertOne(reqBody)
}

func FindAll(filter interface{}, opts ...*options.FindOptions) []models.User {
	//slice f the user
	var db = mongodb.Newmongof()
	var allUser []models.User
	var oneUser models.User
	//will find all the id through filter condition
	findElementRes, err := db.FindAll(context.Background(), filter)
	if err != nil {
		//return err
		echo.New().Logger.Fatal("error in mongo - fetch FindAll")
	}
	//iterating through the slice
	for findElementRes.Next(context.Background()) {
		err := findElementRes.Decode(&oneUser)
		if err != nil {
			fmt.Println(err)
		}
		allUser = append(allUser, oneUser)
	}
	return allUser
}
func FindOne(filter interface{}, key string) models.User {
	var db = mongodb.Newmongof()
	var ans models.User

	//checking the data in redisdb db
	val2, err := cache.Get(key).Result()

	//if we do not find it in the reddis
	if err == redis.Nil {
		// add the key
		// checking if the leadId exists in db or not
		var findOneUser models.User
		// finding in  the database
		err1 := db.FindOne(filter).Decode(&findOneUser)

		if err1 != nil {
			fmt.Println("error in FindOne service layer")
		}

		findOneUser2, _ := json.Marshal(findOneUser) // converting to byte for setting in the reddis
		key := strconv.Itoa(findOneUser.Id)
		// setting in the reddis
		err2 := cache.Set(key, findOneUser2, 0).Err()

		if err2 != nil {
			fmt.Println("not able to set the values in redisdb")
		}
		return findOneUser
	} else {

		// exists in redisdb db , we will get the data from redisDb and will unmarshal it and return it
		// this is done to beautify in showing in the postman
		err := json.Unmarshal([]byte(val2), &ans)
		if err != nil {
			fmt.Println("error in unmarshalling")
		}
		return ans
	}

}
func UpdateOne(reqBody models.User, key string) (*mongo.UpdateResult, error) {
	var db = mongodb.Newmongof()
	keyInt, _ := strconv.Atoi(key)
	oneUser := FindOne(bson.M{"id": keyInt}, key)
	var empty models.User
	if oneUser == empty {
		return nil, errors.New("users doesn't exists")
	}

	reqBody.Id = keyInt

	// fields which we want to update
	updateField := bson.M{"$set": bson.M{"firstname": reqBody.FirstName, "lastname": reqBody.LastName, "email": reqBody.Email, "businessType": reqBody.BusinessType, "phoneNo": reqBody.PhoneNo, "companyName": reqBody.CompanyName, "country": reqBody.Country}}
	//updateFileRes, err := collection.UpdateOne(bson.M{"unique_id": findOneLead.UniqueId}, updateField)
	updateFileRes, err := db.UpdateOne(bson.M{"id": keyInt}, updateField)

	if err != nil {
		fmt.Println("error in updateFileRes")
		return nil, errors.New("error in updateFileRes")
	}

	reqBodyv2, _ := json.Marshal(reqBody) // convert it to byte so we can store it in redisdb

	//updating in redisdb
	err1 := cache.Set(key, reqBodyv2, 0).Err()

	if err1 == redis.Nil {
		fmt.Println("not able to set the values in redisdb")
	}
	return updateFileRes, nil
}
