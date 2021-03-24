func sumOddLengthSubarrays(arr []int) int {
    result := 0
    offset := 1
 
    for ; offset <= len(arr); offset += 2 {
        for i := 0; i <= len(arr) - offset; i++ {
            result += sum(arr[i:offset+i])
        }
    }
    
    return result
}

func sum(arr []int) int {
    r := 0
    for _, v := range arr {
       r += v 
    }
    
    return r
}
