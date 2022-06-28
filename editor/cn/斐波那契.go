package main

func fb1(n int) int {
	var fb func(num int) int
	fb = func(num int) int {
		if num == 0 {
			return 0
		}
		if num == 1 || num == 2 {
			return 1
		}
		return fb(num-1) + fb(num-2)
	}
	return fb(n)
}

/*
从斐波那契看动态规划
	graph TD
    A[fb10]-->B[fb9]
    A[fb10]-->C[fb8]
    B[fb9]-->D[fb8]
    B[fb9]-->F[fb7]
    C[fb8]-->E[fb7]
    C[fb8]-->G[fb6]
递归问题先画出递归树,这样才能发现重复子问题

	动态规划主要问题是穷举出所有可能的值,然后求其中的最值.穷举的过程中往往会伴随着大量重复子问题的操作,导致算法效率低下.
	我们要解决的问题:
		1.解决重复子问题:
			通常我们可以用一个array 或者hashmap 去备忘录记录子问题的结果,每次迭代/递归时先检查一边备忘录
		2.最优子结构:
		3.状态转移:
			res = 1                    ; num=1/2
			res = fb(num-1) + fb(num-2); num>2
	###最困难是分析这个状态转移的过程###
*/

// 基于解决重复子问题去优化函数
// fb2
func fb2(n int) int {
	if n == 0 {
		return 0
	}
	cache := make([]int, n)
	var helper func(num int, c []int) int
	helper = func(num int, c []int) int {
		if num == 1 || num == 2 {
			return 1
		}
		if c[num] != 0 {
			return c[num]
		}
		c[num] = helper(num-1, c) + helper(num-2, c)
		return c[num]
	}
	return helper(n, cache)
}

/*
	fb2通过备忘录(DP table)优化了算法,但是是从自顶向下的逻辑思考,是否可以通过自底向上逻辑实现?
*/

// 基于自底向上
// fb3
func fb3(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	cache := make([]int, n)
	cache[1], cache[2] = 1, 1
	for i := 3; i < n; i++ {
		cache[i] = cache[i-1] + cache[i-2]
	}
	return cache[n]
}

/*
	由fb3发现DP table状态转移就用到了两个,进一步优化,压缩内存
*/

// 基于DP table 压缩
// fb4
func fb4(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	prev, curr := 1, 1
	for i := 3; i < n; i++ {
		sum := prev + curr
		prev = curr
		curr = sum
	}
	return curr
}
