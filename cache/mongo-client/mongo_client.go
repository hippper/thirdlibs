package mongoclient

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Mongo client config
type MongoClientConfig struct {
	URL      string
	Database string
}

type MongoClient struct {
	Client         *mongo.Client
	DatabaseClient *mongo.Database
}

func NewMongoClient(c MongoClientConfig) (*MongoClient, error) {
	// 设置客户端参数
	clientOptions := options.Client().ApplyURI(c.URL)

	// 连接到MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		Log.Error(err)
		return nil, err
	}

	// 检查链接
	err = client.Ping(context.Background(), nil)
	if err != nil {
		client.Disconnect(context.Background())
		Log.Error(err)
		return nil, err
	}

	databaseClient := client.Database(c.Database)

	mongoClient := &MongoClient{
		Client:         client,
		DatabaseClient: databaseClient,
	}

	return mongoClient, nil
}
