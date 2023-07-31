package model

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	Status  bool        `json:"status"`
	Message interface{} `json:"message"`
}

func NewMessage(status bool, message interface{}) *Message {
	return &Message{
		Status:  status,
		Message: message,
	}
}
func (m *Message) ToJSON() string {
	json, er := marshal(m)
	if er != nil {
		fmt.Println(er)
		return "{status:false}"
	}

	return json
}

func UnMarshal(payload string, dataInterface interface{}) error {
	return json.Unmarshal([]byte(payload), dataInterface)
}

func unMarshal(payload string, dataInterface interface{}) error {
	return json.Unmarshal([]byte(payload), dataInterface)
}

func marshal(dataInterface interface{}) (string, error) {
	byteJSON, er := json.Marshal(dataInterface)
	if er != nil {
		fmt.Println(er)
		return "{}", er
	}

	return string(byteJSON[:]), nil
}
