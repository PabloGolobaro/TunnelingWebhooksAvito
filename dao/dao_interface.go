package dao

import (
	"AvitoMessengeBot/api/auth"
	"AvitoMessengeBot/api/user"
	"AvitoMessengeBot/db"
)

type ClientDao interface {
	Create(client db.Client) error
	ReadById(id string) (db.Client, error)
	AppendClientUser(client db.Client, info user.UserInfoSelf) (db.Client, error)
	ReadAll() ([]db.Client, error)
	CheckIfTokenExpired(id string) (bool, error)
	UpdateToken(client db.Client, response200 auth.InlineResponse200) db.Client
	Delete(id string) error
	DeleteWebhook(client db.Client) (db.Client, error)
	AppendWebhook(client db.Client, webhook db.Webhook) (db.Client, error)
	GetByUserId(id uint) (db.User, error)
}
type ClientService struct {
	Dao ClientDao
}

func NewClientService(dao ClientDao) *ClientService {
	return &ClientService{Dao: dao}
}
