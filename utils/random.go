package utils

import (
	"crypto/rand"
	"math/big"
)

func Random(min, max int) int64 {
	res, _ := rand.Int(rand.Reader, big.NewInt(int64(max-min)))
	return res.Int64() + int64(min)
}
