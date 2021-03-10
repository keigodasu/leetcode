func intersect(nums1 []int, nums2 []int) []int {
    
    if len(nums1) > len(nums2) {
        temp := nums1
        nums1 = nums2
        nums2 = temp
    }
    
    var result []int
    dic := map[int]int{}
    
    for _, v := range nums1 {
        dic[v]++
    }
    
    for _, v := range nums2 {
        if vv, ok := dic[v]; ok {
            if vv > 0 {
                result = append(result, v)
                dic[v]--
            }
        }
    }
    
    
    return result
}
