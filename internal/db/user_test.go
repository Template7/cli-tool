package db

import (
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"testing"
)

func TestClient_ListFakeUsers(t *testing.T) {
	viper.AddConfigPath("../../config")

	ctx := context.WithValue(context.Background(), "traceId", uuid.NewString())
	fu := New().ListFakeUsers(ctx)
	t.Log(fu)
}
