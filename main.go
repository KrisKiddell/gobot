package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/kriskiddell/gobot/internal/bot"
	"github.com/kriskiddell/plog"
)

func main() {
	godotenv.Load()

	bot, err := bot.NewBot(os.Getenv("TOKEN"), os.Getenv("APP_ID"))

	if err != nil {
		plog.Error.Fatal(err)
	}

	bot.RegisterSlashCommands()

	err = bot.Session.Open()

	if err != nil {
		fmt.Println("error opening connection,", err)
		bot.Session.Close()
		return
	}

	plog.Success.Println("Bot is online")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	bot.RemoveSlashCommands()

	bot.Session.Close()

}
