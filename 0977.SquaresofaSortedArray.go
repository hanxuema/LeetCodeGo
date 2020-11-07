package main


func sortedSquares(A []int) []int {
    res := make([]int, len(A))
    if A == nil || len(A) == 0 {
        return res
    }
    neg := -1
    pos := 0 
	for i := 0; i < len(A); i++ {
		if A[i] >= 0 {
			pos = i
			break
		}
	}
	neg = pos - 1
    // fmt.Println("neg", neg)
    // fmt.Println("pos", pos)
    i := 0
    for neg >= 0 && pos < len(A) {
        if A[neg] * A[neg] < A[pos] * A[pos] {
            res[i] = A[neg] * A[neg]
            neg--
        } else{
            res[i] = A[pos] * A[pos]
            pos++
        }
        i++
    }
    for neg >= 0 {
        res[i] = A[neg] * A[neg]
        neg--
        i++
    }
    for pos < len(A){
         res[i] = A[pos] * A[pos]
        i++ 
        pos++
    }
    return res
}