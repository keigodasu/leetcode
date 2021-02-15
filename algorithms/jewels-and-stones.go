func numJewelsInStones(jewels string, stones string) int {
    count := 0
    for _, v := range jewels {
        count += strings.Count(stones, string(v))
    }
   
    return count
}
