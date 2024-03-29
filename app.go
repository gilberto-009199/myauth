package main

import (
	"context"
	"fmt"
	"myauth/application/model"
	"myauth/application/service"
	"myauth/application/util"
	"strings"

	"github.com/skip2/go-qrcode"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx        context.Context
	appService *service.ApplicationService
}

func (a *App) ListAlgoritm() string {
	return model.NewMessage(true, util.LIST_ALGORITM).ToJSON()
}

func (a *App) GetSettings() string {
	return model.NewMessage(true, a.appService.Settings).ToJSON()
}

func (a *App) GetFile(pattern string) string {

	file, e := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Select File",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "File (" + pattern + ")",
				Pattern:     pattern,
			},
		},
	})

	if e != nil {
		fmt.Println(e)
		return model.NewMessage(false, nil).ToJSON()
	}

	fmt.Println(file)
	return model.NewMessage(true, file).ToJSON()
}

func (a *App) getFileInternal(pattern string) (string, error) {

	file, e := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Select File",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "File (" + pattern + ")",
				Pattern:     pattern,
			},
		},
	})

	if e != nil {
		fmt.Println(e)
		return "", e
	}

	fmt.Println(file)
	return file, nil
}

func (a *App) ExportToken(uid, exportType, pass string) string {

	token := a.appService.MapToken[uid]

	payload, decript := util.Decrypt(token.Algoritm, token.Payload, pass, token.Name)
	if !decript {
		fmt.Println("FALED Decript")
		return model.NewMessage(false, nil).ToJSON()
	}

	var file string
	var buffFile []byte
	var e error

	switch exportType {
	case "csv":
		file, e = a.SaveFileDialog("token.csv")
		if e != nil {
			fmt.Println(e)
			return model.NewMessage(false, nil).ToJSON()
		}
		buffFile = []byte(token.Name + ";" + token.Algoritm + ";" + payload)
		break
	case "qrcode":
		file, e = a.SaveFileDialog("token.png")
		if e != nil {
			fmt.Println(e)
			return model.NewMessage(false, nil).ToJSON()
		}
		buffFile, e = qrcode.Encode(payload, qrcode.Low, 256)
		if e != nil {
			fmt.Println(e)
		}
		break
	case "myauth":
		file, e = a.SaveFileDialog("token.bin")
		if e != nil {
			fmt.Println(e)
			return model.NewMessage(false, nil).ToJSON()
		}

		MapToken := map[string]util.Token{uid: token}
		buffFile = []byte(append(util.Prefix, util.ToBSON(MapToken)...))
		break
	}

	e = util.SaveInFile(file, buffFile)
	if e != nil {
		fmt.Println(e)
	}

	return model.NewMessage(true, file).ToJSON()
}

func (a *App) ImportToken(exportType, pass string) string {

	var file string
	var e error
	var buffFile []byte

	switch exportType {
	case "csv":
		file, e = a.getFileInternal("*.csv")
		if e != nil {
			fmt.Println(e)
			return model.NewMessage(false, nil).ToJSON()
		}
		buffFile, e = util.ReadInFile(file)
		if e != nil {
			fmt.Println(e)
			return model.NewMessage(false, nil).ToJSON()
		}

		csv := string(buffFile)

		vctString := strings.Split(csv, ";")

		a.appService.AddToken(model.TokenRequest{
			Name:     vctString[0],
			Algoritm: vctString[1],
			Url:      vctString[2],
			Passwrd:  pass,
		})

		break

	case "qrcode":
		file, e = a.getFileInternal("*.png;*.jpg;*.jpeg")
		if e != nil {
			fmt.Println(e)
			return model.NewMessage(false, nil).ToJSON()
		}

		url, e := util.FileQRCode(file)
		if e != nil {
			fmt.Println(e)
			return model.NewMessage(false, nil).ToJSON()
		}

		// Cifrar
		fmt.Println(url.RawPath)
		a.appService.AddToken(model.TokenRequest{
			Name:     url.Query().Get("issuer"),
			Algoritm: a.appService.Settings.AlgoritmDefault,
			Url:      url.RawPath,
			Passwrd:  pass,
		})

		break

	case "myauth":
		file, e = a.getFileInternal("*.bin;*.myauth")
		if e != nil {
			fmt.Println(e)
			return model.NewMessage(false, nil).ToJSON()
		}

		mapToken, e := util.ReadTokensInFile(file)
		if e != nil {
			fmt.Println(e)
			return model.NewMessage(false, nil).ToJSON()
		}

		// Adicionando os valores do map à lista
		for _, value := range mapToken {
			a.appService.AddTokenNotPassword(model.TokenRequest{
				Name:     value.Name,
				Algoritm: value.Algoritm,
				Url:      value.Payload,
			})
		}

		break
	}

	return model.NewMessage(true, file).ToJSON()
}

func (a *App) SaveFileDialog(filename string) (string, error) {

	file, e := runtime.SaveFileDialog(a.ctx, runtime.SaveDialogOptions{
		DefaultFilename: filename,
	})

	if e != nil {
		fmt.Println(e)
		return "", e
	}

	return file, nil
}

func (a *App) GetDiretory() string {
	file, e := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{})

	if e != nil {
		fmt.Println(e)
		return model.NewMessage(false, nil).ToJSON()
	}

	fmt.Println(file)
	return model.NewMessage(true, file).ToJSON()
}

func (a *App) SetSettings(settings string) string {

	data := model.ToSettingRequest(settings)

	a.appService.SetSettings(data)

	return model.NewMessage(true, a.appService.Settings).ToJSON()
}

// NewApp creates a new App application struct
func Build(service *service.ApplicationService) *App {
	return &App{
		appService: service,
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}
