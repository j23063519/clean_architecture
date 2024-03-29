package hash

import (
	"github.com/j23063519/clean_architecture/pkg/log"
	"golang.org/x/crypto/bcrypt"
)

func BcryptHash(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	log.ErrorJSON("Hash", "BcryptHash", err)

	return string(bytes)
}

func BcryptCheck(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}

func BcryptIsHashed(str string) bool {
	return len(str) == 60
}
