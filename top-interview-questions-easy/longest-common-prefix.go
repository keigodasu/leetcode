func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }
    
    var sb strings.Builder
    for i := 0; i < len(strs[0]);i++ {
        for _, str := range strs {
            if i == len(str) || strs[0][i] != str[i] {
                return sb.String()
            }
        }
        sb.WriteByte(strs[0][i])
    }
   
    return sb.String()
}
