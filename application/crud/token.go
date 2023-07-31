package crud

import (
	"fmt"
	"myauth/application/model"
	"myauth/application/service"
)

type CrudToken struct {
	appService *service.ApplicationService
}

// Crud Token
func (a *CrudToken) TokenList(passwrd string) string {

	listToken := []model.TokenResponse{}

	// Adicionando os valores do map à lista
	for uuid, value := range a.appService.MapToken {
		item := model.ToTokenResponse(value, passwrd)
		item.Id = uuid
		listToken = append(listToken, item)
	}

	return model.NewMessage(true, listToken).ToJSON()
}
func (a *CrudToken) TokenCreate(res string) string {

	fmt.Println(res)
	request := model.ToTokenRequest(res)
	fmt.Println(request.Name)
	fmt.Println(request.Secret)
	fmt.Println(request.Url)
	fmt.Println(request.Algoritm)
	a.appService.AddToken(request)

	return model.NewMessage(true, nil).ToJSON()
}

func (a *CrudToken) TokenUpdate() string {
	return ""
}
func (a *CrudToken) TokenDelete() string {
	return ""
}

func Build(service *service.ApplicationService) *CrudToken {
	return &CrudToken{
		appService: service,
	}
}
