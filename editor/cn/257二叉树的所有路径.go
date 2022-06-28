// 给你一个二叉树的根节点 root ，按 任意顺序 ，返回所有从根节点到叶子节点的路径。
//
// 叶子节点 是指没有子节点的节点。
//
//
// 示例 1：
//
//
// 输入：root = [1,2,3,null,5]
// 输出：["1->2->5","1->3"]
//
//
// 示例 2：
//
//
// 输入：root = [1]
// 输出：["1"]
//
//
//
//
// 提示：
//
//
// 树中节点的数目在范围 [1, 100] 内
// -100 <= Node.val <= 100
//
// Related Topics 树 深度优先搜索 字符串 回溯 二叉树 👍 764 👎 0

package main

import (
	"fmt"
	"strconv"
)

func main() {
	result := binaryTreePaths(new(TreeNode))
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
func binaryTreePaths(root *TreeNode) []string {
	var paths []string
	var constructPaths func(node *TreeNode, path string)
	constructPaths = func(node *TreeNode, path string) {
		if node != nil {
			pathSub := path
			pathSub += strconv.Itoa(node.Val)
			if node.Left == nil && node.Right == nil {
				paths = append(paths, pathSub)
			} else {
				pathSub += "->"
				constructPaths(node.Left, pathSub)
				constructPaths(node.Right, pathSub)
			}
		}
	}
	constructPaths(root, "")
	return paths
}

// leetcode submit region end(Prohibit modification and deletion)

/*
	递归:
		遍历时拼接path,直至叶子节点时,加入结果集中
*/
