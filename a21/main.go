package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func maxInt(x int, y int) int {
	return int(math.Max(float64(x), float64(y)))
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(reader, &n)

	var conditionInput int
	conditions := make([]int, n)
	points := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &conditionInput)
		conditions[i] = conditionInput - 1
		fmt.Fscan(reader, &points[i])
	}

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	for length := n - 2; length >= 0; length-- {
		for l := 0; l < n-length; l++ {
			r := l + length

			if l-1 >= 0 {
				point := 0
				if conditions[l-1] >= l && conditions[l-1] <= r {
					point = points[l-1]
				}
				dp[l][r] = dp[l-1][r] + point
			}

			if r+1 < n {
				point := 0
				if conditions[r+1] >= l && conditions[r+1] <= r {
					point = points[r+1]
				}
				dp[l][r] = maxInt(dp[l][r], dp[l][r+1]+point)
			}
		}
	}

	ans := -1
	for i := 0; i < n; i++ {
		ans = maxInt(ans, dp[i][i])
	}
	fmt.Println(ans)
}
