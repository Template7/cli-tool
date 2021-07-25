package db

import (
	"cli-tool/internal/pkg/config"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
)

type client struct {
	admin *mongo.Collection
	user  *mongo.Collection
	token *mongo.Collection
}

var (
	once     sync.Once
	instance *client
)

func New() *client {
	once.Do(func() {
		c, err := mongo.Connect(nil, options.Client().ApplyURI(config.New().Mongo.ConnectionString))
		if err != nil {
			log.Fatal(err)
		}
		db := c.Database(config.New().Mongo.Db)
		instance = &client{
			admin: db.Collection("admin"),
			user:  db.Collection("user"),
			token: db.Collection("token"),
		}
		if err := c.Ping(nil, nil); err != nil {
			log.Fatal(err)
		}
		//instance.initIndex(db)
		log.Debug("mongo client initialized")
	})
	return instance
}

//func (c client) initIndex(db *mongo.Database) {
//	ctx := context.Background()
//	for col, idx := range CollectionIndexes {
//		log.Debug("create index for collection: ", col)
//
//		//db := c.Database(config.New().Mongo.Db)
//		_, err := db.Collection(col).Indexes().CreateMany(ctx, idx)
//		if err != nil {
//			log.Error("unable to create index for collection: ", col, ". ", err.Error())
//			panic(err)
//		}
//	}
//}