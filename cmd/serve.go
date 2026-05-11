package cmd

import (
	"log/slog"
	"os"

	"github.com/0xpelamar/chatbot/internal/repository"
	"github.com/0xpelamar/chatbot/internal/repository/redis"
	"github.com/0xpelamar/chatbot/internal/service"
	"github.com/0xpelamar/chatbot/internal/telegram"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:           "serve",
	Short:         "Start serving",
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE:          serve,
}

func serve(cmd *cobra.Command, args []string) error {
	// setup repositories
	redisClient, err := redis.NewRedisClient(os.Getenv("REDIS_URL"))
	if err != nil {
		slog.Error("could not connect to the redis", "error", err)
		return err
	}
	slog.Info("Connected to the redis successfully.")
	accRepo := repository.NewAccountRedis(redisClient)

	accService := service.NewAccountService(accRepo)

	app := service.NewApp(accService)

	tel, err := telegram.NewTelegram(app)
	if err != nil {
		slog.Error("could not create telegram", "error", err)
		return err
	}
	slog.Info("Connected to the telegram successfully.")
	tel.Start()
	return nil

}

func init() {
	rootCmd.AddCommand(serveCmd)

}
