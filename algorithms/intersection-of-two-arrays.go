func intersection(nums1 []int, nums2 []int) []int {
    var result []int
    hashmap :=  map[int]bool{}
    
    for _, v := range nums1 {
        hashmap[v] = false
    }
   
    for _, v := range nums2 {
        if flag, ok := hashmap[v]; ok{
            if !flag {
                result = append(result, v)
                hashmap[v] = true
            }
        }
    }
    
    return result
}
