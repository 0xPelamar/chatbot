package main

import (
	"github.com/0xpelamar/chatbot/cmd"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load(".env")
	cmd.Execute()
}
