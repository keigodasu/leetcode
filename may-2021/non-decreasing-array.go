func checkPossibility(nums []int) bool {
    count := 0
    
    for i := 1; i < len(nums); i++ {
        if nums[i-1] > nums[i] {
            if count == 1 {
                return false
            }
            
            count++
            
            if i < 2 || nums[i-2] <= nums[i] {
                nums[i-1] = nums[i]
            } else {
                nums[i] = nums[i-1]
            }
        }
    }
    return true
}
