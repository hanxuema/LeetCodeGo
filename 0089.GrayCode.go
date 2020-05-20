package main

func grayCode(n int) []int {
	res := []int{}
	//res = append(res, intLength(30))
	cur := 0
	backtrack89(n, cur, 0, &res)
	return res
}

func backtrack89(n int, cur int, start int, res *[]int) {

	if intLength(cur) == n {
		*res = append(*res, cur)
		return
	}
	binary := []int{0, 1}

	for i := start; i < len(binary); i++ {
		cur = cur*10 + binary[i]
		//fmt.Printf("cur %+v\n", cur)
		backtrack89(n, cur, i+1, res)

		cur = (cur - binary[i]) / 10
	}
}

func power(n int) int {
	res := 1
	for i := 0; i <= n; i++ {
		res *= 2
	}
	return res
}

func intLength(n int) int {
	res := 0
	for n > 0 {
		n = n / 10
		res++
	}

	return res
}
