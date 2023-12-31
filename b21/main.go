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
	var s string
	fmt.Fscan(reader, &s)

	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i < n; i++ {
		dp[i][i+1] = 1
	}

	for i := 2; i <= n; i++ {
		for head := 0; head < n; head++ {
			tail := head + i
			if tail > n {
				continue
			}

			dp[head][tail] = dp[head][tail-1]
			dp[head][tail] = maxInt(dp[head][tail], dp[head+1][tail])
			if s[head] == s[tail-1] {
				dp[head][tail] = maxInt(dp[head][tail], dp[head+1][tail-1]+2)
			}
		}
	}

	fmt.Println(dp[0][n])
}
