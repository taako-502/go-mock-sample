package sample

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockMongoDB struct {
	mock.Mock
}

func (m *MockMongoDB) FindUserByID(ctx context.Context, id primitive.ObjectID) (*user, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*user), args.Error(1)
}

func TestSample_usingTestifyMock(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	userID, _ := primitive.ObjectIDFromHex("000000000000000000000001")
	expectedUser := &user{ID: userID, Name: "Test User"}

	mockDB := new(MockMongoDB)
	mockDB.On("FindUserByID", mock.Anything, userID).Return(expectedUser, nil)

	user, err := Sample(mockDB, ctx)
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, expectedUser, user)

	mockDB.AssertExpectations(t)
}
