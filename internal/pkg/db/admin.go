package db

import (
	"context"
	"github.com/Template7/common/structs"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

func (c client) CreateAdmin(admin structs.Admin) (err error) {
	log.Debug("create admin")

	_, err = c.admin.InsertOne(context.Background(), admin)
	return
}

func (c client) GetAdmin() (data structs.Admin, err error) {
	log.Debug("get admin")

	err = c.admin.FindOne(context.Background(), bson.M{}).Decode(&data)
	return
}
