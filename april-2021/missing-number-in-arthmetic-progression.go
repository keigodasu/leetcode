func missingNumber(arr []int) int {
    diff := (arr[len(arr) - 1] - arr[0]) / len(arr)
    
    expectedValue := arr[0]
    for _, v := range arr {
        if v != expectedValue {
            return expectedValue
        }else {
            expectedValue += diff
        }
    }
    return expectedValue
}
