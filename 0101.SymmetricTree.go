package main

func isSymmetricIterative(root *TreeNode) bool {
    if root == nil {
        return true
    }
    if root.Left == nil && root.Right == nil {
        return true
    }
    if root.Left == nil || root.Right == nil{
        return false
    }
    if root.Left.Val != root.Right.Val {
        return false
    }
    var q []*TreeNode
    q = append(q, root.Left)
    q = append(q, root.Right)
    for len(q) > 0 {
        //add next level to array
        length := len(q)
        level := []int{}
        for i := 0; i < length; i++{
            node := q[0]
            var nodeVal int
            if node != nil{
                nodeVal = node.Val
                q = append(q, node.Left)
                q = append(q, node.Right)
            }
            level = append(level, nodeVal)
            q = q[1:]
        }
        left, right := 0, len(level) - 1
        for left < right {
            if level[left] != level[right] {
                return false
            }
            left++
            right--
        }
    }
    return true
}

func isSymmetricRecursively(root *TreeNode) bool {
    if root == nil {
        return true
    }
    return checkSymmetric(root.Left, root.Right)
}

func checkSymmetric(left, right *TreeNode) bool{
    if left == nil && right == nil { 
        return true
    }
    if (left == nil || right == nil) || (left.Val != right.Val) {
        return false
    }
    return checkSymmetric(left.Left, right.Right) && checkSymmetric(left.Right,right.Left)
}


// [1,2,2,3,4,4,3]
// [1,2,2,3,4,4,3]
// [1,2,2,null,3,null,3]
// [1,2,2,null,3,3,null]
// [1]
// [1,2,2,3,4,4,3,5,6,7,8,8,7,6,5]
// [1,2,2,3,4,4,3,5,6,7,8,8,7,7,5]
// [1,2,2,3,4,4,3,5,6,7,8,8,7,6,5,9]