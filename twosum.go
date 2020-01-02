package main

 
// func main()  {
// 	nums := []int{2,7,11,15}
// 	target := 9
// 	twoSum(nums, target)
// }

func twoSum(nums []int, target int) []int {
	res := []int{0,0}
    if nums == nil || len(nums) < 2{ 
        return res
    }
    dic := make(map[int]int)
    for i, num := range nums {
         if _, ok := dic[target - num]; ok {
			 res[0] = i
			 res[1] = dic[target - num]
            return res
         }else{
             dic[num] = i
         }         
    }
    return res
}