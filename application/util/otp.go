package util

import (
	"log"
	"net/url"
	"time"

	"github.com/pquerna/otp/totp"
)

const otp_url_aws string = "otpauth://totp/Amazon%20Web%20Services:teste4534@gilbertoramos?secret=GQMWATC3JYIGQBZRDPT7JEVWVDZQISEOXZXUP42552H6VRJTIXM3YLHI37J3YW6N&issuer=Amazon%20Web%20Services"
const otp_url_github string = "otpauth://totp/GitHub:gilberto-009199?secret=VNZH25YKY7W2HWS5&issuer=GitHub"

type DataCode struct {
	Code string
}

func ReadOTPInPayloadToDataCode(otp string) DataCode {

	otp_url := otp

	data := DataCode{
		Code: "",
	}

	parsedURL, err := url.Parse(otp_url)
	if err != nil {
		log.Fatal(err)
		return data
	}

	secret := parsedURL.Query().Get("secret")

	timeCurrent := time.Now()
	code, err := totp.GenerateCode(secret, timeCurrent)
	if err != nil {
		log.Fatal(err)
	}
	data.Code = code

	return data
}

// gil por que vc nao simplesmente usa gson ou o jsonp padrao do Go? Em?

func ReadOTPInURLToJSON(otp string) (*url.URL, error) {

	otp_url := otp

	parsedURL, err := url.Parse(otp_url)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	/*
		message += fmt.Sprintf(`"scheme": "%s",`, parsedURL.Scheme)
		message += fmt.Sprintf(`"host": "%s",`, parsedURL.Host)
		message += fmt.Sprintf(`"path": "%s",`, parsedURL.Path)

		for key, values := range parsedURL.Query() {
			message += fmt.Sprintf(`"%s": "%s",`, key, values[0])
		}

		message += fmt.Sprintf(`"url":"%s"}`, otp_url)
	*/
	return parsedURL, nil
}
