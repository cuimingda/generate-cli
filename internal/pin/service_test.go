package pin

import (
	"crypto/rand"
	"errors"
	"io"
	"testing"
)

func TestGenerateValidLength(t *testing.T) {
	s := NewService()

	testCases := []int{4, 6, 8}
	for _, length := range testCases {
		pin, err := s.Generate(length)
		if err != nil {
			t.Fatalf("unexpected error for length %d: %v", length, err)
		}
		if len(pin) != length {
			t.Fatalf("expected length %d, got %d", length, len(pin))
		}
	}
}

func TestGenerateInvalidLength(t *testing.T) {
	s := NewService()

	testCases := []int{0, 3, 5, 7, 9}
	for _, length := range testCases {
		_, err := s.Generate(length)
		if !errors.Is(err, ErrInvalidLength) {
			t.Fatalf("expected ErrInvalidLength for length %d, got %v", length, err)
		}
	}
}

func TestGenerateRandomReadError(t *testing.T) {
	s := NewService()

	oldReader := rand.Reader
	rand.Reader = errorReader{}
	defer func() {
		rand.Reader = oldReader
	}()

	_, err := s.Generate(6)
	if !errors.Is(err, io.ErrUnexpectedEOF) {
		t.Fatalf("expected io.ErrUnexpectedEOF, got %v", err)
	}
}

type errorReader struct{}

func (errorReader) Read(_ []byte) (int, error) {
	return 0, io.ErrUnexpectedEOF
}
