func arraysIntersection(arr1 []int, arr2 []int, arr3 []int) []int {
    var comOneAndTwo []int
    for _, v := range arr1 {
        for _, vv := range arr2 {
            if v == vv {
                comOneAndTwo = append(comOneAndTwo, v)
            }
        }
    }
    
    var comThreeAndRest []int
    for _, v := range comOneAndTwo {
        for _, vv := range arr3 {
            if v == vv { 
                comThreeAndRest = append(comThreeAndRest, v) 
            }
        }
    }
    
    return comThreeAndRest
}
