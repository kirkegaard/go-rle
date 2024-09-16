package rle

import (
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
		encoded := Encode(test.input)
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
		decoded := Decode(test.input)
		if decoded != test.expected {
			t.Errorf("expected %s, but got %s", test.expected, decoded)
		}
	}
}
