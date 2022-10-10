package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"hexagonal_boilerplate/shared/config"
)

func Connect(cfg *config.Config) *mongo.Client {
	uri := fmt.Sprintf("mongodb://%v:%v@%v:%v/?replicaSet=rs&readPreference=primary&ssl=false", cfg.Database.Mongodb.Username, cfg.Database.Mongodb.Password, cfg.Database.Mongodb.Host, cfg.Database.Mongodb.Port)
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))

	err = client.Connect(context.Background())
	if err != nil {
		panic(err)
	}

	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		panic(err)
	}
	return client
}
