package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func minInt(x int, y int) int {
	return int(math.Min(float64(x), float64(y)))
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, w int
	fmt.Fscan(reader, &n, &w)

	weights := make([]int, n+1)
	values := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(reader, &weights[i], &values[i])
	}

	maxValue := 100 * 1000
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, maxValue+1)
	}
	for i := 0; i <= maxValue; i++ {
		dp[0][i] = 1000000001
	}
	dp[0][0] = 0

	for i := 1; i <= n; i++ {
		for j := 0; j <= maxValue; j++ {
			dp[i][j] = dp[i-1][j]
			if j-values[i] >= 0 {
				dp[i][j] = minInt(dp[i][j], dp[i-1][j-values[i]]+weights[i])
			}
		}
	}

	ans := -1
	for i := maxValue; i >= 0; i-- {
		if dp[n][i] <= w {
			ans = i
			break
		}
	}
	fmt.Println(ans)
}
