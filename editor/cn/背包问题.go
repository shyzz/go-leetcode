package main

import "fmt"

/*
0-1背包问题:

	月黑风高的夜晚，张三开启了法外狂徒模式：他背着一个可装载重量为 W 的背包去地主家偷东西。

地主家有 N 个物品，每个物品有重量和价值两个属性，其中第 i 个物品的重量为 wt[i]，价值为 val[i]。

问张三现在用这个背包装物品，最多能装的价值是多少？

*/

func main() {
	w := 9
	nums := []int{2, 2, 4, 6, 3}
	// maxW := knapsack(w, nums)
	maxW := knapsack2(w, nums)
	fmt.Println(maxW)
}

// knapsack 回溯算法
// w: 背包最大重量, nums: 物品重量的集合
func knapsack(w int, nums []int) int {
	n := len(nums)
	maxW := 0
	var add func(i, currentW int)
	add = func(i, currentW int) {
		if (currentW == w) || (i == n) { // 当前重量等于最大重量/已经决策到最后一个物品
			if currentW > maxW {
				maxW = currentW
			}
			return
		}
		add(i+1, currentW) // 第i件物品不加入背包
		if (currentW + nums[i]) <= w {
			add(i+1, currentW+nums[i]) // 第i件物品加入背包
		}
	}
	add(0, 0)
	return maxW
}

// knapsack1 使用 DP table 优化
func knapsack1(w int, nums []int) int {
	n := len(nums)
	maxW := 0
	dbTable := make([][]bool, n)
	for i := 0; i < n; i++ {
		dbTable[i] = make([]bool, w+1)
	}
	var add func(i, currentW int)
	add = func(i, currentW int) {
		if (currentW == w) || (i == n) {
			if currentW > maxW {
				maxW = currentW
			}
			return
		}
		if dbTable[i][currentW] {
			return
		}
		dbTable[i][currentW] = true
		add(i+1, currentW)
		if (currentW + nums[i]) <= w {
			add(i+1, currentW+nums[i])
		}
	}
	add(0, 0)
	return maxW
}

// knapsack2 动态规划求解0-1背包问题
// 思路:
// 		循环遍历所有物品,检索到每次物品时都基于上一个物品所有的状态去做决策,是否加入背包
func knapsack2(w int, nums []int) int {
	n := len(nums)
	// 初始化dbTable
	dbTable := make([][]bool, n)
	for i := 0; i < n; i++ {
		dbTable[i] = make([]bool, w+1)
	}
	dbTable[0][0] = true // 第一行第一个数据需要初始化
	if nums[0] < w {
		dbTable[0][nums[0]] = true
	}
	for i := 1; i < n; i++ {
		for j := 0; j <= w; j++ { // dbTable 最大就只有w列
			if dbTable[i-1][j] { // 这个判断是为了判断上一个物品是否有满足条件的解, 此例中不判断也可,相当于穷举
				dbTable[i][j] = dbTable[i-1][j] // 第i个物品不加入背包,所有状态表状态和上一个物品一致
			}
		}
		for j := 0; j <= w-nums[i]; j++ { // 当前物品需要加入背包,所以判断时要判断列小于剩余重量
			if dbTable[i-1][j] {
				dbTable[i][j+nums[i]] = true
			}
		}
	}
	for i := w; i >= 0; i-- {
		if dbTable[n-1][i] {
			return i
		}
	}
	return 0
}

func knapsack3(w int, nums []int) int {
	n := len(nums)
	dbTable := make([]bool, w+1)
	dbTable[0] = true
	if nums[0] < w {
		dbTable[nums[0]] = true
	}
	for i := 1; i < n; i++ {
		for j := w - nums[i]; j <= w; j++ {
			if dbTable[j] {
				dbTable[j+nums[i]] = true
			}
		}
	}
	for i := w; i > 0; i-- {
		if dbTable[i] {
			return i
		}
	}
	return 0
}
