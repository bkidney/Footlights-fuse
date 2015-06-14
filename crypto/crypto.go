package crypto

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
