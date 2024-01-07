package main

import (
	"bufio"
	"fmt"
	"math"
	"os"

	"golang.org/x/exp/slices"
)

func maxInt(x int, y int) int {
	return int(math.Max(float64(x), float64(y)))
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(reader, &n)

	values := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &values[i])
	}

	dp := make([]int, n)
	minValues := make([]int, n+1)
	for i := 1; i <= n; i++ {
		minValues[i] = 1000000
	}

	for i := 0; i < n; i++ {
		pos, _ := slices.BinarySearch(minValues, values[i])
		dp[i] = maxInt(1, pos)
		if minValues[dp[i]] > values[i] {
			minValues[dp[i]] = values[i]
		}
	}

	ans := 0
	for i := 0; i < n; i++ {
		ans = maxInt(ans, dp[i])
	}
	fmt.Println(ans)
}
