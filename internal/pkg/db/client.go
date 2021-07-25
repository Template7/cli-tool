package db

import (
	"cli-tool/internal/pkg/config"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
)

type client struct {
	User          *mongo.Collection
	Decision      *mongo.Collection
	Game          *mongo.Collection
	Gambling      *mongo.Collection
	Betting       *mongo.Collection
	Gambler       *mongo.Collection
	GambleHistory *mongo.Collection
	Strategy      *mongo.Collection
	StrategyMeta  *mongo.Collection
	Simulation    *mongo.Collection
}

var (
	once     sync.Once
	instance *client
)

func New() *client {
	once.Do(func() {
		c, err := mongo.Connect(nil, options.Client().ApplyURI(config.New().Mongo.ConnectionString))
		if err != nil {
			log.Fatal(err.Error())
			return
		}
		if err := c.Ping(nil, nil); err != nil {
			log.Fatal(err.Error())
		}
		db := c.Database(config.New().Mongo.Db)
		instance = &client{
			//Client:        c,
			User:          db.Collection("user"),
			Game:          db.Collection("game"),
			Decision:      db.Collection("decision"),
			Gambling:      db.Collection("gambling"),
			Betting:       db.Collection("betting"),
			Gambler:       db.Collection("gambler"),
			GambleHistory: db.Collection("gamble_history"),
			Strategy:      db.Collection("strategy"),
			StrategyMeta:  db.Collection("strategy_meta"),
			Simulation:    db.Collection("simulation"),
		}

		log.Debug("mongo client initialized")
	})
	return instance
}
