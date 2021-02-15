func smallerNumbersThanCurrent(nums []int) []int {
    list := make([]int, len(nums))
    
   
    for i := 0; i < len(nums); i++ {
        count := 0
        for j := 0; j < len(nums); j++ {
            if nums[i] > nums[j] {
               count++ 
            }
        }
        list[i] = count
    }
    
    return list
}
