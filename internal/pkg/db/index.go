package db

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

var CollectionIndexes = map[string][]mongo.IndexModel{
	"user":               user,
	"transactionHistory": transactionHistory,
}

var (
	user = []mongo.IndexModel{
		{
			Keys: bsonx.Doc{
				{
					Key:   "mobile",
					Value: bsonx.Int32(1),
				},
			},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys: bsonx.Doc{
				{
					Key:   "email",
					Value: bsonx.Int32(1),
				},
			},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys: bsonx.Doc{
				{
					Key:   "user_id",
					Value: bsonx.Int32(1),
				},
			},
			Options: options.Index().SetUnique(true),
		},
	}
	transactionHistory = []mongo.IndexModel{
		{
			Keys: bsonx.Doc{
				{
					Key:   "request_data.from_wallet_id",
					Value: bsonx.Int32(1),
				},
				{
					Key:   "request_data.to_wallet_id",
					Value: bsonx.Int32(1),
				},
				{
					Key:   "transaction_id",
					Value: bsonx.Int32(1),
				},
			},
		},
	}
)
