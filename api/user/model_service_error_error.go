/*
 * Информация о пользователе
 *
 * API для получения баланса кошелька пользователя, истории операций и инфорации об авторизованном пользователе **Авито API для бизнеса предоставляется согласно [Условиям использования](https://api.avito.ru/docs/public/APITermsOfServiceV1.pdf).**
 *
 * API version: 1
 * Contact: supportautoload@avito.ru
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package user

type ServiceErrorError struct {
	// Код ошибки
	Code int32 `json:"code"`
	// Описание ошибки
	Message string `json:"message"`
}