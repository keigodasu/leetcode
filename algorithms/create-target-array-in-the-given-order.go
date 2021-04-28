func createTargetArray(nums []int, index []int) []int {
    result := make([]int, len(nums))
    for i, _ := range result {
        result[i] = -1
    }
    
    for i, v := range index {
        if result[v] != -1 {
            copy(result[v+1:], result[v:])
        }
        result[v] = nums[i]
    }
    
    return result
}
