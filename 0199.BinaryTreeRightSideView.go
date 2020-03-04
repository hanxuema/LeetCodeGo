package main

func rightSideView(root *TreeNode) []int {
    res := []int{}
    if root == nil {
        return res
    }
    
    var q []*TreeNode
    q = append(q, root)
    
    for len(q) > 0 {
        length := len(q)
        level := []int{}
        for i:=0; i<length; i++ {
            node := q[0]
            q = q[1:]
            level = append(level, node.Val)
            if node.Left != nil {
                q = append(q, node.Left)
            }
            if node.Right != nil{
                q = append(q, node.Right)
            }        
        }
        if len(level) >= 1{
            res = append(res, level[len(level) - 1])
        }
    }
 
    return res
}


//    1            <---  [1,3,5]
//  /   \
// 2     3         <---
//  \     
//   5            <---

// [1,2,3,null,5,null,null]
// [1,2,3,null,5,null,4]
// [1,2,3,null,5,null,null,7,null]
// []
// [1]
// [1,2,null]