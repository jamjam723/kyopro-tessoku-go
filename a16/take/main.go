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
	var n int
	fmt.Fscan(reader, &n)

	next := make([]int, n)
	for i := 1; i <= n-1; i++ {
		fmt.Fscan(reader, &next[i])
	}
	nextnext := make([]int, n-1)
	for i := 1; i <= n-2; i++ {
		fmt.Fscan(reader, &nextnext[i])
	}

	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = 1000000000
	}
	dp[1] = 0

	for i := 1; i <= n-1; i++ {
		dp[i+1] = minInt(dp[i+1], dp[i]+next[i])
		if i+2 <= n {
			dp[i+2] = minInt(dp[i+2], dp[i]+nextnext[i])
		}
	}

	fmt.Println(dp[n])
}
