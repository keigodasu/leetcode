func flipAndInvertImage(A [][]int) [][]int {
    output := make([][]int, len(A))
    
    for i := 0; i < len(A); i++ {
        var temp []int
        for j := len(A[i]) - 1; j >= 0; j-- {
            temp = append(temp, 1^A[i][j])
        }
        output[i] = temp
    }
   
    return output
}
