// 给你二叉树的根节点 root ，返回它节点值的 前序 遍历。
//
//
//
// 示例 1：
//
//
// 输入：root = [1,null,2,3]
// 输出：[1,2,3]
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
// 示例 4：
//
//
// 输入：root = [1,2]
// 输出：[1,2]
//
//
// 示例 5：
//
//
// 输入：root = [1,null,2]
// 输出：[1,2]
//
//
//
//
// 提示：
//
//
// 树中节点数目在范围 [0, 100] 内
// -100 <= Node.val <= 100
//
//
//
//
// 进阶：递归算法很简单，你可以通过迭代算法完成吗？
// Related Topics 栈 树 深度优先搜索 二叉树 👍 742 👎 0

package main

import (
	"container/list"
	"fmt"
)

func main() {
	node2 := TreeNode{Val: 2}
	node1 := TreeNode{Val: 1}
	node := TreeNode{Val: 0, Left: &node1, Right: &node2}
	result := preorderTraversal(&node)
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
func preorderTraversal(root *TreeNode) []int {
	var ans []int
	var traver func(root *TreeNode)
	traver = func(root *TreeNode) {
		if root == nil {
			return
		}
		ans = append(ans, root.Val)
		traver(root.Left)
		traver(root.Right)
	}
	traver(root)
	return ans
}

func preorderTraversal1(root *TreeNode) []int {
	var ans []int
	if root == nil {
		return ans
	}
	queue := []*TreeNode{root}
	if len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		ans = append(ans, node.Val)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return ans
}

func preorderTraversal2(root *TreeNode) []int {
	var ans []int
	if root == nil {
		return ans
	}
	queue := list.New()
	queue.PushBack(root)
	if queue.Len() > 0 {
		node := queue.Remove(queue.Back()).(*TreeNode)
		ans = append(ans, node.Val)
		if node.Left != nil {
			queue.PushBack(node.Left)
		}
		if node.Right != nil {
			queue.PushBack(node.Right)
		}
	}
	return ans
}

// leetcode submit region end(Prohibit modification and deletion)
