package db

import (
	"cli-tool/internal/pkg/db/collection"
	"context"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

func (c client) CreateAdmin(admin collection.Admin) (err error) {
	log.Debug("create admin")

	_, err = c.admin.InsertOne(context.Background(), admin)
	return
}

func (c client) GetAdmin() (data collection.Admin, err error) {
	log.Debug("get admin")

	err = c.admin.FindOne(context.Background(), bson.M{}).Decode(&data)
	return
}
