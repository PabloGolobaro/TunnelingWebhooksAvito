package main

import (
	"AvitoMessengeBot/api"
	config "AvitoMessengeBot/config"
	"AvitoMessengeBot/core"
	"AvitoMessengeBot/dao"
	"AvitoMessengeBot/db"
	"AvitoMessengeBot/server"
	"AvitoMessengeBot/tg_bot"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

const server_port string = ":8080"

func main() {
	db_type := flag.String("db", "sqlite", "Type of DB to use.")
	flag.Parse()
	//Load Configuration
	err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	// Open DB
	log.Println("Opening db...")
	config.Config.DB = db.Init(*db_type)
	r := server.NewGinRouter()
	// Create TG_BOT
	config.Config.Bot, err = tg_bot.NewTeleBot()
	if err != nil {
		log.Fatal(err)
	}
	//Create api_cient
	newConfiguration := api.NewConfiguration()
	apiClient := api.NewAPIClient(newConfiguration)
	clientService := dao.NewClientService(dao.NewClientGorm(config.Config.DB))
	// Make SigChanel
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		log.Println("Starting bot...")
		err = tg_bot.SendToAdmins(config.Config.Bot, config.Config.Admins)
		if err != nil {
			log.Println(err)
		}
		config.Config.Bot.Start()
	}()
	go func() {
		log.Println("Starting server...")
		r.Run(server_port)
	}()
	err = core.StartApp(clientService, apiClient)
	if err != nil {
		fmt.Println(err)
	}
	sig := <-shutdown
	fmt.Println("Signal - ", sig)
	err = core.StopApp(clientService, apiClient)
	if err != nil {
		fmt.Println(err)
	}

}
