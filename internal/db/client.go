package db

import (
	"github.com/Template7/common/db"
	"github.com/Template7/common/logger"
	"gorm.io/gorm"
	"sync"
)

type Client struct {
	//mongo struct {
	//	user               *mongo.Collection
	//	transactionHistory *mongo.Collection
	//	depositHistory     *mongo.Collection
	//	withdrawHistory    *mongo.Collection
	//}
	sql struct {
		core *gorm.DB
	}
	log *logger.Logger
}

var (
	once     sync.Once
	instance *Client
)

func New() *Client {
	once.Do(func() {
		// nosql
		//mDb := db.NewNoSql().Database(config.New().Db.NoSql.Db)
		instance = &Client{}
		//instance.mongo.user = mDb.Collection("user")
		//instance.mongo.transactionHistory = mDb.Collection("transactionHistory")
		//instance.mongo.depositHistory = mDb.Collection("depositHistory")
		//instance.mongo.withdrawHistory = mDb.Collection("withdrawHistory")

		// sql
		instance.sql.core = db.NewSql()
		instance.log = logger.New().WithService("db")

		instance.log.Info("db client initialized")
	})
	return instance
}
