package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(reader, &n)

	fromPrev := make([]int, n)
	for i := 1; i <= n-1; i++ {
		fmt.Fscan(reader, &fromPrev[i])
	}
	fromTwoPosBefore := make([]int, n)
	for i := 2; i <= n-1; i++ {
		fmt.Fscan(reader, &fromTwoPosBefore[i])
	}

	dp := make([]int, n)
	dp[0] = 0
	dp[1] = dp[0] + fromPrev[1]
	for i := 2; i <= n-1; i++ {
		dp[i] = int(math.Min(float64(dp[i-1]+fromPrev[i]), float64(dp[i-2]+fromTwoPosBefore[i])))
	}
	fmt.Println(dp[n-1])
}
