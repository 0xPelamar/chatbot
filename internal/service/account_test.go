package service

import (
	"context"
	"github.com/0xpelamar/chatbot/internal/entity"
	"github.com/0xpelamar/chatbot/internal/repository"
	"github.com/0xpelamar/chatbot/internal/repository/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestAccountService_CreateOrUpdateWithUserExists(t *testing.T) {
	accRep := mocks.NewMockAccountRepository(t)
	s := NewAccountService(accRep)

	accRep.On("Get", mock.Anything, entity.NewID("account", 33)).Return(entity.Account{
		ID:        33,
		FirstName: "Pelamar",
	}, nil).Once()

	accRep.On("Save", mock.Anything, mock.MatchedBy(func(acc entity.Account) bool {
		return acc.FirstName == "Ashkan"
	})).Return(nil).Once()

	acc, created, err := s.CreateOrUpdate(context.Background(), entity.Account{
		ID:        33,
		FirstName: "Ashkan",
	})
	assert.NoError(t, err)
	assert.Equal(t, created, false)
	assert.Equal(t, "Ashkan", acc.FirstName)
	assert.Equal(t, int64(33), acc.ID)

	accRep.AssertExpectations(t)
}

func TestAccountService_CreateOrUpdateWithUserNotExists(t *testing.T) {
	accRep := mocks.NewMockAccountRepository(t)
	s := NewAccountService(accRep)

	accRep.On("Get", mock.Anything, entity.NewID("account", 33)).Return(entity.Account{}, repository.ErrorNotFound).Once()

	accRep.On("Save", mock.Anything, mock.MatchedBy(func(acc entity.Account) bool {
		return acc.FirstName == "Ashkan"
	})).Return(nil).Once()

	acc, created, err := s.CreateOrUpdate(context.Background(), entity.Account{
		ID:        33,
		FirstName: "Ashkan",
	})
	assert.NoError(t, err)
	assert.Equal(t, created, true)
	assert.Equal(t, "Ashkan", acc.FirstName)
	assert.Equal(t, int64(33), acc.ID)

	accRep.AssertExpectations(t)
}

func TestAccountService_CreateOrUpdateWithUserHasNotChanged(t *testing.T) {
	accRep := mocks.NewMockAccountRepository(t)
	s := NewAccountService(accRep)

	accRep.On("Get", mock.Anything, entity.NewID("account", 33)).Return(entity.Account{
		ID:        33,
		FirstName: "Pelamar",
	}, nil).Once()

	acc, created, err := s.CreateOrUpdate(context.Background(), entity.Account{
		ID:        33,
		FirstName: "Pelamar",
	})
	assert.NoError(t, err)
	assert.Equal(t, created, false)
	assert.Equal(t, "Pelamar", acc.FirstName)
	assert.Equal(t, int64(33), acc.ID)

	accRep.AssertExpectations(t)
}
