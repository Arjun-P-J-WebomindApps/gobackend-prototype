package utils

import (
	"crypto/rand"
	"encoding/hex"
	"math/big"

	"golang.org/x/crypto/bcrypt"
)

func GenerateNumericOTP(length int) (int, error) {
	min := big.NewInt(1)

	max := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(length)), nil)
	diff := new(big.Int).Sub(max, min)

	n, err := rand.Int(rand.Reader, diff)

	if err != nil {
		return 0, err
	}

	n.Add(n, min)

	return int(n.Int64()), nil
}

func GenerateSecureToken(nBytes int) (string, error) {
	b := make([]byte, nBytes)

	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(b), nil
}

func HashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(hashedBytes), err
}

func CheckPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
