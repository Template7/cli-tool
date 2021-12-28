package db

import (
	"cli-tool/internal/pkg/config"
	"context"
	"fmt"
	"github.com/Template7/common/structs"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

type client struct {
	mongo struct {
		client *mongo.Client
		admin  *mongo.Collection
		user   *mongo.Collection
		token  *mongo.Collection
	}
	mysql struct {
		db      *gorm.DB
		wallet  *gorm.DB
		balance *gorm.DB
	}
}

var (
	once     sync.Once
	instance *client
)

func New() ClientInterface {
	once.Do(func() {
		c, err := mongo.Connect(nil, options.Client().ApplyURI(config.New().Mongo.ConnectionString))
		if err != nil {
			log.Fatal(err)
		}
		db := c.Database(config.New().Mongo.Db)
		instance = &client{}
		instance.mongo.client = c
		instance.mongo.admin = db.Collection("admin")
		instance.mongo.user = db.Collection("user")
		instance.mongo.token = db.Collection("token")

		// mysql
		sqlDb, err := gorm.Open(mysql.Open(config.New().MySql.ConnectionString), &gorm.Config{})
		if err != nil {
			log.Fatal(err)
		}
		instance.mysql.db = sqlDb
		instance.mysql.wallet = instance.mysql.db.Model(&structs.Wallet{})
		instance.mysql.balance = instance.mysql.db.Model(&structs.Balance{})

		if err := c.Ping(nil, nil); err != nil {
			log.Fatal(err)
		}
		//instance.initIndex(db)
		log.Debug("mongo client initialized")
	})
	return instance
}

func (c *client) InitDb(force bool) {
	if force {
		log.Debug("delete db")
		if err := c.mongo.client.Database(config.New().Mongo.Db).Drop(context.Background()); err != nil {
			log.Fatal("fail to drop databases: ", config.New().Mongo.Db, ". ", err.Error())
		}
		if err := c.mysql.db.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS `%s`", config.New().MySql.Db)).Error; err != nil {
			log.Fatal("fail to drop mysql db: ", config.New().MySql.Db, ". ", err.Error())
		}
	}

	// init mongo index
	ctx := context.Background()
	for col, idx := range CollectionIndexes {
		log.Debug("create index for collection: ", col)
		_, err := c.mongo.client.Database(config.New().Mongo.Db).Collection(col).Indexes().CreateMany(ctx, idx)
		if err != nil {
			log.Error("fail to create index: ", err.Error())
			panic(err)
		}
	}

	// init sql
	//instance.mysql.db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s;", sqlTempDb))
	instance.mysql.db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s;", config.New().MySql.Db))
	instance.mysql.db = instance.mysql.db.Exec(fmt.Sprintf("USE %s;", config.New().MySql.Db))
	if err := instance.mysql.db.AutoMigrate(&structs.Wallet{}, &structs.Balance{}); err != nil {
		log.Error("fail to create table: ", err.Error())
		panic(err)
	}
	instance.mysql.db.Exec(fmt.Sprintf("CREATE USER IF NOT EXISTS '%s'@'localhost' IDENTIFIED BY '%s';", config.New().MySql.Username, config.New().MySql.Password))
	instance.mysql.db.Exec(fmt.Sprintf("GRANT ALL ON %s.* TO '%s'@'localhost';", config.New().MySql.Db, config.New().MySql.Username))
	return
}
