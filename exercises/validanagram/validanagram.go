package validanagram

func isAnagram(s string, t string) bool {

	if (len(s) != len(t)) {
		return false
	}

	mapperS := make(map[rune]int)
	mapperT := make(map[rune]int)

	for _, char := range s {
		mapperS[char]++
	}

	for _, char := range t {
		mapperT[char]++
	}

	for k, v := range mapperS {

		if mapperT[k] != v {
			return false
		}
	}

	return true
} 