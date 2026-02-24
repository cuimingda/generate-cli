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
		pins, err := s.Generate(length)
		if err != nil {
			t.Fatalf("unexpected error for length %d: %v", length, err)
		}
		if len(pins) != 10 {
			t.Fatalf("expected default count 10, got %d", len(pins))
		}
		for _, generated := range pins {
			if len(generated) != length {
				t.Fatalf("expected length %d, got %d", length, len(generated))
			}
		}
	}
}

func TestGenerateCustomCount(t *testing.T) {
	s := NewService()

	pins, err := s.Generate(6, 3)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(pins) != 3 {
		t.Fatalf("expected 3 pins, got %d", len(pins))
	}
	for _, generated := range pins {
		if len(generated) != 6 {
			t.Fatalf("expected length 6, got %d", len(generated))
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

	_, err := s.Generate(6, 1)
	if !errors.Is(err, io.ErrUnexpectedEOF) {
		t.Fatalf("expected io.ErrUnexpectedEOF, got %v", err)
	}
}

type errorReader struct{}

func (errorReader) Read(_ []byte) (int, error) {
	return 0, io.ErrUnexpectedEOF
}
