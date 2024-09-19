package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func read(w http.ResponseWriter, r *http.Request) {
	var ctx context.Context
	var result bson.M
	docTitle := chi.URLParam(r, "docTitle")
	collection.FindOne(ctx, bson.D{{"title", docTitle}}, options.FindOne()).Decode(&result)
	jeson, err := json.Marshal(map[string]bson.M{"val": result})
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(jeson)
}
