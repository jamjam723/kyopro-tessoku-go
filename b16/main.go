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

func absInt(x int) int {
	return int(math.Abs(float64(x)))
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(reader, &n)

	heights := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &heights[i])
	}

	dp := make([]int, n)
	dp[0] = 0
	dp[1] = absInt(heights[1] - heights[0])
	for i := 2; i < n; i++ {
		dp[i] = minInt(dp[i-1]+absInt(heights[i]-heights[i-1]), dp[i-2]+absInt(heights[i]-heights[i-2]))
	}
	fmt.Println(dp[n-1])
}
