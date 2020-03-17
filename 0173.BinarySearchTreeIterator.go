package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 type BSTIterator struct {
    values []int
    index int 
}


func Constructor(root *TreeNode) BSTIterator {
    values := make([]int, 0)
    inorder(root, &values)
    return BSTIterator{
		values: values,
		index:  0,
	}
}


func inorder(root *TreeNode, values *[]int) {
    if root == nil {
        return 
    }
    inorder(root.Left, values)
    *values = append(*values, root.Val)
    inorder(root.Right, values)
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
    if this.HasNext(){
        val := this.values[this.index]
        this.index++
        return val
    }
    return -1
}


/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
    return this.index < len(this.values)
}


/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */