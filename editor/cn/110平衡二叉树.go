// 给定一个二叉树，判断它是否是高度平衡的二叉树。
//
// 本题中，一棵高度平衡二叉树定义为：
//
//
// 一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1 。
//
//
//
//
// 示例 1：
//
//
// 输入：root = [3,9,20,null,null,15,7]
// 输出：true
//
//
// 示例 2：
//
//
// 输入：root = [1,2,2,3,3,null,null,4,4]
// 输出：false
//
//
// 示例 3：
//
//
// 输入：root = []
// 输出：true
//
//
//
//
// 提示：
//
//
// 树中的节点数在范围 [0, 5000] 内
// -10⁴ <= Node.val <= 10⁴
//
// Related Topics 树 深度优先搜索 二叉树 👍 1053 👎 0

package main

import (
	"fmt"
	"math"
)

func main() {
	result := isBalanced(new(TreeNode))
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
func isBalanced(root *TreeNode) bool {
	var balance func(node *TreeNode) int
	balance = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := balance(node.Left)
		right := balance(node.Right)
		if left == -1 || right == -1 {
			return -1
		}
		if math.Abs(float64(left-right)) > 1 {
			return -1
		}
		return int(math.Max(float64(left), float64(right))) + 1
	}
	return balance(root) != -1
}

// leetcode submit region end(Prohibit modification and deletion)

/*
节点的高度：节点到最远叶子节点的最长路径上边的数量。叶子节点高度为0。
节点的深度：节点到根节点的路径上边的数量。所有根节点深度为0。
树的高度：树的高度等于根节点的高度，等于最远叶子节点的深度。
树的深度：树的深度等于树的高度。
树的宽度：两个最长路径的叶子节点之间节点数。

    自底向上 `剪枝`
    递归返回值:
        如果当前节点的左右子树高度差<2,则返回当前节点的最大高度
        如果当前节点的左右子树高度差>=2,则返回-1,此子树不是平衡树
    递归终止条件:
        叶子节点时,返回高度为0
        左右子树返回有一个是-1时,退出 //剪枝
*/
