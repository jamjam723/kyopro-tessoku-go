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
	var s, t string
	fmt.Fscan(reader, &s, &t)

	dp := make([][]int, len(s)+1)
	for i := 0; i <= len(s); i++ {
		dp[i] = make([]int, len(t)+1)
	}

	for i := 0; i <= len(s); i++ {
		for j := 0; j <= len(t); j++ {
			dp[i][j] = math.MaxInt32
		}
	}
	dp[0][0] = 0

	for i := 0; i <= len(s); i++ {
		for j := 0; j <= len(t); j++ {
			if i > 0 && j > 0 {
				cost := 1
				if s[i-1] == t[j-1] {
					cost = 0
				}
				dp[i][j] = minInt(dp[i][j], dp[i-1][j-1]+cost)
			}

			if i > 0 {
				dp[i][j] = minInt(dp[i][j], dp[i-1][j]+1)
			}

			if j > 0 {
				dp[i][j] = minInt(dp[i][j], dp[i][j-1]+1)
			}
		}
	}

	fmt.Println(dp[len(s)][len(t)])
}
