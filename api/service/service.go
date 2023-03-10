package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"npm/pkg/dataAccess"
	"npm/pkg/dataAccess/reddis"
	"npm/pkg/models"
)

// wrapper check ki rdb me hai nhi hai to mongodb ki method call kar

var cache redisdb.Redis

func Create(reqBody models.User) (*mongo.InsertOneResult, error) {
	// unique id of the lead will be the length of the collection
	ans, err := dataAccess.Collection.TotalDocument()
	if err != nil {
		return nil, err
	}
	reqBody.Id = ans

	return dataAccess.Collection.InsertOne(reqBody)
}

func FindAll(filter interface{}, opts ...*options.FindOptions) []models.User {
	var allUser []models.User
	var oneUser models.User
	findElementRes, err := dataAccess.Collection.FindAll(filter)
	if err != nil {
		//return err
		fmt.Println("error in mongo - fetch FindAll")
	}
	//data
	for findElementRes.Next(context.Background()) {
		err := findElementRes.Decode(&oneUser)
		if err != nil {
			fmt.Println(err)
		}
		allUser = append(allUser, oneUser)
	}
	return allUser
}

// getLead
func FindOne(filter interface{}, key string) models.User {
	var ans models.User

	//checking the data in redisdb db
	val2, err2 := cache.Get(key).Result()

	if err2 == redis.Nil {
		//add the key
		fmt.Println("FindOne , adding the key, in if")
		//checking if the leadId exists in db or not
		var findOneLead models.User
		keyInt, _ := strconv.Atoi(key)

		err2 := dataAccess.Collection.FindOne(bson.M{"unique_id": keyInt}).Decode(&findOneLead)

		if err2 != nil {
			fmt.Println("error in FindOne service layer")
		}

		findOneLeadv2, _ := json.Marshal(findOneLead) //converting to byte
		key := strconv.Itoa(findOneLead.Id)

		err4 := cache.Set(key, findOneLeadv2, 0).Err()

		if err4 != nil {
			fmt.Println("not able to set the values in redisdb")
		}
		return findOneLead
	} else {
		fmt.Println("in findOne in else")
		// exists in redisdb db , we will get the data from redisDb and will unmarshal it and return it

		err := json.Unmarshal([]byte(val2), &ans)
		if err != nil {
			fmt.Println("error in unmarshalling")
		}
		return ans
	}

}

// update lead
func UpdateOne(reqBody models.User, key string) (*mongo.UpdateResult, error) {

	oneLead := FindOne(bson.M{"unique_id": key}, key)
	var empty models.User
	if oneLead == empty {
		fmt.Println("Lead doesn't exists")
		return nil, errors.New("Lead doesn't exists")
	}
	keyInt, err := strconv.Atoi(key) //leadId should be same
	reqBody.Id = keyInt

	// fields which we want to update
	updateField := bson.M{"$set": bson.M{"first_name": reqBody.FirstName, "last_name": reqBody.LastName, "email": reqBody.Email, "phone_no": reqBody.PhoneNo, "company_name": reqBody.CompanyName, "country": reqBody.Country}}

	//updateFileRes, err := collection.UpdateOne(bson.M{"unique_id": findOneLead.UniqueId}, updateField)
	updateFileRes, err := dataAccess.Collection.UpdateOne(bson.M{"unique_id": key}, updateField)

	if err != nil {
		fmt.Println("error in updateFileRes")
		return nil, errors.New("error in updateFileRes")
	}

	reqBodyv2, _ := json.Marshal(reqBody) // convert it to byte so we can store it in redisdb

	//updating in redisdb
	err7 := cache.Set(key, reqBodyv2, 0).Err()

	if err7 == redis.Nil {
		fmt.Println("not able to set the values in redisdb")
	}
	return updateFileRes, nil
}
