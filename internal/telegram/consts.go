package telegram

import (
	"github.com/0xpelamar/chatbot/internal/entity"
	"gopkg.in/telebot.v4"
	"time"
)

var (
	txtMainMenu              = "🌹به برنامه چت ناشناس خوش اومدی «%s».\nچه کاری برات انجام بدم؟"
	txtGetDisplayName        = "نام نمایشی رو بفرست."
	getCityMessage           = "🏠 Enter your City "
	txtGetGender             = "🙎🏻لطفا جنسیتت را وارد کن "
	getAgeMessage            = "Enter your age"
	txtOnInvalidGender       = "جنسیت را از کیبورد انتخاب کنید"
	txtProfileUpdatedMessage = "پروفایل شما آپدیت شد"
)
var (
	DefaultInputTimeout     = time.Minute * 5
	DefaultInputTimeoutText = "⏰ We were waiting for you but you did not send anything. Please send message when you come ⏰"
	ConfirmText             = "✅ Confirm"
	DeclineText             = "❌ Decline"
	maleGenderKeyboard      = "🙋‍♂️ مرد"
	femaleGenderKeyboard    = "🙋‍♀️ زن"
	nonBinaryGenderKeyboard = "🏳️‍🌈 نان باینری"
)

var (
	selector           = &telebot.ReplyMarkup{}
	btnEditDisplayName = selector.Data("✏️ Edit Name", "editName")
	btnEditProvince    = selector.Data("✏️ Edit Province", "editProvince")
	btnEditAge         = selector.Data("✏️ Edit Age", "editAge")
	btnEditGender      = selector.Data("✏️ Edit Gender", "editGender")
)

func GetAccount(c telebot.Context) entity.Account {
	return c.Get("account").(entity.Account)
}
