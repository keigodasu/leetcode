func shuffle(nums []int, n int) []int {
    var list []int
    
    firstHalf := nums[:n]
    secondHalf := nums[n:]
   
    for i := 0; i < n; i++ {
        list = append(list, firstHalf[i])
        list = append(list, secondHalf[i])
    }
    
    return list
}
