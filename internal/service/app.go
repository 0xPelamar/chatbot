package service

type App struct {
	Accounts *AccountService
}

func NewApp(accounts *AccountService) *App {
	return &App{
		Accounts: accounts,
	}
}
