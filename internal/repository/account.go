package repository

import "github.com/0xpelamar/chatbot/internal/entity"

var _ AccountRepository = &Account{}

type Account struct {
	CommonBehaviour[entity.Account]
}
