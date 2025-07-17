package utils

import (
	"crypto/rand"
	"math/big"
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
