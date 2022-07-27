package main

import (
	"fmt"
	"math"
)

/*
	假设我们有一个 n 乘以 n 的矩阵 w[n][n]。矩阵存储的都是正整数。棋子起始位置在左上角，终止位置在右下角。
	我们将棋子从左上角移动到右下角。每次只能向右或者向下移动一位。从左上角到右下角，会有很多不同的路径可以走。
	我们把每条路径经过的数字加起来看作路径的长度。那从左上角移动到右下角的最短路径长度是多少呢？
*/

func main() {
	var checkerBoard = [][]int{
		{1, 3, 5, 9},
		{2, 1, 3, 4},
		{5, 2, 6, 7},
		{6, 8, 4, 3},
	}
	minDist := minDistBT(checkerBoard, 3)
	fmt.Println(minDist)
}

func minDistBT(checkerBoard [][]int, n int) int {
	var (
		minDist = math.MaxInt
		minF    func(i, j, dist int)
	)
	minF = func(i, j, dist int) {
		if (i == n) && (j == n) {
			if dist < minDist {
				minDist = dist
			}
			return
		}
		if i < n {
			// 1.
			minF(i+1, j, dist+checkerBoard[i+1][j])
			// 2.
			// dist += checkerBoard[i+1][j]
			// minF(i+1, j, dist)
		}
		if j < n {
			// 1.
			minF(i, j+1, dist+checkerBoard[i][j+1])
			// 2.
			// dist += checkerBoard[i][j+1]
			// minF(i, j+1, dist)
		}
	}
	minF(0, 0, checkerBoard[0][0])
	return minDist
}

func minDistBT1(checkerBoard [][]int, n int) int {
	dbTable := make([][]int, n)
	for i := 0; i < n; n++ {
		dbTable[i] = make([]int, n)
	}
	sum := 0
	for j := 0; j < n; j++ {
		sum += checkerBoard[0][j]
		dbTable[0][j] = sum
	}
	sum = 0
	for i := 0; i < n; i++ {
		sum += checkerBoard[i][0]
		dbTable[i][0] = sum
	}

	for i := 1; i < n; i++ {
		for j := 1; j < n; j++ {
			dbTable[i][j] = dbTable[i][j] + min(dbTable[i][j-1], dbTable[i-1][j])
		}
	}
	return dbTable[n-1][n-1]
}
