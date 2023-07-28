package crud

import (
	"myauth/application/model"
	"myauth/application/service"
	"myauth/application/util"
)

type CrudToken struct {
	appService *service.ApplicationService
}

type TokenRequest struct {
	Name     string `json:name`
	Algoritm string `json:algoritm`
	Secret   string `json:secret`
	Passwrd  string `json:passwrd`
}

type TokenResponse struct {
	Name     string `json:name`
	Algoritm string `json:algoritm`
	Secret   string `json:secret`
	Code     string `json:code`
}

// Crud Token
func (a *CrudToken) TokenList(passwrd string) string {

	listToken := []TokenResponse{}

	// Adicionando os valores do map à lista
	for _, value := range a.appService.MapToken {
		listToken = append(listToken, toTokenResponse(value, passwrd))
	}

	return model.NewMessage(true, listToken).ToJSON()
}
func (a *CrudToken) TokenCreate() string {
	return ""
}
func (a *CrudToken) TokenUpdate() string {
	return ""
}
func (a *CrudToken) TokenDelete() string {
	return ""
}

func toTokenResponse(token util.Token, passwrd string) TokenResponse {

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

func Build(service *service.ApplicationService) *CrudToken {
	return &CrudToken{
		appService: service,
	}
}
