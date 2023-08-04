package service

import (
	"fmt"
	"log"
	"myauth/application/model"
	"myauth/application/util"

	"path/filepath"

	"github.com/google/uuid"
)

type ApplicationService struct {
	Settings util.Settings
	MapToken map[string]util.Token
}

var settingsDefault = util.Settings{
	PathFileSettings: filepath.Join(util.GetDirUser(), ".myauth"),
	PathFileTokens:   filepath.Join(util.GetDirUser(), ".myauth.bin"),
	AlgoritmDefault:  "Camellia",
}

func Build() *ApplicationService {
	return &ApplicationService{}
}

// Start Application
func (s *ApplicationService) Start() {

	settings, er := util.ReadSettingsInFile(settingsDefault.PathFileSettings)
	if er != nil {
		fmt.Println(er)
		if util.SaveSettingsInFile(settingsDefault.PathFileSettings, settingsDefault) {
			settings, _ = util.ReadSettingsInFile(settingsDefault.PathFileSettings)
		} else {
			log.Fatal("Not create file config")
		}
	}

	if settings.PathFileSettings != settingsDefault.PathFileSettings {

		settingsExternal, er := util.ReadSettingsInFile(settings.PathFileSettings)

		if er != nil {
			if util.SaveSettingsInFile(settings.PathFileSettings, settingsDefault) {
				settingsExternal, _ = util.ReadSettingsInFile(settingsDefault.PathFileSettings)
			} else {
				log.Fatal("Not create file config")
			}
		}
		settings = settingsExternal
	}

	s.Settings = settings

	mapToken, er := util.ReadTokensInFile(s.Settings.PathFileTokens)
	if er != nil {

		fmt.Println(er)

		if len(s.MapToken) < 1 {
			mapToken = map[string]util.Token{}
		}

		if util.SaveTokensInFile(s.Settings.PathFileTokens, mapToken) {
			mapToken, _ = util.ReadTokensInFile(s.Settings.PathFileTokens)
		} else {
			log.Fatal("Not create file 	Tokens")
		}
	}

	s.MapToken = mapToken

	fmt.Println(s.MapToken)

}

func (s *ApplicationService) SetSettings(settings model.SettingsRequest) {

	settingsNow := util.Settings{
		PathFileSettings: settings.PathFileSettings,
		PathFileTokens:   settings.PathFileTokens,
		AlgoritmDefault:  settings.AlgoritmDefault,
	}

	util.SaveSettingsInFile(settingsDefault.PathFileSettings, settingsNow)

	util.SaveSettingsInFile(settings.PathFileSettings, settingsNow)

	util.SaveTokensInFile(settings.PathFileTokens, s.MapToken)

	s.Settings = settingsNow

}

func (s *ApplicationService) AddToken(token model.TokenRequest) {
	id := uuid.New()

	payload := util.Encrypt(token.Algoritm, token.Url, token.Passwrd, token.Name)

	tokenStorage := util.Token{
		Name:     token.Name,
		Algoritm: token.Algoritm,
		Payload:  payload,
	}

	s.MapToken[id.String()] = tokenStorage

	fmt.Println(s.MapToken)

	go util.SaveTokensInFile(s.Settings.PathFileTokens, s.MapToken)

}

func (s *ApplicationService) UpdateToken(uid string, token model.TokenRequest, pass string) {

	tokenCurrent := s.MapToken[uid]

	payloadCurrent, decript := util.Decrypt(tokenCurrent.Algoritm, tokenCurrent.Payload, pass, tokenCurrent.Name)
	if !decript {
		fmt.Println("FALED Decript")
	}

	//	fmt.Printf("Decript: %s\n", payloadCurrent)

	payloadNow := util.Encrypt(token.Algoritm, payloadCurrent, token.Passwrd, token.Name)

	tokenStorage := util.Token{
		Name:     token.Name,
		Algoritm: token.Algoritm,
		Payload:  payloadNow,
	}

	s.MapToken[uid] = tokenStorage

	fmt.Println(s.MapToken)

	go util.SaveTokensInFile(s.Settings.PathFileTokens, s.MapToken)

}

func (s *ApplicationService) RemoveToken(uid string) {

	delete(s.MapToken, uid)

	go util.SaveTokensInFile(s.Settings.PathFileTokens, s.MapToken)

}
