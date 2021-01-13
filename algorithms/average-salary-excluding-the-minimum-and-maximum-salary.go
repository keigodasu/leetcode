func average(salary []int) float64 {
    max := findMaxSalary(salary)
    min := findMinSalary(salary)
  
    var total int
    for _, s := range salary {
        if s != max && s != min {
            total += s
        }
    }

    return float64(total)/float64((len(salary)-2))
}

func findMaxSalary(salary []int) int{
    max := 0
    for _, s := range salary {
        if s >= max {
            max = s
        }
    }
    
    return max
}

func findMinSalary(salary []int) int{
    min := salary[0]
    for _, s := range salary {
        if s <= min {
            min = s
        }
    }
    
    return min
}

///////

func average(salary []int) float64 {
    total, min, max := 0, 100000000, 1000
    for _, v := range salary {
        total += v
        if v < min { min = v }
        if v > max { max = v }
    }

    return float64(total - max -min) / float64(len(salary) - 2)
}
