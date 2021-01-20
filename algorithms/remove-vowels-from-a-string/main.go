func removeVowels(s string) string {
    re := regexp.MustCompile(`[aiueo]`)
    return re.ReplaceAllString(s, "") 
}
