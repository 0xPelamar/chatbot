package teleprompt

import (
	"fmt"
	"gopkg.in/telebot.v4"
	"sync"
	"time"
)

type TelePrompt struct {
	accountPrompts sync.Map
}

func NewTelePrompt() *TelePrompt {
	return &TelePrompt{}
}

type Prompt struct {
	TeleCtx telebot.Context
}

func (tp *TelePrompt) Register(userId int64) <-chan Prompt {
	ch := make(chan Prompt, 1)
	if preChannel, loaded := tp.accountPrompts.LoadAndDelete(userId); loaded {
		close(preChannel.(chan Prompt))
	}
	tp.accountPrompts.Store(userId, ch)
	return ch
}

func (tp *TelePrompt) Dispatch(userId int64, c telebot.Context) bool {
	ch, loaded := tp.accountPrompts.LoadAndDelete(userId)
	if !loaded {
		return false
	}
	select {
	case ch.(chan Prompt) <- Prompt{TeleCtx: c}:
	default:
		return false
	}
	return true

}

func (tp *TelePrompt) AsMessage(userId int64, timeout time.Duration) (*telebot.Message, bool) {
	ch := tp.Register(userId)
	select {
	case val := <-ch:
		return val.TeleCtx.Message(), false
	case <-time.After(timeout):
		return nil, true
	}
}
