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
	var s, t string
	fmt.Fscan(reader, &s, &t)

	dp := make([][]int, len(s)+1)
	for i := 0; i <= len(s); i++ {
		dp[i] = make([]int, len(t)+1)
	}

	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(t); j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
			dp[i][j] = maxInt(dp[i][j], dp[i-1][j])
			dp[i][j] = maxInt(dp[i][j], dp[i][j-1])
		}
	}

	fmt.Println(dp[len(s)][len(t)])
}
