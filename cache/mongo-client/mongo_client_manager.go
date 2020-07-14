package mongoclient

import (
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
)

type MongoClientManager struct {
	mongoClientMap *sync.Map
}

var mongoClientManager *MongoClientManager = nil

func MongoClientManagerInit(configs []MongoClientConfig) error {
	mongoClientManager = &MongoClientManager{
		mongoClientMap: &sync.Map{},
	}

	for _, config := range configs {
		client, err := NewMongoClient(config)
		if err != nil {
			Log.Error(err)
			return err
		}

		mongoClientManager.mongoClientMap.Store(config.Database, client)
	}
	return nil
}

func GetMongoClientManager() *MongoClientManager {
	return mongoClientManager
}

func (r *MongoClientManager) GetMongoClient(database string) *mongo.Client {
	clientobj, ok := r.mongoClientMap.Load(database)
	if !ok {
		Log.Warnf("can't find client in map, database=%v", database)
		return nil
	}

	client := clientobj.(*MongoClient)

	return client.Client
}

func (r *MongoClientManager) GetMongoDatabaseClient(database string) *mongo.Database {
	clientobj, ok := r.mongoClientMap.Load(database)
	if !ok {
		Log.Warnf("can't find client in map, database=%v", database)
		return nil
	}

	client := clientobj.(*MongoClient)

	return client.DatabaseClient
}
