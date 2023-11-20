package main

import (
	"bufio"
	"fmt"
	"os"
)

const MAT_LEN = 1502

func main() {
	reader := bufio.NewReader(os.Stdin)
	var h, w, n int
	fmt.Fscan(reader, &h, &w, &n)

	mat := make([][]int, MAT_LEN)
	for i := 0; i < MAT_LEN; i++ {
		mat[i] = make([]int, MAT_LEN)
	}

	var a, b, c, d int
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a, &b, &c, &d)
		mat[a][b] += 1
		mat[a][d+1] -= 1
		mat[c+1][b] -= 1
		mat[c+1][d+1] += 1
	}

	for i := 1; i < MAT_LEN; i++ {
		for j := 1; j < MAT_LEN; j++ {
			mat[i][j] += mat[i][j-1]
		}
	}

	for i := 1; i < MAT_LEN; i++ {
		for j := 1; j < MAT_LEN; j++ {
			mat[j][i] += mat[j-1][i]
		}
	}

	for i := 1; i <= h; i++ {
		for j := 1; j <= w; j++ {
			fmt.Print(mat[i][j])
			if j != w {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
