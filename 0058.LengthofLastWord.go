package main

func lengthOfLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}
	chars := []byte(s)
	first := -1
	for i := len(chars) - 1; i >= 0; i-- {
		if first == -1 {
			if chars[i] == ' ' {
				continue
			} else {
				first = i
			}
		} else {
			if chars[i] == ' ' {
				return first - i
			}
		}
	}

	return first + 1
}

// "Hello World"
// " Hello World "
// "  "
// "a"
// " "
