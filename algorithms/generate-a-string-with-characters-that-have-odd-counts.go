func generateTheString(n int) string {
    base := strings.Repeat("a", n-1)
    if n % 2 == 0 {
       return base + "b" 
    } else {
       return base + "a"
    }
}
