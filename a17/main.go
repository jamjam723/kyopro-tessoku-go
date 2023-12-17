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
		dp[i] = minInt(dp[i-1]+fromPrev[i], dp[i-2]+fromTwoPosBefore[i])
	}

	routes := make([]int, 0, n)
	cur := n - 1
	routes = append(routes, n-1)
	for cur > 1 {
		if dp[cur-1]+fromPrev[cur] <= dp[cur-2]+fromTwoPosBefore[cur] {
			cur = cur - 1
		} else {
			cur = cur - 2
		}
		routes = append(routes, cur)
	}
	if cur == 1 {
		routes = append(routes, 0)
	}

	fmt.Println(len(routes))
	for i := len(routes) - 1; i >= 0; i-- {
		fmt.Print(routes[i] + 1)
		if i != 0 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
