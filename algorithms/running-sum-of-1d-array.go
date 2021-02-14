func runningSum(nums []int) []int {
    list := make([]int, len(nums))
    accumlated := 0
    for i, v := range nums {
        list[i] = v + accumlated
        accumlated += v
    }
    return list
}
