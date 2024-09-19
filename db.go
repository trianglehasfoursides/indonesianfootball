package main

import (
	"context"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func init() {
	var ctx context.Context
	con, err := mongo.Connect(ctx, options.Client())
	if err != nil {
		logrus.Fatalf(err.Error())
	}
	collection = con.Database("indonesianfootball", options.Database()).Collection("doc", options.Collection())
}
