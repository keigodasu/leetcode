func jump(nums []int) int {
    jumps := 0
    currentJumEnd := 0
    farthest := 0
    
    for i := 0; i < len(nums) - 1; i++ {
        farthest = int(math.Max(float64(farthest), float64(i + nums[i])))
        if i == currentJumEnd {
            jumps++
            currentJumEnd = farthest
        }
    }
    
    return jumps
}
