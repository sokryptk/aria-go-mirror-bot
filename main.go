package main

import (
	"aria-go-mirror-bot/commandHandlers"
	"aria-go-mirror-bot/env"
	"fmt"
	"github.com/PaulSonOfLars/gotgbot"
	"github.com/PaulSonOfLars/gotgbot/ext"
	"github.com/PaulSonOfLars/gotgbot/handlers"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"net/http"
	"os"
	"time"
)

func init() {
	fmt.Println("aria-go-mirror-bot started!")
	fmt.Println("VERSION : v1.0.0")
}
func main() {
	cfg := zap.NewProductionEncoderConfig()
	cfg.EncodeLevel = zapcore.CapitalLevelEncoder
	cfg.EncodeTime = zapcore.RFC3339TimeEncoder

	logger := zap.New(zapcore.NewCore(zapcore.NewConsoleEncoder(cfg), os.Stdout, zap.InfoLevel))
	defer logger.Sync() // flushes buffer, if any

	updater, err := gotgbot.NewUpdater(env.Config.BotToken, logger)
	if err != nil {
		log.Fatal(err)
	}

	ext.DefaultTgBotGetter.Client = &http.Client{
		Timeout: 50*time.Second,
	}

	updater.Dispatcher.AddHandler(handlers.NewArgsCommand(env.Config.Commands.MirrorCommands, commandHandlers.MirrorHandler))
	updater.Dispatcher.AddHandler(handlers.NewArgsCommand(env.Config.Commands.MirrorTarCommands, commandHandlers.MirrorHandler))
	updater.Dispatcher.AddHandler(handlers.NewArgsCommand(env.Config.Commands.ListCommands, commandHandlers.MirrorHandler))
	//updater.Dispatcher.AddHandler(handlers.NewCommand(env.Config.Commands.StatusCommands, commandHandlers.MirrorHandler))


	if err := updater.StartPolling(); err != nil {
		panic(err)
	}


	updater.Idle()
}
