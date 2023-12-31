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

	adjacents1 := make([]int, n)
	for i := 1; i < n; i++ {
		fmt.Fscan(reader, &adjacents1[i])
	}
	adjacents2 := make([]int, n)
	for i := 1; i < n; i++ {
		fmt.Fscan(reader, &adjacents2[i])
	}

	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		dp[i] = -1000000000
	}

	for i := 1; i < n; i++ {
		dp[adjacents1[i]] = maxInt(dp[adjacents1[i]], dp[i]+100)
		dp[adjacents2[i]] = maxInt(dp[adjacents2[i]], dp[i]+150)
	}

	fmt.Println(dp[n])
}
