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
	var n, m int
	fmt.Fscan(reader, &n, &m)

	mat := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		mat[i] = make([]int, n)
	}
	for i := 1; i <= m; i++ {
		for j := 0; j < n; j++ {
			fmt.Fscan(reader, &mat[i][j])
		}
	}

	inf := 100000000
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, 1<<n)
	}
	for i := 0; i <= m; i++ {
		for j := 0; j < 1<<n; j++ {
			dp[i][j] = inf
		}
	}
	dp[0][0] = 0

	for i := 1; i <= m; i++ {
		for j := 0; j < 1<<n; j++ {
			selection := make([]int, n)
			for k := 0; k < n; k++ {
				if (j/(1<<k))%2 == 1 {
					selection[k] = 1
				}
			}

			v := 0
			for k := 0; k < n; k++ {
				if selection[k] == 1 || mat[i][k] == 1 {
					v += 1 << k
				}
			}

			dp[i][j] = minInt(dp[i][j], dp[i-1][j])
			dp[i][v] = minInt(dp[i][v], dp[i-1][j]+1)
		}
	}

	if dp[m][1<<n-1] >= inf {
		fmt.Println(-1)
	} else {
		fmt.Println(dp[m][1<<n-1])
	}
}
