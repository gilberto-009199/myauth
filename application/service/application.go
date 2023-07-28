package service

import (
	"fmt"
	"log"
	"myauth/application/util"

	"path/filepath"
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
