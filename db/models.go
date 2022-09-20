package db

import "time"

type Client struct {
	ID           uint    `gorm:"primaryKey"`
	ClientId     string  `json:"client_id"`
	ClientSecret string  `json:"client_secret"`
	AccessToken  string  `json:"access_token,omitempty"`
	ExpiresIn    float64 `json:"expires_in,omitempty"`
	ExpiresDate  time.Time
	User         User    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Webhook      Webhook `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
type User struct {
	ID         uint   `json:"id" gorm:"primaryKey"`
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	ProfileUrl string `json:"profile_url"`
	Email      string `json:"email"`
	ClientID   uint
}
type Webhook struct {
	ID          uint `gorm:"primaryKey"`
	PublicUrl   string
	Tunnel_name string
	ClientID    uint
}
