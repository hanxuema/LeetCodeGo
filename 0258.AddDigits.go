package main

import "fmt"

func addDigits(num int) int {
	for lenOfNum(num) > 1 {
		num = sumOfDigits(num)
	}
	return num
}

func sumOfDigits(num int) int {
	res := 0
	for num > 0 {
		res = res + num%10
		fmt.Printf("sumOfDigits %+v\n", res)

		num = num / 10
		fmt.Printf("sumOfDigits num %+v\n", num)
	}
	fmt.Printf("sumOfDigits return res %+v\n", res)
	return res
}

func lenOfNum(num int) int {
	res := 0
	if num >= 0 {
		num = num / 10
		res = res + 1
	}
	fmt.Printf("lon %+v\n", res)
	return res
}
