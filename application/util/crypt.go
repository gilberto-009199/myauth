package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"crypto/rc4"
	"encoding/hex"
	"fmt"

	"github.com/aead/camellia"
	"golang.org/x/crypto/chacha20"
)

var LIST_ALGORITM = []string{"AES", "3DES", "RC4", "ChaCha20", "Camellia"}
var MAP_ALGORITM = map[string]int{"AES": 0, "3DES": 1, "RC4": 2, "ChaCha20": 3, "Camellia": 4}

func Encrypt(algoritm, payload, passwrd, nonce string) string {

	data := []byte(payload)
	key := []byte(passwrd)
	cipher := make([]byte, 0)

	// Odeio Switch com case recuado , danen-se seus fdp
	switch MAP_ALGORITM[algoritm] {
	case MAP_ALGORITM["AES"]:
		cipher, _ = encryptAES(data, key, []byte(nonce))
	case MAP_ALGORITM["3DES"]:
		cipher, _ = encrypt3DES(data, key)
	case MAP_ALGORITM["RC4"]:
		cipher = encryptRC4(data, key, []byte(nonce))
	case MAP_ALGORITM["ChaCha20"]:
		cipher = encryptChaCha20(data, key, []byte(nonce))
	case MAP_ALGORITM["Camellia"]:
		cipher, _ = encryptCamellia(data, key)
	}

	return hex.EncodeToString(cipher[:])
}

func Decrypt(algoritm, payload, passwrd, nonce string) (string, bool) {

	data, _ := hex.DecodeString(payload)
	key := []byte(passwrd)
	var cipher []byte

	// Odeio Switch com case recuado, danen-se seus fdp
	switch MAP_ALGORITM[algoritm] {
	case MAP_ALGORITM["AES"]:
		cipher, _ = decryptAES(data, key, []byte(nonce))
	case MAP_ALGORITM["3DES"]:
		cipher, _ = decrypt3DES(data, key)
	case MAP_ALGORITM["RC4"]:
		cipher = decryptRC4(data, key, []byte(nonce))
	case MAP_ALGORITM["ChaCha20"]:
		cipher = decryptChaCha20(data, key, []byte(nonce))
	case MAP_ALGORITM["Camellia"]:
		cipher, _ = decryptCamellia(data, key)
	default:
		return "", false
	}

	return string(cipher[:]), true
}

func encryptCamellia(data, key []byte) ([]byte, error) {

	keyBuffer := key
	lenKeyBuffer := len(keyBuffer)
	if lenKeyBuffer < 32 {
		complementBuffer := make([]byte, int(32-lenKeyBuffer))
		keyBuffer = append(keyBuffer, complementBuffer...)
	} else {
		keyBuffer = keyBuffer[:32]
	}

	block, err := camellia.NewCipher(keyBuffer)
	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, camellia.BlockSize+len(data))
	iv := ciphertext[:camellia.BlockSize]
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[camellia.BlockSize:], data)
	return ciphertext, nil
}

func decryptCamellia(data, key []byte) ([]byte, error) {

	keyBuffer := key
	lenKeyBuffer := len(keyBuffer)
	if lenKeyBuffer < 32 {
		complementBuffer := make([]byte, int(32-lenKeyBuffer))
		keyBuffer = append(keyBuffer, complementBuffer...)
	} else {
		keyBuffer = keyBuffer[:32]
	}

	block, err := camellia.NewCipher(keyBuffer)
	if err != nil {
		return nil, err
	}

	plaintext := make([]byte, camellia.BlockSize+len(data))
	iv := plaintext[:camellia.BlockSize]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(plaintext[camellia.BlockSize:], data)

	return plaintext[camellia.BlockSize*2:], nil

}

// ChaCha20
func encryptChaCha20(data, key, nonce []byte) []byte {

	keyBuffer := key
	lenKeyBuffer := len(keyBuffer)
	if lenKeyBuffer < chacha20.KeySize {
		complementBuffer := make([]byte, int(chacha20.KeySize-lenKeyBuffer))
		keyBuffer = append(keyBuffer, complementBuffer...)
	} else {
		keyBuffer = keyBuffer[:chacha20.KeySize]
	}

	nonceBuffer := nonce
	lenNonceBuffer := len(nonceBuffer)
	if lenNonceBuffer < 24 {
		complementBuffer := make([]byte, int(24-lenNonceBuffer))
		nonceBuffer = append(nonceBuffer, complementBuffer...)
	} else {
		nonceBuffer = nonceBuffer[:24]
	}

	cipher, e := chacha20.NewUnauthenticatedCipher(keyBuffer, nonceBuffer)
	if e != nil {
		fmt.Println(e)
	}
	ciphertext := make([]byte, len(data))
	cipher.XORKeyStream(ciphertext, data)
	return ciphertext
}

