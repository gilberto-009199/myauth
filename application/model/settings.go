package model

type SettingsRequest struct {
	PathFileSettings string `json:"file_settings"`
	PathFileTokens   string `json:"file_tokens"`
	AlgoritmDefault  string `json:"algoritm"`
}

func ToSettingRequest(data string) SettingsRequest {

	request := SettingsRequest{}
	UnMarshal(data, &request)

	return request

}
