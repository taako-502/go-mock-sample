package sample

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/taako-502/go-testify-mock/database"
	"github.com/taako-502/go-testify-mock/mockgen"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
)

type MockMongoDB struct {
	mock.Mock
}

func (m *MockMongoDB) FindUserByID(ctx context.Context, id primitive.ObjectID) (*database.User, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*database.User), args.Error(1)
}

// stretchr/testify(https://github.com/stretchr/testify)を使ったモックのサンプル
func TestSample_usingTestifyMock(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	userID, _ := primitive.ObjectIDFromHex("000000000000000000000001")
	expectedUser := &database.User{ID: userID, Name: "Test User"}

	mockDB := new(MockMongoDB)
	mockDB.On("FindUserByID", mock.Anything, userID).Return(expectedUser, nil)

	user, err := Sample(mockDB, ctx)
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, expectedUser, user)

	mockDB.AssertExpectations(t)
}

// uber-go/mock(https://github.com/uber-go/mock)を使ったモックのサンプル
func TestSample_usingUberGoMock(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mockgen.NewMockSampleInteface(ctrl)

	// モックの振る舞いを設定
	userID, _ := primitive.ObjectIDFromHex("000000000000000000000001")
	m.EXPECT().
		FindUserByID(gomock.Any(), userID).
		Return(&database.User{}, nil)

	// テスト対象の関数を実行
	Sample(m, context.Background())
}