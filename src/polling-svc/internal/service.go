package internal

import (
	"context"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
)

type Service interface {
	PollYTApi(ctx context.Context)
	postVideoInfo(ctx context.Context, vList []interface{}) error
}

const (
	// Google APIs
	developerKey = "AIzaSyDJ-7c8MCxZ43wRuIPSHONQyE-eChGkBP8" // for demo

	// Mongo
	db         = "yt"
	collection = "videoinfo"
	username   = "root" 
	password   = "root123" // for demo
)

type VideoInfo struct {
	VideoId            string `bson:"_id"`
	Title              string
	Description        string
	PublishingDatetime string
	Thumbnail          youtube.ThumbnailDetails
}

type PollingServer struct {
	MongoCollection *mongo.Collection
	YtSvc           *youtube.Service
}

func InitPollingSvc() Service {
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

	//initalize google api connection
	httpCli := &http.Client{
		Transport: &transport.APIKey{Key: developerKey},
	}
	service, err := youtube.New(httpCli)
	if err != nil {
		log.Fatalf("Error creating new YouTube client: %v", err)
	}

	return &PollingServer{
		MongoCollection: mongoCollection,
		YtSvc:           service,
	}
}

func (p *PollingServer) PollYTApi(ctx context.Context) {
	call := p.YtSvc.Search.List([]string{"id", "snippet"}).
		MaxResults(10)
	response, err := call.Do()
	if err != nil {
		return
	}

	vList := []interface{}{}
	for _, item := range response.Items {
		vInfo := VideoInfo{
			VideoId:            item.Id.VideoId,
			Title:              item.Snippet.Title,
			Description:        item.Snippet.Description,
			PublishingDatetime: item.Snippet.PublishedAt,
			Thumbnail:          *item.Snippet.Thumbnails,
		}
		vList = append(vList, vInfo)
	}
	err = p.postVideoInfo(ctx, vList)
	if err != nil {
		log.Println(err)
	}
}

func (p *PollingServer) postVideoInfo(ctx context.Context, vList []interface{}) error {
	_, err := p.MongoCollection.InsertMany(ctx, vList)
	if err != nil {
		return err
	}
	return nil
}
