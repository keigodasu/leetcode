package week_4_march_22nd_march_28th

func countSubstrings(s string) int {
	var result int

	for skip := 1; skip <= len(s); skip++ {
		for i := 0; i <= len(s) - skip; i++ {
			str := s[i:skip+i]
			if isPalindrome(str) {
				result++
			}
		}
	}
	return result
}

func isPalindrome(str string) bool {
	if len(str) == 1 {
		return true
	}

	for i := 0; i < len(str); i++ {
		j := len(str) - 1 - i
		if str[i] != str[j] {
			return false
		}
	}
	return true
}
