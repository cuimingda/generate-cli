package slug

import (
	"crypto/rand"
	"errors"
	"math/big"
)

const charset = "abcdefghijklmnopqrstuvwxyz0123456789"

const (
	MinLength    = 8
	MaxLength    = 16
	DefaultLen   = 12
	MinCount     = 1
	MaxCount     = 10
	DefaultCount = 1
)

var (
	ErrInvalidLength = errors.New("invalid length, supported range: 8-16")
	ErrInvalidCount  = errors.New("invalid count, supported range: 1-10")
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Generate(length int, count int) ([]string, error) {
	if length < MinLength || length > MaxLength {
		return nil, ErrInvalidLength
	}

	if count < MinCount || count > MaxCount {
		return nil, ErrInvalidCount
	}

	results := make([]string, 0, count)
	max := big.NewInt(int64(len(charset)))

	for range count {
		bytes := make([]byte, length)
		for i := range bytes {
			n, err := rand.Int(rand.Reader, max)
			if err != nil {
				return nil, err
			}
			bytes[i] = charset[n.Int64()]
		}
		results = append(results, string(bytes))
	}

	return results, nil
}
