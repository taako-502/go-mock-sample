package sample

import (
	"context"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Sample(si SampleInteface, ctx context.Context) (*user, error) {
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
