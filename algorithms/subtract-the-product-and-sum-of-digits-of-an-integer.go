func subtractProductAndSum(n int) int {
    sum := 0
    product := 1
    
    for _, v := range strings.Split(fmt.Sprintf("%d", n), "") {
        num, _ := strconv.Atoi(v)
        sum += num
        product *= num
    }
  
    return product - sum
    
}
