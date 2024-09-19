package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func update(w http.ResponseWriter, r *http.Request) {
	var ctx context.Context
	buff := poolDoc.Get().(*Doc)
	err := json.NewDecoder(r.Body).Decode(buff)
	if err != nil {
		_, err = w.Write([]byte(err.Error()))
		if err != nil {
			logrus.Errorf(err.Error())
		}
	}
	_, err = collection.UpdateOne(ctx, bson.M{"title": buff.Title, "val": buff.Val}, options.InsertOne())
	if err != nil {
		w.Write([]byte(err.Error()))
	}
}
