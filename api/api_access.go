/*
 * Авторизация
 *
 * Получение и обновление авторизационных токенов для персональной авторизации и авторизации приложений **Авито API для бизнеса предоставляется согласно [Условиям использования](https://api.avito.ru/docs/public/APITermsOfServiceV1.pdf).**
 *
 * API version: 1
 * Contact: supportautoload@avito.ru
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package api

import (
	"AvitoMessengeBot/api/auth"
	"context"
	"github.com/antihax/optional"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Linger please
var (
	_ context.Context
)

type AccessApiService service

/*
AccessApiService Получение access token
Получения временного ключа для авторизации
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *AccessApiGetAccessTokenOpts - Optional Parameters:
     * @param "ClientId" (optional.String) -
     * @param "ClientSecret" (optional.String) -
     * @param "GrantType" (optional.String) -
@return InlineResponse200
*/

type AccessApiGetAccessTokenOpts struct {
	ClientId     optional.String
	ClientSecret optional.String
	GrantType    optional.String
}

func (a *AccessApiService) GetAccessToken(ctx context.Context, localVarOptionals *AccessApiGetAccessTokenOpts) (auth.InlineResponse200, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue auth.InlineResponse200
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/token"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.ClientId.IsSet() {
		localVarFormParams.Add("client_id", parameterToString(localVarOptionals.ClientId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ClientSecret.IsSet() {
		localVarFormParams.Add("client_secret", parameterToString(localVarOptionals.ClientSecret.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.GrantType.IsSet() {
		localVarFormParams.Add("grant_type", parameterToString(localVarOptionals.GrantType.Value(), ""))
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v auth.InlineResponse200
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}