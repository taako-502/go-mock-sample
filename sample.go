package sample

import (
	"context"

	"github.com/pkg/errors"
	"github.com/taako-502/go-testify-mock/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

	func Sample(si database.SampleInteface, ctx context.Context) (*database.User, error) {
	id, err := primitive.ObjectIDFromHex("000000000000000000000001")
	if err != nil {
		return nil, errors.Wrap(err, "primitive.ObjectIDFromHex")
	}

	user, err := si.FindUserByID(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "SampleInteface.FindUserByID")
	}

	return user, nil
}
