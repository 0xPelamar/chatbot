package entity

import "time"

type Account struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`

	JoinedAt    time.Time `json:"joined_at"`
	DisplayName string    `json:"display_name"`
	Age         int64     `json:"age"`
	Province    string    `json:"province"`
	City        string    `json:"city"`
	Gender      byte      `json:"gender"`
	Language    int64     `json:"language"`
}

func (a Account) EntityID() ID {
	return NewID("account", a.ID)
}
