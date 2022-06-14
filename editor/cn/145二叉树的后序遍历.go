// 给你一棵二叉树的根节点 root ，返回其节点值的 后序遍历 。
//
//
//
// 示例 1：
//
//
// 输入：root = [1,null,2,3]
// 输出：[3,2,1]
//
//
// 示例 2：
//
//
// 输入：root = []
// 输出：[]
//
//
// 示例 3：
//
//
// 输入：root = [1]
// 输出：[1]
//
//
//
//
// 提示：
//
//
// 树中节点的数目在范围 [0, 100] 内
// -100 <= Node.val <= 100
//
//
//
//
// 进阶：递归算法很简单，你可以通过迭代算法完成吗？
// Related Topics 栈 树 深度优先搜索 二叉树 👍 861 👎 0

package main

import (
	"container/list"
	"fmt"
)

func main() {
	node3 := TreeNode{Val: 2}
	node2 := TreeNode{Val: 1}
	node1 := TreeNode{Val: 0, Left: &node2, Right: &node3}
	result := postorderTraversal(&node1)
	fmt.Println(result)
}

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

// go-leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
	var ans []int
	var traver func(root *TreeNode)
	traver = func(root *TreeNode) {
		if root == nil {
			return
		}
		traver(root.Left)
		traver(root.Right)
		ans = append(ans, root.Val)
	}
	traver(root)
	return ans
}

func postorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var ans []int
	stack := list.New()
	stack.PushBack(root)
	if stack.Len() > 0 {
		node := stack.Remove(stack.Back()).(TreeNode)
		ans = append(ans, node.Val)
		if node.Left != nil {
			stack.PushBack(node.Left)
		}
		if node.Right != nil {
			stack.PushBack(node.Right)
		}
	}
	// 中左右  →  左右中
	for l, r := 0, len(ans)-1; l < r; l, r = l+1, l-1 {
		ans[l], ans[r] = ans[r], ans[l]
	}
	return ans
}

// go-leetcode submit region end(Prohibit modification and deletion)
