package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/api/youtube/v3"
)

type APIServer struct {
	MongoCollection *mongo.Collection
}

type VideoInfo struct {
	Title              string                   `bson:"title"`
	Description        string                   `bson:"description"`
	PublishingDatetime string                   `bson:"publishingdatetime"`
	Thumbnail          youtube.ThumbnailDetails `bson:"thumbnail"`
}

func (s *APIServer) GetVideoInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var (
		maxResults int64 = 10
		offset     int64 = 0
	)
	var err error

	maxResultsQueryParam := r.URL.Query().Get("maxResults")
	if maxResultsQueryParam != "" {
		maxResults, err = strconv.ParseInt(maxResultsQueryParam, 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
	offsetQueryParam := r.URL.Query().Get("offset")
	if offsetQueryParam != "" {
		offset, err = strconv.ParseInt(offsetQueryParam, 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	filter := bson.D{}
	opts := options.Find().SetSort(bson.D{{Key: "publishingdatetime", Value: -1}}).SetLimit(maxResults).SetSkip(offset)
	cursor, _ := s.MongoCollection.Find(context.Background(), filter, opts)
	var results []VideoInfo
	if err := cursor.All(context.Background(), &results); err != nil {
		log.Println("error while parsing result: ", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	res, _ := json.Marshal(results)
	w.Write(res)
}

func (s *APIServer) SearchVideo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	titleQueryParam := r.FormValue("title")
	descQueryParam := r.FormValue("description")
	var filter bson.D
	if titleQueryParam == "" {
		filter = bson.D{
			{Key: "$text", Value: bson.D{{Key: "$search", Value: descQueryParam}}},
		}
	} else if descQueryParam == "" {
		filter = bson.D{
			{Key: "title", Value: titleQueryParam},
		}
	} else {
		filter = bson.D{
			{Key: "title", Value: titleQueryParam},
			{Key: "$text", Value: bson.D{{Key: "$search", Value: descQueryParam}}},
		}
	}

	opts := options.FindOne()
	var result VideoInfo
	_ = s.MongoCollection.FindOne(context.Background(), filter, opts).Decode(&result)

	w.WriteHeader(http.StatusOK)
	res, _ := json.Marshal(result)
	w.Write(res)
}
