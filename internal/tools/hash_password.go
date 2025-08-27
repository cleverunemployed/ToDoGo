package tools

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	// cost определяет сложность хеширования (рекомендуется 10-14)
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
