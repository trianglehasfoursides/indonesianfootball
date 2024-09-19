package main

import (
	"context"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func readAll(w http.ResponseWriter, r *http.Request) {
	var ctx context.Context
	var all []byte
	result, err := collection.Find(ctx, bson.D{}, options.Find())
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	err = result.All(ctx, *&all)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(all)
}
