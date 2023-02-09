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
	Title              string
	Description        string
	PublishingDatetime string
	Thumbnail          youtube.ThumbnailDetails
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
		}
	}
	offsetQueryParam := r.URL.Query().Get("offset")
	if offsetQueryParam != "" {
		offset, err = strconv.ParseInt(offsetQueryParam, 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
	}

	filter := bson.D{}
	opts := options.Find().SetSort(bson.D{{Key: "publishingdatetime", Value: -1}}).SetLimit(maxResults).SetSkip(offset)
	cursor, _ := s.MongoCollection.Find(context.Background(), filter, opts)
	var results []VideoInfo
	if err := cursor.All(context.Background(), &results); err != nil {
		log.Println(err)
		return
	}
	w.WriteHeader(http.StatusBadRequest)
	res, _ := json.Marshal(results)
	w.Write(res)

	// for _, result := range results {
	// 	res, _ := json.Marshal(result)
	// 	fmt.Println(string(res))
	// }
}
