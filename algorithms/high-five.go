func highFive(items [][]int) [][]int {
    var studentScores = make(map[int][]int)
    result := [][]int{}
    
    for _, score := range items {
        studentScores[score[0]] = append(studentScores[score[0]], score[1])
    }
    
    for k, v := range studentScores {
        result = append(result, []int{k, average(topK(v, 5))})
    }
    
    sort.Slice(result, func(i, j int) bool {
        return result[i][0] < result[j][0]
    })
    
    return result
}

func topK(nums []int, k int)  []int {
    sort.Ints(nums)
    return nums[len(nums)-k : len(nums)]
}

func average(nums []int) int {
    sum := 0
    for _, val := range nums {
        sum += val
    }
    
    return sum / len(nums)
}
