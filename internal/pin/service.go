package pin

import (
	"crypto/rand"
	"errors"
	"math/big"
)

const digits = "0123456789"

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Generate(length int) (string, error) {
	if length != 4 && length != 6 && length != 8 {
		return "", errors.New("invalid length, supported values: 4/6/8")
	}

	bytes := make([]byte, length)
	max := big.NewInt(int64(len(digits)))

	for i := range bytes {
		n, err := rand.Int(rand.Reader, max)
		if err != nil {
			return "", err
		}
		bytes[i] = digits[n.Int64()]
	}

	return string(bytes), nil
}
