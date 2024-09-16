package rle

import (
	"errors"
	"strconv"
)

func rleEncode(input string) (string, error) {
	if len(input) == 0 {
		return "", errors.New("Input is empty")
	}

	result := ""
	count := 1

	for i := 1; i < len(input); i++ {
		if input[i] == input[i-1] {
			count++
		} else {
			result += strconv.Itoa(count) + string(input[i-1])
			count = 1
		}
	}

	result += strconv.Itoa(count) + string(input[len(input)-1])

	return result, nil
}

func rleDecode(input string) (string, error) {
	if len(input) == 0 {
		return "", errors.New("Input is empty")
	}

	result := ""
	count := 0

	for i := 0; i < len(input); i++ {
		if input[i] >= '0' && input[i] <= '9' {
			count = count*10 + int(input[i]-'0')
		} else {
			for j := 0; j < count; j++ {
				result += string(input[i])
			}
			count = 0
		}
	}

	return result, nil
}
