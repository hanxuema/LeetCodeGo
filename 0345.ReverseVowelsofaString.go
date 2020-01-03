package main

func reverseVowels(s string) string {
	sByte := []byte(s)
	if sByte == nil || len(sByte) <= 1 {
		return s
	}
	set := make(map[byte]bool)
	set['a'] = true
	set['e'] = true
	set['i'] = true
	set['o'] = true
	set['u'] = true
	set['A'] = true
	set['E'] = true
	set['I'] = true
	set['O'] = true
	set['U'] = true

	left, right := 0, len(s)-1
	for left < right {
		_, leftok := set[sByte[left]]
		if !leftok {
			left++
		}
		_, rightok := set[sByte[right]]
		if !rightok {
			right--
		}
		if leftok && rightok {
			temp := sByte[left]
			sByte[left] = sByte[right]
			sByte[right] = temp
			left++
			right--
		}
	}

	return string(sByte)
}
