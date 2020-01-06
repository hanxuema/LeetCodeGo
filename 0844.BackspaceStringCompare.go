package main

func backspaceCompare(S string, T string) bool {
	s := []byte(S)
	t := []byte(T)
	first, second := len(s)-1, len(t)-1
	countFirst, countSecond := 0, 0

	for first >= 0 || second >= 0 {
		for s[first] == '#' && first > 0 {
			countFirst++
			first--
		}
		for t[second] == '#' && second > 0 {
			countSecond++
			second--
		}
		for countFirst > 0 && first > 0 {
			first--
			countFirst--
		}
		for countSecond > 0 && second > 0 {
			second--
			countSecond--
		}
		if s[first] != t[second] {
			return false
		} else {
			first--
			second--
		}
	}
	if first > 0 || second > 0 {
		return false
	}
	return true
}
