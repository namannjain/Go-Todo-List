package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type manager struct {
	connection *mongo.Client
	ctx        context.Context
	cancel     context.CancelFunc
}

var Mgr Manager

type Manager interface {
	Insert(interface{}, string) (interface{}, error)
	Delete(primitive.ObjectID, string) error
	DeleteAll(string) error
	Fetch(primitive.ObjectID, string) (interface{}, error)
	FetchAll(string) (interface{}, error)
	Update(primitive.ObjectID, map[string]interface{}, string) (interface{}, error)
}

func ConnectDb() {
	dbUri := os.Getenv("DB_HOST")
	client, err := mongo.NewClient(options.Client().ApplyURI(fmt.Sprintf("%s%s", "mongodb://", dbUri)))
	panicErr(err)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	panicErr(err)

	err = client.Ping(ctx, readpref.Primary())
	panicErr(err)

	Mgr = &manager{connection: client, ctx: ctx, cancel: cancel}
}

func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}
