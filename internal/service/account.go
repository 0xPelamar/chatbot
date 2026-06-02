package service

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/0xpelamar/chatbot/internal/consts"
	"github.com/0xpelamar/chatbot/internal/entity"
	"github.com/0xpelamar/chatbot/internal/repository"
	"github.com/google/uuid"
)

type AccountService struct {
	repo repository.AccountRepository
}

func NewAccountService(repo repository.AccountRepository) *AccountService {
	return &AccountService{
		repo: repo,
	}
}

func (a *AccountService) Get(ctx context.Context, ID entity.ID) (entity.Account, error) {
	return a.repo.Get(ctx, ID)
}

func (a *AccountService) CreateOrUpdate(ctx context.Context, account entity.Account) (entity.Account, bool, error) {
	savedAccount, err := a.repo.Get(ctx, account.EntityID())

	// User exists
	if err == nil {
		if isChanged(savedAccount, account) {
			savedAccount.FirstName = account.FirstName
			savedAccount.LastName = account.LastName
			savedAccount.Username = account.Username
			slog.Info("User existed and updated")
			return savedAccount, false, a.repo.Save(ctx, savedAccount)
		}
		slog.Info("User existed and not updated")
		return savedAccount, false, nil
	}

	// User does not exist
	if errors.Is(err, repository.ErrNotFound) {
		account.JoinedAt = time.Now()
		account.Language = consts.English
		account.UUID = uuid.New().String()
		slog.Info("New account created")
		return account, true, a.repo.Save(ctx, account)

	}
	return entity.Account{}, false, err
}

func (a *AccountService) Update(ctx context.Context, account entity.Account) error {
	return a.repo.Save(ctx, account)
}

func isChanged(savedAccount, newAccount entity.Account) bool {
	return savedAccount.FirstName != newAccount.FirstName ||
		savedAccount.LastName != newAccount.LastName ||
		savedAccount.Username != newAccount.Username
}
