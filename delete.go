package main

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func delete(w http.ResponseWriter, r *http.Request) {
	var ctx context.Context
	docTitle := chi.URLParam(r, "docTitle")
	_, err := collection.DeleteOne(ctx, bson.D{{"title", docTitle}}, options.Delete())
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("ok"))
}
