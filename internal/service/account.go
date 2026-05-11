package service

import (
	"context"
	"errors"
	"time"

	"github.com/0xpelamar/chatbot/internal/consts"
	"github.com/0xpelamar/chatbot/internal/entity"
	"github.com/0xpelamar/chatbot/internal/repository"
)

type AccountService struct {
	repo repository.AccountRepository
}

func NewAccountService(repo repository.AccountRepository) *AccountService {
	return &AccountService{
		repo: repo,
	}
}

func (a *AccountService) CreateOrUpdate(ctx context.Context, account entity.Account) (entity.Account, bool, error) {
	savedAccount, err := a.repo.Get(ctx, account.EntityID())

	// User exists
	if err == nil {
		if isChanged(savedAccount, account) {
			savedAccount.FirstName = account.FirstName
			savedAccount.LastName = account.LastName
			savedAccount.Username = account.Username
			savedAccount.DisplayName = account.DisplayName
			savedAccount.Age = account.Age
			savedAccount.Province = account.Province
			savedAccount.Gender = account.Gender
			savedAccount.Language = account.Language

			return savedAccount, false, a.repo.Save(ctx, savedAccount)
		}
		return savedAccount, false, nil
	}

	// User does not exist
	if errors.Is(err, repository.ErrNotFound) {
		account.JoinedAt = time.Now()
		account.Language = consts.English
		return account, true, a.repo.Save(ctx, account)

	}
	return entity.Account{}, false, err
}

func isChanged(savedAccount, newAccount entity.Account) bool {
	return savedAccount.FirstName != newAccount.FirstName ||
		savedAccount.LastName != newAccount.LastName ||
		savedAccount.Username != newAccount.Username ||
		savedAccount.DisplayName != newAccount.DisplayName ||
		savedAccount.Age != newAccount.Age ||
		savedAccount.Gender != newAccount.Gender ||
		savedAccount.Province != newAccount.Province ||
		savedAccount.Language != newAccount.Language
}
