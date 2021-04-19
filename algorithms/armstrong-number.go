func isArmstrong(N int) bool {
    t := strconv.Itoa(N)
    sum := 0
    
    for _, v := range t {
        sum += int(math.Pow(float64(v - '0'), float64(len(t))))
    }
    
    return sum == N
}
