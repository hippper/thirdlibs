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

		mongoClientManager.mongoClientMap.Store(config.Name, client)
	}
	return nil
}

func GetMongoClientManager() *MongoClientManager {
	return mongoClientManager
}

func (r *MongoClientManager) GetMongoClient(name string) *mongo.Client {
	clientobj, ok := r.mongoClientMap.Load(name)
	if !ok {
		Log.Warnf("can't find client in map, name=%v", name)
		return nil
	}

	client := clientobj.(*MongoClient)

	return client.Client
}

func (r *MongoClientManager) GetMongoDatabaseClient(name string) *mongo.Database {
	clientobj, ok := r.mongoClientMap.Load(name)
	if !ok {
		Log.Warnf("can't find client in map, name=%v", name)
		return nil
	}

	client := clientobj.(*MongoClient)

	return client.DatabaseClient
}
