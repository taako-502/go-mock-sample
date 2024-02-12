package sample

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongodb struct {
	db *mongo.Database
}

type user struct {
	ID    int    `bson:"_id"`
	Name  string `bson:"name"`
	Email string `bson:"email"`
}

func (m *mongodb) FindUserByID(ctx context.Context, id primitive.ObjectID) (*user, error) {
	var result user
	filter := bson.M{"_id": id} 
	if err := m.db.Collection("users").FindOne(ctx, filter).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func newMongoDB(connectString string, DBName string) (*mongodb, error) {
	clientOptions := options.Client().ApplyURI(connectString)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	if err = client.Ping(context.TODO(), nil);err != nil {
		return nil, err
	}

	return &mongodb{
		db: client.Database(DBName),
	}, nil
}
