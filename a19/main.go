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
	var n, w int
	fmt.Fscan(reader, &n, &w)

	weights := make([]int, n+1)
	values := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(reader, &weights[i], &values[i])
	}

	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, w+1)
	}
	for i := 0; i <= w; i++ {
		dp[0][i] = -1000000000
	}
	dp[0][0] = 0

	for i := 1; i <= n; i++ {
		for j := 1; j <= w; j++ {
			dp[i][j] = dp[i-1][j]
			if j-weights[i] >= 0 {
				dp[i][j] = maxInt(dp[i][j], dp[i-1][j-weights[i]]+values[i])
			}
		}
	}

	ans := -1
	for i := 1; i <= w; i++ {
		ans = maxInt(ans, dp[n][i])
	}
	fmt.Println(ans)
}
