package cli

import (
	"cli-tool/internal/pkg/config"
	"cli-tool/internal/pkg/db"
	"cli-tool/internal/pkg/util"
	"context"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
)

func Test_createAdmin(t *testing.T) {
	viper.AddConfigPath("../../../configs")
	viper.Set("Mongo.Db", "testDb")

	testUsername := "admin"
	testPassword := "password"
	createAdmin(testUsername, testPassword)
	admin, err := db.New().GetAdmin()
	if err != nil {
		t.Error(err)
	}
	if !util.CheckPasswordHash(testPassword, admin.Password) {
		t.Error("hashed password did not passed")
	}

	c, _ := mongo.Connect(nil, options.Client().ApplyURI(config.New().Mongo.ConnectionString))
	if err != nil {
		t.Error(err)
	}
	db := c.Database(config.New().Mongo.Db)
	db.Drop(context.Background())
}
