package tg_bot

import (
	config "AvitoMessengeBot/config"
	tele "gopkg.in/telebot.v3"
	"log"
	"time"
)

func NewTeleBot() (*tele.Bot, error) {
	pref := tele.Settings{
		Token:  config.Config.Token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return b, err
	}
	b.Handle("/add", ClientAddCommand)
	b.Handle("/start", StartCommand)
	b.Handle("/list", ListCommand)
	return b, err
}
