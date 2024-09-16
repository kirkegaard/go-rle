package rle

import (
	"strconv"
)

func Encode(input string) string {
	if len(input) == 0 {
		return ""
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

	return result
}

func Decode(input string) string {
	if len(input) == 0 {
		return ""
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

	return result
}
