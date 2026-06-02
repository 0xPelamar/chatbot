package telegram

func mainMenuKeyboard(language byte) [][]string {
	// TODO: implement other languages too
	_ = language

	return [][]string{
		{"Connect me to my special contact 💌"},
		{"List of my contacts 📓"},
		{"Edit profile 🧑", "My anonymous link 🔗"},
		{"🗣 Language", "Manual"},
	}
}

func agesKeyboard() [][]string {
	return [][]string{
		{"12", "13", "14", "15", "16", "17"},
		{"18", "19", "20", "21", "22", "23"},
		{"24", "25", "26", "27", "28", "29"},
		{"30", "31", "32", "33", "34", "35"},
		{"36", "37", "38", "39", "40", "41"},
		{"42", "43", "44", "45", "46", "47"},
		{"48", "49", "50", "51", "52", "53"},
		{"54", "55", "56", "57", "58", "59"},
		{"60", "61", "62", "63", "64", "65"},
		{"66", "67", "68", "69", "70", "71"},
		{"72", "73", "74", "75", "76", "77"},
	}
}
func provincesKeyboard() [][]string {
	return [][]string{
		{"آذربایجان شرقی", "آذربایجان غربی", "اردبیل"},
		{"اصفهان", "البرز", "ایلام"},
		{"بوشهر", "تهران", "چهارمحال و بختیاری"},
		{"خراسان جنوبی", "خراسان رضوی", "خراسان شمالی"},
		{"خوزستان", "زنجان", "سمنان"},
		{"سیستان و بلوچستان", "فارس", "قزوین"},
		{"قم", "کردستان", "کرمان"},
		{"کرمانشاه", "کهگیلویه و بویراحمد", "گلستان"},
		{"گیلان", "لرستان", "مازندران"},
		{"مرکزی", "هرمزگان", "همدان"},
		{"یزد"},
	}
}

func genderKeyboard() [][]string {
	return [][]string{
		{txtFemale, txtMale},
		{txtNonBinary},
	}
}

func languageKeyboard() [][]string {
	return [][]string{
		{txtEnglish},
		{txtKurdish},
		{txtPersian},
		{txtArabic},
		{txtSpanish},
	}
}
