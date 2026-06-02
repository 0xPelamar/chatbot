package prompt

import (
	"errors"
	"log/slog"
	"sync"
	"time"

	"gopkg.in/telebot.v4"
)

var (
	ErrCanceled = errors.New("prompt canceled: a new prompt was registered for this user")
	ErrTimeout  = errors.New("prompt timed out: user did not respond in time")
)

// reply is an internal envelope passed through a user's channel.
// It either carries a real telegram context or a cancellation signal.
type reply struct {
	tCtx     telebot.Context
	canceled bool
}

// Prompter waits for a user's next Telegram message.
// Register a prompt with WaitForMessage, then route incoming
// messages through Fulfill to unblock it.
type Prompter struct {
	pending sync.Map // map[int64]chan reply
}

func NewPrompter() *Prompter {
	return &Prompter{}
}

// subscribe creates a fresh channel for the user and cancels
// any previously pending prompt for that same user.
func (w *Prompter) subscribe(userID int64) <-chan reply {
	ch := make(chan reply, 1)

	if old, existed := w.pending.LoadAndDelete(userID); existed {
		old.(chan reply) <- reply{canceled: true}
	}

	w.pending.Store(userID, ch)
	return ch
}

// Deliver delivers an incoming message to whoever is waiting for
// this user's reply. Returns false if no one is waiting.
func (w *Prompter) Deliver(userID int64, c telebot.Context) bool {
	ch, waiting := w.pending.LoadAndDelete(userID)
	if !waiting {
		return false
	}

	select {
	case ch.(chan reply) <- reply{tCtx: c}:
		return true
	default:
		return false
	}
}

// WaitForMessage blocks until the user sends a message, the timeout
// elapses, or a newer prompt for the same user cancels this one.
func (w *Prompter) WaitForMessage(userID int64, timeout time.Duration) (*telebot.Message, error) {
	ch := w.subscribe(userID)
	slog.Info("waiting for message", "user", userID)
	select {
	case r := <-ch:
		if r.canceled {
			return nil, ErrCanceled
		}
		return r.tCtx.Message(), nil
	case <-time.After(timeout):
		return nil, ErrTimeout
	}
}
