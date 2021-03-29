package reconstruct_original_digits_from_english

import (
	"sort"
	"strconv"
)

func originalDigits(s string) string {
	result := []int{}
	dic := map[string]int{}

	for _, v := range s {
		dic[string(v)]++
	}

	if _, num := dic["z"]; num {
		for dic["z"] > 0 {
			result = append(result, 0)
			dic["z"]--
			dic["e"]--
			dic["r"]--
			dic["o"]--
		}
	}

	if _, num := dic["w"]; num {
		for dic["w"] > 0 {
			result = append(result, 2)
			dic["t"]--
			dic["w"]--
			dic["o"]--
		}
	}

	if _, num := dic["u"]; num {
		for dic["u"] > 0 {
			result = append(result, 4)
			dic["f"]--
			dic["o"]--
			dic["u"]--
			dic["r"]--
		}
	}

	if _, num := dic["x"]; num {
		for dic["x"] > 0 {
			result = append(result, 6)
			dic["s"]--
			dic["i"]--
			dic["x"]--
		}
	}

	if _, num := dic["g"]; num {
		for dic["g"] > 0 {
			result = append(result, 8)
			dic["e"]--
			dic["i"]--
			dic["g"]--
			dic["h"]--
			dic["t"]--
		}
	}

	if _, num := dic["r"]; num {
		for dic["r"] > 0 {
			result = append(result, 3)
			dic["t"]--
			dic["h"]--
			dic["r"]--
			dic["e"]--
			dic["e"]--
		}
	}

	if _, num := dic["s"]; num {
		for dic["s"] > 0 {
			result = append(result, 7)
			dic["s"]--
			dic["e"]--
			dic["v"]--
			dic["e"]--
			dic["n"]--
		}
	}

	if _, num := dic["v"]; num {
		for dic["v"] > 0 {
			result = append(result, 5)
			dic["f"]--
			dic["i"]--
			dic["v"]--
			dic["e"]--
		}
	}

	if _, num := dic["i"]; num {
		for dic["i"] > 0 {
			result = append(result, 9)
			dic["n"]--
			dic["i"]--
			dic["n"]--
			dic["e"]--
		}
	}

	if _, num := dic["o"]; num {
		for dic["o"] > 0 {
			result = append(result, 1)
			dic["o"]--
			dic["n"]--
			dic["e"]--
		}
	}

	sort.Ints(result)

	resultstring := ""
	for _, v := range result {
		resultstring += strconv.Itoa(v)
	}

	return resultstring
}
