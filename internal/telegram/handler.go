package telegram

import (
	"context"
	"fmt"
	"log/slog"

	"gopkg.in/telebot.v4"
)

func (t *Telegram) setupHandlers() {
	t.bot.Use(t.registerMiddleware)
	t.bot.Handle("/start", t.start)
	t.bot.Handle(telebot.OnText, t.textHandler)
	t.bot.Handle(&btnEditDisplayName, t.editDisplayName)
	t.bot.Handle(&btnEditAge, t.editAge)
	t.bot.Handle(&btnEditProvince, t.editProvince)
	t.bot.Handle(&btnEditGender, t.editGender)
}

func (t *Telegram) textHandler(c telebot.Context) error {
	if t.prompter.Deliver(c.Sender().ID, c) {
		return nil
	}

	switch c.Message().Text {
	case "Edit profile 🧑":
		return t.editProfile(c)
	case "🗣 Language":
		return t.editLanguage(c)
	case "My anonymous link 🔗":
		return t.getAnonymousLink(c)
	default:
		return t.home(c)
	}
}

func (t *Telegram) editLanguage(c telebot.Context) error {
	msg, err := t.Input(c, InputConfig{
		Prompt:         "Edit Language",
		PromptKeyboard: languageKeyboard(),
		Validator:      choiceValidator(make1DArray(languageKeyboard())...),
		confirm: confirmStep{
			BuildPrompt: func(msg *telebot.Message) string {
				return fmt.Sprintf("Are you sure to change the bot language to %s", msg.Text)
			},
		},
	})
	if err != nil {
		slog.Error("error while editing language", "err", err)
		return err
	}
	account := getAccount(c)
	account.Language = getLanguageCode(msg.Text)
	slog.Info(fmt.Sprintf("The user entered language: %d", account.Language))
	if err := t.App.Account.Update(context.Background(), account); err != nil {
		_ = c.Send(fmt.Sprintln("Couldn't update the langauge"))
		return err
	}
	c.Set("account", account)
	_ = c.Send(fmt.Sprintf("the language updated successfully to %d", account.Language))
	return nil
}

func (t *Telegram) getAnonymousLink(c telebot.Context) error {
	account := getAccount(c)
	link := fmt.Sprintf("https://t.me/chatrobotchatbot?start=%s", account.UUID)
	return c.Reply(fmt.Sprintf("This is your anonymous link\n%s", link))
}
