package helper

import "golang.org/x/crypto/bcrypt"

func PasswordHash(password string) string {
	bytePassword := []byte(password)
	hashPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)

	PanicIfError(err)

	return string(hashPassword)
}
