package crud

import (
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

func (a *CrudToken) TokenTimeCode(secreat string) string {

	times := ""

	// Adicionando os valores do map à lista
	for uuid, value := range a.appService.MapToken {
		item := model.ToTokenResponse(value, passwrd)
		item.Id = uuid
		listToken = append(listToken, item)
	}

	return model.NewMessage(true, listToken).ToJSON()
}

func (a *CrudToken) TokenInfo(uid, pass string) string {

	token := a.appService.MapToken[uid]

	if !token.IsInterfaceNil() {
		return model.NewMessage(false, nil).ToJSON()
	}

	item := model.ToTokenResponse(token, pass)
	item.Id = uid

	return model.NewMessage(true, item).ToJSON()
}

func (a *CrudToken) TokenCreate(res string) string {

	request := model.ToTokenRequest(res)

	a.appService.AddToken(request)

	return model.NewMessage(true, nil).ToJSON()
}

func (a *CrudToken) TokenUpdate(uid, res, pass string) string {

	request := model.ToTokenRequest(res)

	if len(request.Passwrd) < 1 {
		request.Passwrd = pass
	}

	a.appService.UpdateToken(uid, request, pass)

	return model.NewMessage(true, nil).ToJSON()
}
func (a *CrudToken) TokenDelete(uid string) string {

	a.appService.RemoveToken(uid)

	return model.NewMessage(true, nil).ToJSON()
}

func Build(service *service.ApplicationService) *CrudToken {
	return &CrudToken{
		appService: service,
	}
}
