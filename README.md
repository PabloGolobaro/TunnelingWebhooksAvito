# Tunneling Webhooks from Avito
 
 System sets up webhooks using AvitoAPI and ngrok tunnel for each one automaticaly.
 Therefore service accepts webhook messages through ngrok tunnels and sends specific webhook information to recipients via TelegramBot.
 - GORM
 - Gin router
 - Telebot for TG_Bot
 - PostgreSQL as DB
 - Ngrok agent as tunnel system
 - Docker-compose to deploy
