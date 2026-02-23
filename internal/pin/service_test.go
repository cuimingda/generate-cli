package pin

import "testing"

func TestGenerateValidLength(t *testing.T) {
	s := NewService()

	pin, err := s.Generate(6)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(pin) != 6 {
		t.Fatalf("expected length 6, got %d", len(pin))
	}
}

func TestGenerateInvalidLength(t *testing.T) {
	s := NewService()

	_, err := s.Generate(5)
	if err == nil {
		t.Fatalf("expected error for invalid length")
	}
}
