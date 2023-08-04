package util

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/mgo.v2/bson"
)

var Prefix_name = []byte("myauth.gilberto-009199.github.com")
var Prefix_version = []byte("001")
var Prefix = append(Prefix_name, Prefix_version...)

var sizeBytesFromPrefixName = len(Prefix_name)
var sizeBytesFromPrefixVersion = len(Prefix_version)

type Settings struct {
	PathFileSettings string `json:"file_settings"`
	PathFileTokens   string `json:"file_tokens"`
	AlgoritmDefault  string `json:"algoritm"`
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

	bytesFromSettings := ToBSON(setting)

	er := SaveInFile(filename, append(Prefix, bytesFromSettings...))

	if er != nil {
		fmt.Println(er)
		return false
	}

	return true
}

func SaveTokensInFile(filename string, listToken map[string]Token) bool {
	bytesFromListToken := ToBSON(listToken)

	er := SaveInFile(filename, append(Prefix, bytesFromListToken...))

	if er != nil {
		fmt.Println(er)
		return false
	}

	return true
}

// UTil Bson And File's
func ToBSON(ent interface{}) []byte {
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

func SaveInFile(filename string, data []byte) error {
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

func (f *Token) IsInterfaceNil() bool {
	if nil == f {
		return true
	}
	return false
}
