package passwords

import (
	"golang.org/x/crypto/bcrypt"
	"math/rand"

	"Nutrition/tools/logger"
)

const symbols = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func Encode(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	if err != nil {
		logger.Warning(err.Error())
	}
	return string(hash)
}

func Validate(password string, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

func Generate() string {
	b := make([]byte, 8)
	for i := range b {
		b[i] = symbols[rand.Int63()%int64(len(symbols))]
	}
	return string(b)
}
