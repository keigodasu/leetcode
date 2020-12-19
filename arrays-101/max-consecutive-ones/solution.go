func findMaxConsecutiveOnes(nums []int) int {   
    count := 0
    maxCount := 0
    
    for i := 0; i < len(nums); i++ {
        if(nums[i] == 1) {
            count++
        } else {
            maxCount = int(math.Max(float64(maxCount), float64(count)))
            count = 0
        }
    }
    
    return int(math.Max(float64(maxCount), float64(count)))
}
