func fizzBuzz(n int) []string  {
	list := []string{}

	for i:=1; i<=n; i++ {
		switch  {
		case i % 3 == 0 && i % 5 != 0 :
			list = append(list, "Fizz")
		case i % 3 != 0 && i % 5 == 0:
			list = append(list, "Buzz")
		case i % 15 == 0 :
			list = append(list, "FizzBuzz")
		default:
			list = append(list, strconv.Itoa(i))
		}
	}

	return list
}


