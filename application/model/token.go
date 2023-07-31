package model

import (
	"fmt"
	"log"
	"myauth/application/util"
	"net/url"
)

type TokenRequest struct {
	Name     string `json:"name"`
	Algoritm string `json:"algoritm"`
	Secret   string `json:"secret"`
	Passwrd  string `json:"passwrd"`
	Url      string `json:"url"`
}

type TokenResponse struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Algoritm string `json:"algoritm"`
	Secret   string `json:"secret"`
	Code     string `json:"code"`
	Url      string `json:"url"`
}

func ToTokenResponse(token util.Token, passwrd string) TokenResponse {

	secret, decript := util.Decrypt(token.Algoritm, token.Payload, passwrd, token.Name)
	dataCode := util.DataCode{
		Code: "****",
	}

	if decript {
		dataCode = util.ReadOTPInPayloadToDataCode(secret)
	}

	return TokenResponse{
		Name:     token.Name,
		Algoritm: token.Algoritm,
		Secret:   "****",
		Code:     dataCode.Code,
	}
}

func ToTokenRequest(data string) TokenRequest {

	request := TokenRequest{}
	UnMarshal(data, &request)

	if len(request.Url) > 0 {
		parsedURL, err := url.Parse(request.Url)
		if err != nil {
			log.Fatal(err)
		}
		parsedURL.Query().Set("secret", request.Secret)
		request.Url = parsedURL.String()
	} else {
		request.Url = fmt.Sprintf("otpauth://totp/MyAuth:user?secret=%s&issuer=%s", request.Secret, request.Name)
	}

	return request
}
