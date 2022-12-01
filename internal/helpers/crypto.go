package helpers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
	"log"
)

func Encrypt(key, stringToEncrypt string) string {
	if stringToEncrypt == "" {
		return stringToEncrypt
	}
	cipherBlock, err := aes.NewCipher([]byte(key))
	if err != nil {
		log.Fatalf("Encrypt - aes.NewCipher - %v", err)
	}

	aead, err := cipher.NewGCM(cipherBlock)
	if err != nil {
		log.Fatalf("Encrypt - cipher.NewGCM - %v", err)
	}

	nonce := make([]byte, aead.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		log.Fatalf("Encrypt - io.ReadFull(rand.Reader, nonce) - %v", err)
	}

	return base64.URLEncoding.EncodeToString(aead.Seal(nonce, nonce, []byte(stringToEncrypt), nil))
}

func Decrypt(key, encryptedString string) (decryptedString string) {
	if encryptedString == "" {
		return encryptedString
	}
	encryptData, err := base64.URLEncoding.DecodeString(encryptedString)
	if err != nil {
		log.Fatal(err)
	}

	cipherBlock, err := aes.NewCipher([]byte(key))
	if err != nil {
		log.Fatalf("Decrypt - aes.NewCipher - %v", err)
	}

	aead, err := cipher.NewGCM(cipherBlock)
	if err != nil {
		log.Fatalf("Decrypt - cipher.NewGCM - %v", err)
	}

	nonceSize := aead.NonceSize()
	if len(encryptData) < nonceSize {
		log.Fatalf("Decrypt - aead.NonceSize - %v", err)
	}

	nonce, cipherText := encryptData[:nonceSize], encryptData[nonceSize:]
	plainData, err := aead.Open(nil, nonce, cipherText, nil)
	if err != nil {
		log.Fatalf("Decrypt - aead.Open - %v", err)
	}

	return string(plainData)
}

func EncryptStream(key string, reader io.Reader) io.Reader {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		log.Fatalf("EncryptStream - NewCipher - %v", err)
	}
	var iv [aes.BlockSize]byte
	stream := cipher.NewOFB(block, iv[:])

	return &cipher.StreamReader{S: stream, R: reader}
}

func DecryptStream(key string, reader io.Writer) io.Writer {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		log.Fatalf("DecryptStream - NewCipher - %v", err)
	}
	var iv [aes.BlockSize]byte
	stream := cipher.NewOFB(block, iv[:])

	return &cipher.StreamWriter{S: stream, W: reader}
}
