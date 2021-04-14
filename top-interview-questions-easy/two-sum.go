func twoSum(nums []int, target int) []int {
    dic := map[int]int{}
    
    for i, v := range nums {
        diff := target - v
        if idx, ok := dic[diff]; ok {
            return []int{i, idx}
        }
        dic[v] = i
    }
    
    return nil
}
