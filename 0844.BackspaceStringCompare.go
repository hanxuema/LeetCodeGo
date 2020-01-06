package main

func backspaceCompare(S string, T string) bool {
    s, t := []byte(S), []byte(T)
    first, second := len(s) -1, len(t) -1
    firstCount, secondCount := 0,0
    for first >= 0 || second >= 0 {
        for first >= 0 {
            if s[first] == '#'{
                first--
                firstCount++
            } else if firstCount > 0 {
                first--
                firstCount--
            } else {
                break
            }
        }
        for second >= 0  {
            if t[second] == '#'{
                second--
                secondCount++
            } else if secondCount > 0 {
                second--
                secondCount--
            } else {
                break
            }
        }
        if (first < 0) != (second < 0) {
            return false
        }
        if first >= 0 && second >= 0 && s[first] != t[second]{
            return false
		}  
		first--
		second--
    }
    return true
}