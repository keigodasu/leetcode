
func longestValidParentheses(s string) int {
    stack := []int{-1}
    max := 0
    
    for i := 0; i < len(s); i++ {
        if s[i] == '(' {
            stack = append(stack, i)
        } else {
            stack = stack[:len(stack) - 1]
            if len(stack) == 0 {
                stack = append(stack, i)
            } else {
                if i - stack[len(stack) - 1] > max {
                    max = i - stack[len(stack) - 1]
                }
            }
        }
    }
    
    return max
}

