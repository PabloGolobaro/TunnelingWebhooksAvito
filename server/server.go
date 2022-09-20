package server

import (
	"AvitoMessengeBot/config"
	"AvitoMessengeBot/dao"
	"fmt"
	"github.com/gin-gonic/gin"
	tele "gopkg.in/telebot.v3"
	"log"
	"net/http"
	"strconv"
)

func NewGinRouter() *gin.Engine {
	r := gin.Default()
	log.Println("Creating routes...")
	r.POST("/", PostWebhook)
	r.GET("/", GetServer)
	return r
}
func PostWebhook(c *gin.Context) {
	var gotWebhook webhook
	err := c.BindJSON(&gotWebhook)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	fmt.Println(gotWebhook.String())
	var IDs = make([]int, 0)
	for _, recepient := range config.Config.Recepient {
		id, err := strconv.Atoi(recepient)
		if err != nil {
			continue
		}
		IDs = append(IDs, id)
	}
	service := dao.NewClientService(dao.NewClientGorm(config.Config.DB))
	var userid int
	for _, message := range gotWebhook.Messages {
		userid = message.UserId
	}
	user, err := service.Dao.GetByUserId(uint(userid))
	if err != nil {
		for _, id := range IDs {
			_, err := config.Config.Bot.Send(tele.ChatID(id), err.Error())
			if err != nil {
				log.Println(err)
			}
			log.Println(err)
		}

		return
	}
	answer := "Получено сообщение:\nАккаунт: " + user.Name + " " + user.ProfileUrl + "\n" + gotWebhook.String()
	for _, id := range IDs {
		_, err := config.Config.Bot.Send(tele.ChatID(id), answer)
		if err != nil {
			log.Println(err)
		}
		c.JSON(http.StatusOK, "ok")
	}

}
func GetServer(c *gin.Context) {

	c.JSON(http.StatusOK, "Got server")
}
