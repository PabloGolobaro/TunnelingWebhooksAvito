package core

import (
	"AvitoMessengeBot/api"
	"AvitoMessengeBot/dao"
	"AvitoMessengeBot/db"
	"context"
	"github.com/antihax/optional"
	"log"
	"time"
)

const grant_type string = "client_credentials"

func StartApp(service *dao.ClientService, apiClient *api.APIClient) error {
	clients, err := service.Dao.ReadAll()
	if err != nil {
		return err
	}
	for _, client := range clients {
		if client.ExpiresDate.Before(time.Now()) {
			opts := &api.AccessApiGetAccessTokenOpts{
				ClientId:     optional.NewString(client.ClientId),
				ClientSecret: optional.NewString(client.ClientSecret),
				GrantType:    optional.NewString(grant_type),
			}
			response200, _, err := apiClient.AccessApi.GetAccessToken(context.Background(), opts)
			if err != nil {
				return err
			}
			client = service.Dao.UpdateToken(client, response200)
		}
		if client.User.ProfileUrl == "" {
			userInfoSelf, _, err := apiClient.UserApi.GetUserInfoSelf(context.Background(), client.AccessToken)
			if err != nil {
				log.Println(err)
			}
			client, err = service.Dao.AppendClientUser(client, userInfoSelf)
			if err != nil {
				return err
			}

		}
		if client.Webhook.PublicUrl != "" {
			UpsetWebhook(apiClient, client.Webhook.PublicUrl, client.AccessToken)
			client, err = service.Dao.DeleteWebhook(client)
			if err != nil {
				log.Println(err)
			}
		}
		publicUrl := StartTunel(apiClient, client.ClientId)
		SetWebhook(apiClient, publicUrl, client.AccessToken)
		webhook := db.Webhook{PublicUrl: publicUrl, Tunnel_name: client.ClientId}
		client, err = service.Dao.AppendWebhook(client, webhook)
		if err != nil {
			return err
		}
	}
	return nil
}

func StopApp(service *dao.ClientService, apiClient *api.APIClient) error {
	clients, err := service.Dao.ReadAll()
	if err != nil {
		return err
	}
	for _, client := range clients {
		if client.ExpiresDate.Before(time.Now()) {
			opts := &api.AccessApiGetAccessTokenOpts{
				ClientId:     optional.NewString(client.ClientId),
				ClientSecret: optional.NewString(client.ClientSecret),
				GrantType:    optional.NewString(grant_type),
			}
			response200, _, err := apiClient.AccessApi.GetAccessToken(context.Background(), opts)
			if err != nil {
				return err
			}
			client = service.Dao.UpdateToken(client, response200)
		}

		UpsetWebhook(apiClient, client.Webhook.PublicUrl, client.AccessToken)
		StopTunnel(apiClient, client.ClientId)
		client, err = service.Dao.DeleteWebhook(client)
		if err != nil {
			return err
		}

	}
	return nil
}
func StartNewClient(service *dao.ClientService, apiClient *api.APIClient, client db.Client) error {
	opts := &api.AccessApiGetAccessTokenOpts{
		ClientId:     optional.NewString(client.ClientId),
		ClientSecret: optional.NewString(client.ClientSecret),
		GrantType:    optional.NewString(grant_type),
	}
	response200, _, err := apiClient.AccessApi.GetAccessToken(context.Background(), opts)
	if err != nil {
		return err
	}
	client = service.Dao.UpdateToken(client, response200)

	if client.User.ProfileUrl == "" {
		userInfoSelf, _, err := apiClient.UserApi.GetUserInfoSelf(context.Background(), client.AccessToken)
		if err != nil {
			log.Println(err)
		}
		client, err = service.Dao.AppendClientUser(client, userInfoSelf)
		if err != nil {
			return err
		}

	}
	if client.Webhook.PublicUrl != "" {
		UpsetWebhook(apiClient, client.Webhook.PublicUrl, client.AccessToken)
		client, err = service.Dao.DeleteWebhook(client)
		if err != nil {
			log.Println(err)
		}
	}
	publicUrl := StartTunel(apiClient, client.ClientId)
	SetWebhook(apiClient, publicUrl, client.AccessToken)
	webhook := db.Webhook{PublicUrl: publicUrl, Tunnel_name: client.ClientId}
	client, err = service.Dao.AppendWebhook(client, webhook)
	if err != nil {
		return err
	}
	return nil
}
