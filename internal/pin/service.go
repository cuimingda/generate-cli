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

func (s *Service) Generate(length int, count ...int) ([]string, error) {
	if !isValidLength(length) {
		return nil, ErrInvalidLength
	}

	total := 10
	if len(count) > 0 && count[0] > 0 {
		total = count[0]
	}

	results := make([]string, 0, total)
	max := big.NewInt(int64(len(digits)))

	for range total {
		bytes := make([]byte, length)
		for i := range bytes {
			n, err := rand.Int(rand.Reader, max)
			if err != nil {
				return nil, err
			}
			bytes[i] = digits[n.Int64()]
		}
		results = append(results, string(bytes))
	}

	return results, nil
}

func isValidLength(length int) bool {
	switch length {
	case 4, 6, 8:
		return true
	default:
		return false
	}
}
