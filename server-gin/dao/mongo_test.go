package dao

import (
	"context"
	"encoding/json"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestMongoInit(t *testing.T) {
	result := bson.M{}
	uid, err := primitive.ObjectIDFromHex("61ea2ab8eccea804bbcbce49")
	if err != nil {
		t.Error(err)
	}
	err = Mongo.Collection("users").FindOne(context.TODO(), bson.D{{Key: "_id", Value: uid}}).Decode(&result)
	if err != nil && err != mongo.ErrNoDocuments {
		t.Error(err)
		return
	}
	jsonData, err := json.MarshalIndent(result, "", "")
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%s\n", jsonData)
}
