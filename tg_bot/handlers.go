package tg_bot

import (
	"AvitoMessengeBot/api"
	"AvitoMessengeBot/config"
	"AvitoMessengeBot/core"
	"AvitoMessengeBot/dao"
	"AvitoMessengeBot/db"
	tele "gopkg.in/telebot.v3"
)

var ClientAddCommand = func(ctx tele.Context) error {
	tags := ctx.Args() // list of arguments splitted by a space
	client := db.Client{}
	for i, tag := range tags {
		switch i {
		case 0:
			client.ClientId = tag
		case 1:
			client.ClientSecret = tag
		}
	}
	newConfiguration := api.NewConfiguration()
	apiClient := api.NewAPIClient(newConfiguration)
	clientService := dao.NewClientService(dao.NewClientGorm(config.Config.DB))
	err := core.StartNewClient(clientService, apiClient, client)
	if err != nil {
		ctx.Send(err.Error())
		return err
	}
	answer := "Успешно добавлен аккаунт\n"
	return ctx.Send(answer)
}
var StartCommand = func(ctx tele.Context) error {
	answer := "Здравствуй" + ctx.Sender().Username + "\n"
	return ctx.Send(answer)
}

var ListCommand = func(ctx tele.Context) error {
	service := dao.NewClientService(dao.NewClientGorm(config.Config.DB))
	clients, err := service.Dao.ReadAll()
	if err != nil {
		return err
	}
	var answer = "Отслеживаемые аккаунты:\n"
	for _, client := range clients {
		answer += client.User.Name + " " + client.User.ProfileUrl + "\n"
	}
	return ctx.Send(answer)
}
