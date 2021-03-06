func findErrorNums(nums []int) []int {
    dup := 0
    mis := 0
    dic := map[int]int{}
    
    for _, v := range nums {
        dic[v]++
    }
    
    for i := 1; i <= len(nums); i++ {
        if v, ok := dic[i]; ok {
            if v == 2 { 
                dup = i
            }
        } else {
            mis = i
        }
    }
    
    return []int{dup, mis}
}
