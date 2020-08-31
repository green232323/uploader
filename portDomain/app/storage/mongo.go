package storage

import (
	"context"
	"github.com/dnahurnyi/uploader/portDomain/app/contracts"
	pb "github.com/dnahurnyi/uploader/portDomain/proto/github.com/dnahurnyi/uploader/portDomain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoDB struct {
	collection *mongo.Collection
	url        string
}

func NewMongoRepository(ctx context.Context, url string) (contracts.PortRepository, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		return nil, err
	}
	collection := client.Database("uploader").Collection("ports")
	return &mongoDB{
		url:        url,
		collection: collection,
	}, nil
}

func (db *mongoDB) SaveOrUpdate(key string, value *pb.Port) error {
	opts := options.Update().SetUpsert(true)
	_, err := db.collection.UpdateOne(context.TODO(), bson.D{{"_id", key}}, bson.D{{"$set", value}}, opts)
	return err
}

func (db *mongoDB) Get(ctx context.Context, key string) (*pb.Port, error) {
	bPort := &pb.Port{}
	err := db.collection.FindOne(ctx, bson.D{{"_id", key}}).Decode(bPort)
	return bPort, err
}
