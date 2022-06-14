// 给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。
//
//
//
// 示例 1：
//
//
// 输入：root = [3,9,20,null,null,15,7]
// 输出：[[3],[9,20],[15,7]]
//
//
// 示例 2：
//
//
// 输入：root = [1]
// 输出：[[1]]
//
//
// 示例 3：
//
//
// 输入：root = []
// 输出：[]
//
//
//
//
// 提示：
//
//
// 树中节点数目在范围 [0, 2000] 内
// -1000 <= Node.val <= 1000
//
// Related Topics 树 广度优先搜索 二叉树 👍 1357 👎 0

package main

import (
	"container/list"
	"fmt"
)

func main() {
	node5 := TreeNode{Val: 7}
	node4 := TreeNode{Val: 15}
	node3 := TreeNode{Val: 20, Left: &node4, Right: &node5}
	node2 := TreeNode{Val: 9}
	node1 := TreeNode{Val: 3, Left: &node2, Right: &node3}
	result := levelOrder(&node1)
	fmt.Println(result)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// go-leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	queue := list.New()
	queue.PushBack(root)
	var ans [][]int
	for queue.Len() > 0 {
		var val []int
		n := queue.Len()
		for n > 0 {
			node := queue.Remove(queue.Front()).(*TreeNode)
			val = append(val, node.Val)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			n--
		}
		ans = append(ans, val)
	}
	return ans
}

// go-leetcode submit region end(Prohibit modification and deletion)
