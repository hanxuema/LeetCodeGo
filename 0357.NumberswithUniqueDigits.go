package main

import "fmt"

func countNumbersWithUniqueDigits(n int) int {
	res := 0
	if n == 0 {
		return 1
	}
	cur := []int{}
	visited := make([]bool, 10)
	backtrack(n, cur, visited, &res)
	return res
}

func backtrack(n int, cur []int, visited []bool, res *int) {
	if len(cur) == n {
		fmt.Println(cur)
		(*res)++
		return
	}

	for i := 0; i < 10; i++ {
		if visited[i] {
			//fmt.Printf("removed %+v\n", cur)
			continue
		}
		cur = append(cur, i)

		visited[i] = true
		backtrack(n, cur, visited, res)
		visited[i] = false

		cur = cur[:len(cur)-1]
	}
}

// func checkNumberDigit(num int) int{
//     res := 0
//     for num > 0 {
//         num = num / 10
//             res++
//     }
//     return res
// }

// func power(num int)int{
//     res := 1
//     for i := 0; i < num; i++{
//          res = res * 10
//     }
//     return res
// }
