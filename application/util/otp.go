package util

import (
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/pquerna/otp/totp"
)

const otp_url_aws string = "otpauth://totp/Amazon%20Web%20Services:teste4534@gilbertoramos?secret=GQMWATC3JYIGQBZRDPT7JEVWVDZQISEOXZXUP42552H6VRJTIXM3YLHI37J3YW6N&issuer=Amazon%20Web%20Services"
const otp_url_github string = "otpauth://totp/GitHub:gilberto-009199?secret=VNZH25YKY7W2HWS5&issuer=GitHub"

func readOTP() {

	otp_url := otp_url_github

	parsedURL, err := url.Parse(otp_url)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("scheme: %s\n", parsedURL.Scheme)
	fmt.Printf("host: %s\n", parsedURL.Host)
	fmt.Printf("path: %s\n", parsedURL.Path)
	fmt.Println("query args:")
	for key, values := range parsedURL.Query() {
		fmt.Printf("  %s = %s\n", key, values[0])
	}
	secret := parsedURL.Query().Get("secret")
	fmt.Printf("  secret = %s \n", secret)

	key, err := totp.GenerateCode(secret, time.Now())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("  code = %s \n", key)

}
