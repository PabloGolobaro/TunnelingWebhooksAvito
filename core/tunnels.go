package core

import (
	"AvitoMessengeBot/api"
	"AvitoMessengeBot/api/ngrok"
	"context"
	"fmt"
	"github.com/antihax/optional"
	"log"
)

func StartTunel(apiClient *api.APIClient, tunnel_name string) (public string) {
	requestBody := ngrok.WebhookSubscribeRequestBody{
		Addr:  "8080",
		Proto: "http",
		Name:  tunnel_name,
	}
	opt := &api.NgorkStartTunnelOpt{Body: optional.NewInterface(requestBody)}
	detail, r, err := apiClient.NgrokApi.StartTunnel(context.Background(), opt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(r.Status)
	fmt.Println(detail.Name, " - ", detail.PublicUrl, "\n")
	fmt.Println("Created Tunnel ", detail.Name)
	return detail.PublicUrl
}
func StopTunnel(apiClient *api.APIClient, name string) {
	r, err := apiClient.NgrokApi.StopTunnel(context.Background(), name)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(r.Status)
	fmt.Println("Deleted Tunnel ", name)
}
