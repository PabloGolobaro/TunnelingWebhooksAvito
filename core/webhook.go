package core

import (
	"AvitoMessengeBot/api"
	"AvitoMessengeBot/api/message"
	"context"
	"fmt"
	"github.com/antihax/optional"
	"log"
)

func GetUserInfo(apiClient *api.APIClient, auth_token string) {
	userInfoSelf, _, err := apiClient.UserApi.GetUserInfoSelf(context.Background(), auth_token)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(userInfoSelf.String())
}
func GetChats(apiClient *api.APIClient, id int64, auth_token string) {
	opts := &api.MessengerApiGetChatsOpts{
		UnreadOnly: optional.NewBool(false),
		Limit:      optional.NewInt32(3),
	}
	chats, _, err := apiClient.MessengerApi.GetChats(context.Background(), auth_token, id, opts)
	if err != nil {
		log.Println(err)
	}
	for _, chat := range chats.Chats {
		for i, user := range chat.Users {
			fmt.Println(i, " - ", user.Name, " ", user.PublicUserProfile.Url, " ", user.Id, "\n")
		}
	}

}
func SetWebhook(apiClient *api.APIClient, url string, auth_token string) {
	webhookOpts := &api.MessengerApiPostWebhookOpts{
		Body: optional.NewInterface(message.WebhookSubscribeRequestBody{Url: url}),
	}
	_, r, err := apiClient.MessengerApi.PostWebhook(context.Background(), auth_token, webhookOpts)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Webhook is set: ", r.Status)
}
func UpsetWebhook(apiClient *api.APIClient, url string, auth_token string) {
	webhookOpts := &api.MessengerApiPostWebhookUnsubscribeOpts{
		Body: optional.NewInterface(message.WebhookSubscribeRequestBody{Url: url}),
	}
	_, r, err := apiClient.MessengerApi.PostWebhookUnsubscribe(context.Background(), auth_token, webhookOpts)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Webhook is upset: ", r.Status)
}
