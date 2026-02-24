package slug

import (
	"crypto/rand"
	"errors"
	"io"
	"strings"
	"testing"
)

func TestGenerateValidInput(t *testing.T) {
	s := NewService()

	slugs, err := s.Generate(12, 3)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(slugs) != 3 {
		t.Fatalf("expected 3 slugs, got %d", len(slugs))
	}

	for _, generated := range slugs {
		if len(generated) != 12 {
			t.Fatalf("expected length 12, got %d", len(generated))
		}

		for _, ch := range generated {
			if !strings.ContainsRune(charset, ch) {
				t.Fatalf("generated slug contains invalid character %q", ch)
			}
		}
	}
}

func TestGenerateInvalidLength(t *testing.T) {
	s := NewService()

	testCases := []int{0, 7, 17}
	for _, length := range testCases {
		_, err := s.Generate(length, 1)
		if !errors.Is(err, ErrInvalidLength) {
			t.Fatalf("expected ErrInvalidLength for length %d, got %v", length, err)
		}
	}
}

func TestGenerateInvalidCount(t *testing.T) {
	s := NewService()

	testCases := []int{0, 11}
	for _, count := range testCases {
		_, err := s.Generate(12, count)
		if !errors.Is(err, ErrInvalidCount) {
			t.Fatalf("expected ErrInvalidCount for count %d, got %v", count, err)
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

	_, err := s.Generate(12, 1)
	if !errors.Is(err, io.ErrUnexpectedEOF) {
		t.Fatalf("expected io.ErrUnexpectedEOF, got %v", err)
	}
}

type errorReader struct{}

func (errorReader) Read(_ []byte) (int, error) {
	return 0, io.ErrUnexpectedEOF
}
