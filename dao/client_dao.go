package dao

import (
	"AvitoMessengeBot/api/auth"
	"AvitoMessengeBot/api/user"
	"AvitoMessengeBot/db"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type ClientGorm struct {
	DB *gorm.DB
}

func NewClientGorm(DB *gorm.DB) *ClientGorm {
	return &ClientGorm{DB: DB}
}

func (u ClientGorm) Create(client db.Client) error {
	result := u.DB.Create(&client)
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected == 0 {
		return gorm.ErrNotImplemented
	}
	return nil
}

func (u ClientGorm) ReadById(id string) (db.Client, error) {
	var client db.Client
	db := u.DB.Where("client_id=?", id).Preload("User").Preload("Webhook").First(&client)
	if db.Error != nil {
		return client, db.Error
	}
	return client, nil
}
func (u ClientGorm) AppendClientUser(client db.Client, info user.UserInfoSelf) (db.Client, error) {
	user := db.User{
		ID:         uint(info.Id),
		Name:       info.Name,
		Phone:      info.Phone,
		ProfileUrl: info.ProfileUrl,
		Email:      info.Email,
	}
	association := u.DB.Model(&client).Association("User")
	association.Append(&user)
	if association.Error != nil {
		return client, association.Error
	}
	return client, nil
}

func (u ClientGorm) CheckIfTokenExpired(id string) (bool, error) {
	client, err := u.ReadById(id)
	if err != nil {
		return false, err
	}
	if client.ExpiresDate.Before(time.Now()) {
		return true, nil
	} else {
		return false, nil
	}
}

func (u ClientGorm) UpdateToken(client db.Client, response200 auth.InlineResponse200) db.Client {
	client.AccessToken = "Bearer " + response200.AccessToken
	client.ExpiresIn = response200.ExpiresIn
	dur := time.Duration(response200.ExpiresIn)
	dur *= time.Second
	client.ExpiresDate = time.Now().Add(dur)
	u.DB.Save(&client)
	return client
}

func (u ClientGorm) Delete(id string) error {
	var client db.Client
	tx := u.DB.Where("client_id=?", id).First(&client)
	if tx.Error != nil {
		return tx.Error
	}
	u.DB.Delete(&client)
	return nil

}
func (u ClientGorm) ReadAll() ([]db.Client, error) {
	var clients []db.Client
	tx := u.DB.Preload("User").Preload("Webhook").Find(&clients)
	fmt.Println(tx.RowsAffected)
	if tx.Error != nil {
		return clients, tx.Error
	}
	return clients, nil
}
func (u ClientGorm) AppendWebhook(client db.Client, webhook db.Webhook) (db.Client, error) {
	err := u.DB.Model(&client).Association("Webhook").Append(&webhook)
	if err != nil {
		return client, err
	}
	return client, nil
}
func (u ClientGorm) DeleteWebhook(client db.Client) (db.Client, error) {
	webhookId := client.Webhook.ID
	err := u.DB.Model(&client).Association("Webhook").Clear()
	if err != nil {
		return client, err
	}
	u.DB.Where("id = ?", webhookId).Delete(&db.Webhook{})
	return client, nil
}

func (u ClientGorm) GetByUserId(id uint) (db.User, error) {
	var user db.User
	db := u.DB.First(&user, id)
	if db.Error != nil {
		return user, db.Error
	}
	return user, nil
}
