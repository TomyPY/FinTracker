package encrypt

type FakeEncrypter struct {
	MockHashPassword   func(password string) (string, error)
	MockVerifyPassword func(password, hash string) error
	MockEncryptToken   func(plainText string) (string, error)
	MockDecryptToken   func(cipherText string) (string, error)
}

func NewFakeEncrypter() Encrypter {
	return &FakeEncrypter{
		MockHashPassword: func(password string) (string, error) {
			return "", nil
		},
		MockVerifyPassword: func(password, hash string) error {
			if password != hash {
				return ErrInvalidPassword
			}
			return nil
		},
		MockEncryptToken: func(plainText string) (string, error) {
			return "", nil
		},
		MockDecryptToken: func(cipherText string) (string, error) {
			return "", nil
		},
	}
}

func (e FakeEncrypter) HashPassword(password string) (string, error) {
	return e.MockHashPassword(password)
}

func (e FakeEncrypter) VerifyPassword(password, hash string) error {
	return e.MockVerifyPassword(password, hash)
}

func (e FakeEncrypter) EncryptToken(plainText string) (string, error) {
	return e.MockEncryptToken(plainText)
}

func (e FakeEncrypter) DecryptToken(cipherText string) (string, error) {
	return e.MockDecryptToken(cipherText)
}
