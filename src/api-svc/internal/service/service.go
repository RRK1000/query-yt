package service

import (
	"context"

	"github.com/gorilla/mux"
	"github.com/query-yt/src/api-svc/internal/handler"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	// Mongo
	db         = "yt"
	collection = "videoinfo"
	username   = "root"
	password   = "root123" // for demo
)

func NewHTTPServer() *handler.APIServer {
	//initialize mongo client
	mongoCli, err := mongo.Connect(
		context.Background(),
		options.Client().
			ApplyURI(getMongoConnStr()).
			SetAuth(options.Credential{Username: username, Password: password}),
	)
	if err != nil {
		panic(err)
	}
	mongoCollection := mongoCli.Database(db).Collection(collection)
	return &handler.APIServer{
		MongoCollection: mongoCollection,
	}
}

func SetupRoutes(r *mux.Router, svc *handler.APIServer) {
	r.HandleFunc("/api/v1/videoinfo", svc.GetVideoInfo).Methods("GET")
	r.HandleFunc("/api/v1/video", svc.SearchVideo).Methods("GET")
}
