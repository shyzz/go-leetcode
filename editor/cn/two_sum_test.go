/*
 * @lc app=leetcode.cn id=1 lang=golang
 * @lcpr version=
 *
 * [1] 两数之和
 */

package main

import "testing"

// @lc code=start
func twoSum(nums []int, target int) []int {
	// 维护 val -> index 的映射
	valToIndex := make(map[int]int)
	for i, num := range nums {
		// 查表，看看是否有能和 nums[i] 凑出 target 的元素
		need := target - num
		if j, ok := valToIndex[need]; ok {
			return []int{j, i}
		}
		// 存入 val -> index 的映射
		valToIndex[num] = i
	}
	return nil
}

// @lc code=end

func TestTwoSum(t *testing.T) {
	// your test code here
	nums := []int{2, 7, 11, 15}
	target := 9
	expected := []int{0, 1}
	result := twoSum(nums, target)
	if len(result) != len(expected) {
		t.Errorf("Expected length %d, got %d", len(expected), len(result))
	}

}

/*
// @lcpr case=start
// [2,7,11,15]\n9\n
// @lcpr case=end

// @lcpr case=start
// [3,2,4]\n6\n
// @lcpr case=end

// @lcpr case=start
// [3,3]\n6\n
// @lcpr case=end

*/
