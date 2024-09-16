package rle

import (
	"github.com/stretchr/testify/assert"
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
		assert.Equal(t, test.expected, encoded, "Expected encoding to be %s, but got %s", test.expected, encoded)
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
		assert.Equal(t, test.expected, decoded, "Expected decoding to be %s, but got %s", test.expected, decoded)
	}
}
