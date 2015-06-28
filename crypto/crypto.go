package crypto

/*
	- Will want to move to OpenSSL in future
		- http://sosedoff.com/2015/05/22/data-encryption-in-go-using-openssl.html
	- Error handling
		- IV must be block size
		- Key must be block size
		- Data size must be a multiple of block size
	- Need to be able to handle different ciphers
	- Need to be able to handle different hashes
*/

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
)

func GetHash(bin []byte) []byte {

	// Create hash and feed data into ongoing hash
	hash := sha256.New()
	hash.Write(bin)

	// Return the hash and append to nil string
	return hash.Sum(nil)
}

func Encrypt(plaintext []byte, key []byte, iv []byte) (ciphertext []byte) {

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	ciphertext = make([]byte, len(plaintext))

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, plaintext)

	return
}

func Decrypt(ciphertext []byte, key []byte, iv []byte) (plaintext []byte) {

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	plaintext = make([]byte, len(ciphertext))

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(plaintext, ciphertext)

	return
}
