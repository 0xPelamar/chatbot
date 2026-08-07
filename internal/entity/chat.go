package entity

type Chat struct {
	ID            string  `json:"id"`
	Users         []int64 `json:"users"`
	CreatedAtUnix int64   `json:"created_at_unix"`
	State         string  `json:"state"`
}

func (ch Chat) EntityID() ID {
	return NewID("chat", ch.ID)

}
