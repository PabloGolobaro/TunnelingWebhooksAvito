/*
 * Мессенджер
 *
 * API Мессенджера - набор методов для получения списка чатов пользователя на Авито, получения сообщений в чате, отправки сообщения в чат и другие Через API Мессенджера можно организовать интеграцию между мессенджером Авито и сторонней системой в обе стороны  **Авито API для бизнеса предоставляется согласно [Условиям использования](https://api.avito.ru/docs/public/APITermsOfServiceV1.pdf).**
 *
 * API version: 1
 * Contact: supportautoload@avito.ru
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package message

type ChatPublicUserProfileAvatarImages struct {
	Var128x128 string `json:"128x128,omitempty"`
	Var192x192 string `json:"192x192,omitempty"`
	Var24x24   string `json:"24x24,omitempty"`
	Var256x256 string `json:"256x256,omitempty"`
	Var36x36   string `json:"36x36,omitempty"`
	Var48x48   string `json:"48x48,omitempty"`
	Var64x64   string `json:"64x64,omitempty"`
	Var72x72   string `json:"72x72,omitempty"`
	Var96x96   string `json:"96x96,omitempty"`
}