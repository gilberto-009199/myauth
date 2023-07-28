// You can edit this code!
// Click here and start typing.
package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"crypto/rc4"
	"fmt"

	"github.com/aead/camellia"
	"golang.org/x/crypto/chacha20"
)

func main() {

	data := []byte("otpauth://totp/Amazon%20Web%20Services:teste4534@gilbertoramos?secret=GQMWATC3JYIGQBZRDPT7JEVWVDZQISEOXZXUP42552H6VRJTIXM3YLHI37J3YW6N&issuer=Amazon%20Web%20Services")
	key := []byte("123456789")
	nonce := []byte("gil")

	/* AES */
	ciphertext, _ := encryptAES(data, key, nonce)
	plaintext, _ := decryptAES(ciphertext, key, nonce)
	fmt.Println("## AES ##")
	fmt.Println("\tOrigin:\n\t", string(data))
	fmt.Println("\tCipher:\n", string(ciphertext))
	fmt.Println("\tPlain:\n\t", string(plaintext))
	/**/

	/* 3DES */
	ciphertext, _ = encrypt3DES(data, key)
	plaintext, _ = decrypt3DES(ciphertext, key)
	fmt.Println("## 3DES ##")
	fmt.Println("\tOrigin:\n\t", string(data))
	fmt.Println("\tCipher:\n", string(ciphertext))
	fmt.Println("\tPlain:\n\t", string(plaintext))
	/**/

	/* RC4 */
	ciphertext = encryptRC4(data, key, nonce)
	plaintext = decryptRC4(ciphertext, key, nonce)
	fmt.Println("## RC4 ##")
	fmt.Println("\tOrigin:\n\t", string(data))
	fmt.Println("\tCipher:\n", string(ciphertext))
	fmt.Println("\tPlain:\n\t", string(plaintext))
	/**/

	/* chacha20 */
	ciphertext = encryptChaCha20(data, key, nonce)
	plaintext = decryptChaCha20(ciphertext, key, nonce)
	fmt.Println("## chacha20 ##")
	fmt.Println("\tOrigin:\n\t", string(data))
	fmt.Println("\tCipher:\n", string(ciphertext))
	fmt.Println("\tPlain:\n\t", string(plaintext))
	/**/

	/* Camellia */
	ciphertext, _ = encryptCamellia(data, key)
	plaintext, _ = decryptCamellia(ciphertext, key)
	fmt.Println("## Camellia ##")
	fmt.Println("\tOrigin:\n\t", string(data))
	fmt.Println("\tCipher:\n", string(ciphertext))
	fmt.Println("\tPlain:\n\t", string(plaintext))
	/**/
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
