package service

import (
	"context"

	"github.com/0xpelamar/chatbot/internal/consts"
	"github.com/0xpelamar/chatbot/internal/entity"
	"github.com/0xpelamar/chatbot/internal/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"testing"
)

func TestAccountService_CreateOrUpdateWithUserExists(t *testing.T) {
	accRep := repository.NewMockAccountRepository(t)
	s := NewAccountService(accRep)
	ctx := context.Background()
	accRep.On("Get", mock.Anything, entity.NewID("account", 21)).
		Return(entity.Account{
			ID:        21,
			FirstName: "zxcv",
		}, nil).Once()
	accRep.On("Save", mock.Anything, mock.MatchedBy(func(acc entity.Account) bool {
		return acc.FirstName == "changed"
	})).Return(nil).Once()
	newAcc, created, err := s.CreateOrUpdate(ctx, entity.Account{
		ID:        21,
		FirstName: "changed",
		Language:  consts.Kurdish,
	})
	assert.NoError(t, err)
	assert.Equal(t, int64(21), newAcc.ID)
	assert.Equal(t, "changed", newAcc.FirstName)
	assert.Equal(t, consts.Kurdish, newAcc.Language)
	assert.False(t, created)
	accRep.AssertExpectations(t)
}

func TestAccountService_CreateOrUpdateWithUserNotExists(t *testing.T) {
	accRep := repository.NewMockAccountRepository(t)
	s := NewAccountService(accRep)
	ctx := context.Background()
	accRep.On("Get", mock.Anything, entity.NewID("account", 21)).
		Return(entity.Account{}, repository.ErrNotFound).Once()
	accRep.On("Save", mock.Anything, mock.MatchedBy(func(acc entity.Account) bool {
		return acc.FirstName == "qwer"
	})).Return(nil).Once()
	newAcc, created, err := s.CreateOrUpdate(ctx, entity.Account{
		ID:        21,
		FirstName: "qwer",
	})
	assert.NoError(t, err)
	assert.Equal(t, int64(21), newAcc.ID)
	assert.Equal(t, "qwer", newAcc.FirstName)
	assert.True(t, created)
	accRep.AssertExpectations(t)
}

func TestAccountService_CreateOrUpdateWithUserHasNotChanged(t *testing.T) {
	accRep := repository.NewMockAccountRepository(t)
	s := NewAccountService(accRep)
	ctx := context.Background()
	accRep.On("Get", mock.Anything, entity.NewID("account", 21)).
		Return(entity.Account{ID: 21, FirstName: "qwer"}, nil).Once()

	newAcc, created, err := s.CreateOrUpdate(ctx, entity.Account{
		ID:        21,
		FirstName: "qwer",
	})
	assert.NoError(t, err)
	assert.Equal(t, int64(21), newAcc.ID)
	assert.Equal(t, "qwer", newAcc.FirstName)
	assert.False(t, created)
	accRep.AssertExpectations(t)
}
