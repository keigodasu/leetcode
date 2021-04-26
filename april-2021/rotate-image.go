func rotate(matrix [][]int)  {
    var newMatrix [][]int
    
    for i := 0; i < len(matrix); i++ {
        var tempList []int
        for j := len(matrix) - 1; j >= 0; j-- {
            tempList = append(tempList, matrix[j][i])
        }
        newMatrix = append(newMatrix, tempList)
        fmt.Println(newMatrix)
    }

    for i, v := range newMatrix {
        matrix[i] = v
    }
}
