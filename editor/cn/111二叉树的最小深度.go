// 给定一个二叉树，找出其最小深度。
//
// 最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
//
// 说明：叶子节点是指没有子节点的节点。
//
//
//
// 示例 1：
//
//
// 输入：root = [3,9,20,null,null,15,7]
// 输出：2
//
//
// 示例 2：
//
//
// 输入：root = [2,null,3,null,4,null,5,null,6]
// 输出：5
//
//
//
//
// 提示：
//
//
// 树中节点数的范围在 [0, 10⁵] 内
// -1000 <= Node.val <= 1000
//
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 782 👎 0

package main

import "fmt"

func main() {
	result := minDepth(new(TreeNode))
	fmt.Println(result)
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
func minDepth(root *TreeNode) int {
	var depth func(node *TreeNode) int
	depth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		if node.Left == nil && node.Right != nil {
			return 1 + depth(node.Right)
		}
		if node.Left != nil && node.Right == nil {
			return 1 + depth(node.Left)
		}
		return min(depth(node.Left), depth(node.Right)) + 1

	}
	return depth(root)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a

}

// leetcode submit region end(Prohibit modification and deletion)
/*
	递归终止条件:
		遍历到叶子节点则返回

	递归逻辑:
		左子树不存在,右字数存在,则遍历右子树
		左子树存在,右子树不存在,则遍历左子树
		左右子树都不存在或都存在,则递归取其最小值+1
*/
