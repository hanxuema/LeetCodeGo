package main

func findMode(root *TreeNode) []int {
    res := []int{}
    dic := make(map[int]int)
    inorder501(root,   &dic)
    max := 0
    for _,v := range  dic{
        if v > max{
            max = v
        }
    }
    for k, v := range dic {
        if v == max{
            res = append(res, k)
        }
    }
    return res
}

func inorder501(root *TreeNode,   dic *map[int]int) {
    if root == nil {
        return 
    }
    inorder501(root.Left, dic)
    (*dic)[root.Val]++
    inorder501(root.Right, dic)
}