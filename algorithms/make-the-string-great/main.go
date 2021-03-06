package make_the_string_great

import (
	"strings"
	"unicode"
)

func makeGoodUsingStack(s string) string  {
	stack := []rune{}
	for _, r := range s {
		if len(stack) > 0 && (stack[len(stack)-1] == r+32 || stack[len(stack)-1] == r-32) {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, r)
		}
	}
	return string(stack)
}

func makeGood(s string) string {
	newString := ""
	duplicatedFlag := false
	for i := 0; i <= len(s) -1; i++ {
		if i == len(s)-1 {
			newString += s[i:i+1]
			break
		}
		if unicode.IsUpper(rune(s[i])) {
			if strings.ToLower(s[i:i+1]) == s[i+1:i+2] {
				duplicatedFlag = true
				i++
				continue
			} else {
				newString += s[i:i+1]
			}
		} else {
			if strings.ToUpper(s[i:i+1]) == s[i+1:i+2] {
				duplicatedFlag = true
				i++
				continue
			} else {
				newString += s[i:i+1]
			}
		}
	}

	if !duplicatedFlag {
		return s
	}

	return  makeGood(newString)
}