package service

import (
	"context"
	"errors"
	"github.com/0xpelamar/chatbot/internal/entity"
	"github.com/0xpelamar/chatbot/internal/repository"
	"time"
)

const (
	DefaultState = "home"
)

type AccountService struct {
	accounts repository.AccountRepository
}

func NewAccountService(accounts repository.AccountRepository) *AccountService {
	return &AccountService{
		accounts: accounts,
	}
}

func (a *AccountService) CreateOrUpdate(ctx context.Context, account entity.Account) (entity.Account, bool, error) {
	savedAccount, err := a.accounts.Get(ctx, account.EntityID())
	// user exists
	if err == nil {
		if savedAccount.Username != account.Username ||
			savedAccount.FirstName != account.FirstName ||
			savedAccount.LastName != account.LastName {
			savedAccount.Username = account.Username
			savedAccount.FirstName = account.FirstName
			savedAccount.LastName = account.LastName
			return savedAccount, false, a.accounts.Save(ctx, savedAccount)
		}
		return savedAccount, false, nil
	}

	// user does not exist
	if errors.Is(err, repository.ErrorNotFound) {
		account.JoinedAt = time.Now()
		account.State = DefaultState
		return account, true, a.accounts.Save(ctx, account)
	}

	return entity.Account{}, false, err
}

func (a *AccountService) Update(ctx context.Context, account entity.Account) error {
	return a.accounts.Save(ctx, account)
}
