package sample

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Sample() (*user, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	DB, err := newMongoDB("mongodb://localhost:27017", "testdb")
	if err != nil {
		return nil, errors.Wrap(err, "newMongoDB")
	}

	id , err:=  primitive.ObjectIDFromHex("000000000000000000000001")
	if err != nil {
		return nil, errors.Wrap(err, "primitive.ObjectIDFromHex")
	}

	user, err := DB.FindUserByID(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "DB.FindUserByID")
	}

	return user, nil
}