func decryptChaCha20(ciphertext, key, nonce []byte) []byte {
	keyBuffer := key
	lenKeyBuffer := len(keyBuffer)
	if lenKeyBuffer < chacha20.KeySize {
		complementBuffer := make([]byte, int(chacha20.KeySize-lenKeyBuffer))
		keyBuffer = append(keyBuffer, complementBuffer...)
	} else {
		keyBuffer = keyBuffer[:chacha20.KeySize]
	}

	nonceBuffer := nonce
	lenNonceBuffer := len(nonceBuffer)
	if lenNonceBuffer < 24 {
		complementBuffer := make([]byte, int(24-lenNonceBuffer))
		nonceBuffer = append(nonceBuffer, complementBuffer...)
	} else {
		nonceBuffer = nonceBuffer[:24]
	}

	cipher, e := chacha20.NewUnauthenticatedCipher(keyBuffer, nonceBuffer)
	if e != nil {
		fmt.Println(e)
	}

	plaintext := make([]byte, len(ciphertext))
	cipher.XORKeyStream(plaintext, ciphertext)
	return plaintext
}

// RC4 is Insecure
func encryptRC4(data, key, salt []byte) []byte {
	cipher, _ := rc4.NewCipher(append(key, salt...))
	ciphertext := make([]byte, len(data))
	cipher.XORKeyStream(ciphertext, data)
	return ciphertext
}

func decryptRC4(data, key, nonce []byte) []byte {
	cipher, _ := rc4.NewCipher(append(key, nonce...))
	plaintext := make([]byte, len(data))
	cipher.XORKeyStream(plaintext, data)
	return plaintext
}

// 3DES (Triple DES) NOT USE NONCE
func encrypt3DES(data, key []byte) ([]byte, error) {

	keyBuffer := key
	lenKeyBuffer := len(keyBuffer)
	if lenKeyBuffer < 24 {
		complementBuffer := make([]byte, int(24-lenKeyBuffer))
		keyBuffer = append(keyBuffer, complementBuffer...)
	} else {
		keyBuffer = keyBuffer[:24]
	}

	block, e := des.NewTripleDESCipher(keyBuffer)
	if e != nil {
		fmt.Println(e)
		return nil, e
	}

	ciphertext := make([]byte, des.BlockSize+len(data))
	iv := ciphertext[:des.BlockSize]
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[des.BlockSize:], data)
	return ciphertext, nil
}

func decrypt3DES(data, key []byte) ([]byte, error) {

	keyBuffer := key
	lenKeyBuffer := len(keyBuffer)
	if lenKeyBuffer < 24 {
		complementBuffer := make([]byte, int(24-lenKeyBuffer))
		keyBuffer = append(keyBuffer, complementBuffer...)
	} else {
		keyBuffer = keyBuffer[:24]
	}

	block, e := des.NewTripleDESCipher(keyBuffer)
	if e != nil {
		fmt.Println(e)
		return nil, e
	}

	ciphertext := make([]byte, des.BlockSize+len(data))
	iv := ciphertext[:des.BlockSize]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(ciphertext[des.BlockSize:], data)

	return ciphertext[des.BlockSize+des.BlockSize:], nil
}

// AES (Advanced Encryption Standard)
func encryptAES(data, key []byte, nonce []byte) ([]byte, error) {

	keyBuffer := key
	lenKeyBuffer := len(keyBuffer)
	if lenKeyBuffer < 32 {
		complementBuffer := make([]byte, int(32-lenKeyBuffer))
		keyBuffer = append(keyBuffer, complementBuffer...)
	} else {
		keyBuffer = keyBuffer[:32]
	}

	nonceBuffer := nonce
	lenNonceBuffer := len(nonceBuffer)
	if lenNonceBuffer < 12 {
		complementBuffer := make([]byte, int(12-lenNonceBuffer))
		nonceBuffer = append(nonceBuffer, complementBuffer...)
	} else {
		nonceBuffer = nonceBuffer[:12]
	}

	block, e := aes.NewCipher(keyBuffer)
	if e != nil {
		fmt.Println(e)
		return nil, e
	}

	gcm, e := cipher.NewGCM(block)
	if e != nil {
		fmt.Println(e)
		return nil, e
	}

	ciphertext := gcm.Seal(nil, nonceBuffer, data, nil)

	return ciphertext, nil
}

func decryptAES(ciphertext, key []byte, nonce []byte) ([]byte, error) {

	keyBuffer := key
	lenKeyBuffer := len(keyBuffer)
	if lenKeyBuffer < 32 {
		complementBuffer := make([]byte, int(32-lenKeyBuffer))
		keyBuffer = append(keyBuffer, complementBuffer...)
	} else {
		keyBuffer = keyBuffer[:32]
	}

	nonceBuffer := nonce
	lenNonceBuffer := len(nonceBuffer)
	if lenNonceBuffer < 12 {
		complementBuffer := make([]byte, int(12-lenNonceBuffer))
		nonceBuffer = append(nonceBuffer, complementBuffer...)
	} else {
		nonceBuffer = nonceBuffer[:12]
	}

	block, e := aes.NewCipher(keyBuffer)
	if e != nil {
		fmt.Println(e)
		return nil, e
	}
	gcm, e := cipher.NewGCM(block)
	if e != nil {
		fmt.Println(e)
		return nil, e
	}
	plaintext, e := gcm.Open(nil, nonceBuffer, ciphertext, nil)
	if e != nil {
		fmt.Println(e)
		return nil, e
	}

	return plaintext, nil
}
