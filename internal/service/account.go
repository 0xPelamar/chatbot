package service

import (
	"context"
	"errors"
	"github.com/0xpelamar/chatbot/internal/entity"
	"github.com/0xpelamar/chatbot/internal/repository"
	"time"
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
		//var isChanged = savedAccount.Username != account.Username ||
		//	savedAccount.FirstName != account.FirstName ||
		//	savedAccount.LastName != account.LastName ||
		//	savedAccount.DisplayName != account.DisplayName ||
		//	savedAccount.Age != account.Age ||
		//	savedAccount.Province != account.Province ||
		//	savedAccount.City != account.City ||
		//	savedAccount.Gender != account.Gender
		//if isChanged {
		//	savedAccount.Username = account.Username
		//	savedAccount.FirstName = account.FirstName
		//	savedAccount.LastName = account.LastName
		//	savedAccount.DisplayName = account.DisplayName
		//	savedAccount.Age = account.Age
		//	savedAccount.Province = account.Province
		//	savedAccount.City = account.City
		//	savedAccount.Gender = account.Gender
		//	return savedAccount, false, a.accounts.Save(ctx, savedAccount)
		//}
		return savedAccount, false, nil
	}

	// user does not exist
	if errors.Is(err, repository.ErrorNotFound) {
		account.JoinedAt = time.Now()
		account.DisplayName = account.FirstName
		return account, true, a.accounts.Save(ctx, account)
	}

	return entity.Account{}, false, err
}

func (a *AccountService) Update(ctx context.Context, account entity.Account) error {
	return a.accounts.Save(ctx, account)
}
