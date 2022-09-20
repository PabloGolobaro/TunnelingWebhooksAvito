package ngrok

type WebhookSubscribeRequestBody struct {
	// Url на который будут отправляться нотификации
	Addr  string `json:"addr"`
	Proto string `json:"proto"`
	Name  string `json:"name"`
}
