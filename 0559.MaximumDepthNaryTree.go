package main

func maxDepth559(root *Node) int {
    if root == nil {
        return 0
    }
    
    res := 0
    var q []*Node
    q = append(q,root)
    
    for len(q) > 0 {
        length := len(q)
        level := []int{}
        
        for i:= 0; i < length; i++{
            node := q[0]
            q = q[1:]
            level = append(level, node.Val)
            for _, v := range node.Children {
                q = append(q, v)
            }
        }
        if len(level) > 0 {
            res++
        }
    }
    
    return res
}