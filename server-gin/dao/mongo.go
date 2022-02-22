package dao

import (
	"context"
	"log"
	"server-gin/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MongoCli *mongo.Client
	Mongo    *mongo.Database
)

func init() {
	conf := config.GetConfig().Mongo
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(conf.Uri))
	if err != nil {
		log.Fatalf("mongo.Connect error:%v", err)
	}
	MongoCli, Mongo = client, client.Database(conf.DbName)
}
