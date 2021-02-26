func reverseWords(s string) string {
    array := strings.Split(s, " ")
    f := ""
    for _, v := range array  {
        rev := ""
        for i := len(v) - 1; i >= 0; i-- {
            rev += string(v[i])
        }
        f += rev + " "
    }

    return strings.TrimSpace(f)
}
