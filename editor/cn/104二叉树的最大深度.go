// //给定一个二叉树，找出其最大深度。
// //
// // 二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
// //
// // 说明: 叶子节点是指没有子节点的节点。
// //
// // 示例：
// //给定二叉树 [3,9,20,null,null,15,7]，
// //
// // 3
// // / \
// // 9 20
// // / \
// // 15 7
// //
// // 返回它的最大深度 3 。
// // Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 1276 👎 0
//

package main

import "fmt"

func main() {
	result := maximumDepthOfBinaryTree(nil)
	fmt.Println(result)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	var ans int
	if root == nil {
		return ans
	}

	var depth func(node *TreeNode, n int)
	depth = func(node *TreeNode, n int) {
		ans = max(ans, n)
		if node == nil {
			return
		}
		if node.Left != nil {
			depth(node.Left, n+1)
		}
		if node.Right != nil {
			depth(node.Right, n+1)
		}
	}
	depth(root, 1)
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxDepth1(root *TreeNode) int {
	var ans int
	if root == nil {
		return ans
	}
	stack := []*TreeNode{root}
	for n := len(stack); n > 0; {
		m := n
		for i := 0; i < m; i++ {
			node := stack[i]
			if node.Left != nil {
				stack = append(stack, node.Left)

			}
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
		}
		stack = stack[m:]
		ans++
		n = len(stack)
	}
	return ans
}

// leetcode submit region end(Prohibit modification and deletion)
