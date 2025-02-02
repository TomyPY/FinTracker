package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidPassword = errors.New("invalid password")
)

type Encrypter interface {
	HashPassword(password string) (string, error)
	VerifyPassword(password, hash string) error
	EncryptToken(plainText string) (string, error)
	DecryptToken(encryptedText string) (string, error)
}

func NewEncrypter(key string) Encrypter {
	return encrypter{
		encryptionKey: []byte(key),
	}
}

type encrypter struct {
	encryptionKey []byte
}

func (e encrypter) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// VerifyPassword verifies if the given password matches the stored hash.
func (e encrypter) VerifyPassword(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return ErrInvalidPassword
		}
		return err
	}
	return nil
}

func (e encrypter) EncryptToken(plainText string) (string, error) {
	block, err := aes.NewCipher(e.encryptionKey)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, 12) // GCM estándar
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	cipherText := aesGCM.Seal(nil, nonce, []byte(plainText), nil)
	encrypted := append(nonce, cipherText...)

	return base64.URLEncoding.EncodeToString(encrypted), nil
}

func (e encrypter) DecryptToken(encryptedText string) (string, error) {
	data, err := base64.URLEncoding.DecodeString(encryptedText)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(e.encryptionKey)
	if err != nil {
		return "", err
	}

	nonceSize := 12
	if len(data) < nonceSize {
		return "", errors.New("data too short")
	}

	nonce, cipherText := data[:nonceSize], data[nonceSize:]
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	plainText, err := aesGCM.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return "", err
	}

	return string(plainText), nil
}
