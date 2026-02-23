package pin

import (
	"crypto/rand"
	"errors"
	"math/big"
)

const digits = "0123456789"

var ErrInvalidLength = errors.New("invalid length, supported values: 4/6/8")

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Generate(length int) (string, error) {
	if !isValidLength(length) {
		return "", ErrInvalidLength
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

func isValidLength(length int) bool {
	switch length {
	case 4, 6, 8:
		return true
	default:
		return false
	}
}
