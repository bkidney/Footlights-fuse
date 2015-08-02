package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"hash"

	"github.com/spf13/viper"
)

func Encrypt(plaintext []byte, key []byte, iv []byte) (ciphertext []byte) {

	switch viper.GetString("cipher") {
	case "AES":
		return EncryptAES(plaintext, key, iv)
	}

	return
}

func Decrypt(ciphertext []byte, key []byte, iv []byte) (plaintext []byte) {

	switch viper.GetString("cipher") {
	case "AES":
		return DecryptAES(ciphertext, key, iv)
	}

	return
}

func Hash(in []byte) (out []byte) {

	var fn hash.Hash

	switch viper.GetString("hash") {
	case "SHA256":
		fn = sha256.New()
	}

	fn.Write(in)

	// Return the hash and append to nil string
	return fn.Sum(nil)
}
func EncryptAES(plaintext []byte, key []byte, iv []byte) (ciphertext []byte) {

	cipher, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	switch viper.GetString("cipher-mode") {
	case "CBC":
		return EncryptAESCBC(cipher, plaintext, iv)
	}

	return
}

func DecryptAES(ciphertext []byte, key []byte, iv []byte) (plaintext []byte) {

	cipher, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	switch viper.GetString("cipher-mode") {
	case "CBC":
		return DecryptAESCBC(cipher, ciphertext, iv)
	}

	return
}

func EncryptAESCBC(aes cipher.Block, plaintext []byte, iv []byte) (ciphertext []byte) {

	enc := cipher.NewCBCEncrypter(aes, iv)

	ciphertext = make([]byte, len(plaintext))
	enc.CryptBlocks(ciphertext, plaintext)
	return
}

func DecryptAESCBC(aes cipher.Block, ciphertext []byte, iv []byte) (plaintext []byte) {

	dec := cipher.NewCBCDecrypter(aes, iv)

	plaintext = make([]byte, len(ciphertext))
	dec.CryptBlocks(plaintext, ciphertext)
	return
}
