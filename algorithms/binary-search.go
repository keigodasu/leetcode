func search(nums []int, target int) int {
    left, right := 0, len(nums) - 1
    
    if len(nums) == 0 {
        return -1
    }
    
    for right >= left  {
        mid := left + ((right - left) /2)
        if nums[mid] == target {
            return mid
        } else if target > nums[mid] {
            left = mid + 1
        } else if target < nums[mid] {
            right = mid - 1 
        }
    }
    
    return -1
}
