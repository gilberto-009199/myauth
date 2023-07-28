package util

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/mgo.v2/bson"
)

var prefix_name = []byte("myauth.gilberto-009199.github.com")
var prefix_version = []byte("001")
var prefix = append(prefix_name, prefix_version...)

var sizeBytesFromPrefixName = len(prefix_name)
var sizeBytesFromPrefixVersion = len(prefix_version)

type Settings struct {
	PathFileSettings string
	PathFileTokens   string
	AlgoritmDefault  string
}

type Token struct {
	Name     string
	Algoritm string
	Payload  string
}

func ReadTokensInFile(filename string) (map[string]Token, error) {

	vectToken := map[string]Token{}

	bytes, er := readInFile(filename)
	if er != nil {
		fmt.Println(er)
		return nil, er
	}

	data := bytes[sizeBytesFromPrefixName+sizeBytesFromPrefixVersion:]

	readInArrayToObject(data, &vectToken)

	return vectToken, nil
}

func ReadSettingsInFile(filename string) (Settings, error) {

	settings := Settings{}

	bytes, er := readInFile(filename)
	if er != nil {
		fmt.Println(er)
		return settings, er
	}

	data := bytes[sizeBytesFromPrefixName+sizeBytesFromPrefixVersion:]

	readInArrayToObject(data, &settings)

	return settings, nil
}

func SaveSettingsInFile(filename string, setting Settings) bool {

	bytesFromSettings := toBSON(setting)

	er := saveInFile(filename, append(prefix, bytesFromSettings...))

	if er != nil {
		fmt.Println(er)
		return false
	}

	return true
}

func SaveTokensInFile(filename string, listToken map[string]Token) bool {
	bytesFromListToken := toBSON(listToken)

	er := saveInFile(filename, append(prefix, bytesFromListToken...))

	if er != nil {
		fmt.Println(er)
		return false
	}

	return true
}

// UTil Bson And File's
func toBSON(ent interface{}) []byte {
	data, err := bson.Marshal(ent)
	if err != nil {
		fmt.Println(err)
	}
	return data
}

func readInArrayToObject(byteArray []byte, obj interface{}) {
	if err := bson.Unmarshal(byteArray, obj); err != nil {
		fmt.Println(err)
	}
}

func saveInFile(filename string, data []byte) error {
	erro := ioutil.WriteFile(filename, data, 0644)
	return erro
}

func readInFile(filename string) ([]byte, error) {

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return content, nil
}
func GetDirUser() string {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return dirname
}
