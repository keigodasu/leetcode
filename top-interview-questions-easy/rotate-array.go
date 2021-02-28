func rotate(nums []int, k int)  {
    if len(nums) < k {
        k = k % len(nums)
    }
    var result []int
    result = append(result, nums[len(nums)-k:]...)
    result = append(result, nums[:len(nums)-k]...)
    for i:=0; i < len(nums); i++ {
        nums[i] = result[i]
    }
}
