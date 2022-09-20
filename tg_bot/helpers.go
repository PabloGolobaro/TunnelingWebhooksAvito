package tg_bot

import (
	"fmt"
	tele "gopkg.in/telebot.v3"
	"log"
	"strconv"
)

func SendToAdmins(bot *tele.Bot, admins []string) error {
	for _, admin := range admins {
		id, err := strconv.Atoi(admin)
		if err != nil {
			log.Println(fmt.Errorf("Ошибка отправки оповещения админам: %v", err))
			continue
		}
		_, err = bot.Send(tele.ChatID(id), "Бот запущен")
		if err != nil {
			return fmt.Errorf("Ошибка отправки оповещения админам: %v", err)
		}
	}
	return nil
}
