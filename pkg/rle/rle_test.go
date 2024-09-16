package rle

import (
	"github.com/kirkegaard/go-rle"
	"testing"
)

func TestEncode(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"aaaaaaabbbbccc", "7a4b3c"},
		{"abc", "1a1b1c"},
		{"", ""},
	}

	for _, test := range tests {
		encoded, err := Encode(test.input)
		if err != nil {
			t.Fatalf("expected no error, but got %v", err)
		}
		if encoded != test.expected {
			t.Errorf("expected %s, but got %s", test.expected, encoded)
		}
	}
}

func TestDecode(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"7a4b3c", "aaaaaaabbbbccc"},
		{"1a1b1c", "abc"},
		{"", ""},
	}

	for _, test := range tests {
		decoded, err := Decode(test.input)
		if err != nil {
			t.Fatalf("expected no error, but got %v", err)
		}
		if decoded != test.expected {
			t.Errorf("expected %s, but got %s", test.expected, decoded)
		}
	}
}
