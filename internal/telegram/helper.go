package telegram

import (
	"github.com/0xpelamar/chatbot/internal/consts"
	"github.com/samber/lo"
	"gopkg.in/telebot.v4"
)

func generateKeyboard(rows [][]string) *telebot.ReplyMarkup {
	mu := &telebot.ReplyMarkup{
		ResizeKeyboard:  true,
		OneTimeKeyboard: true,
		RemoveKeyboard:  true,
		ForceReply:      true,
	}

	mu.Reply(lo.Map(rows, func(row []string, _ int) telebot.Row {
		return mu.Row(lo.Map(row, func(btn string, _ int) telebot.Btn {
			return telebot.Btn{Text: btn}
		})...)
	})...)
	return mu
}

func make1DArray(arr [][]string) []string {
	if len(arr) == 0 {
		return nil
	}
	ln := len(arr)
	res := make([]string, 0, ln*len(arr[0]))

	for _, row := range arr {
		for _, cell := range row {
			res = append(res, cell)
		}
	}
	return res
}

func getLanguageCode(inp string) byte {
	if inp == txtKurdish {
		return consts.Kurdish
	} else if inp == txtPersian {
		return consts.Persian
	} else if inp == txtSpanish {
		return consts.Spanish
	} else if inp == txtArabic {
		return consts.Arabic
	} else {
		return consts.English
	}
}

func getGenderCode(inp string) byte {
	if inp == txtFemale {
		return consts.FemaleGender
	} else if inp == txtNonBinary {
		return consts.NonBinaryGender
	} else {
		return consts.MaleGender
	}
}
