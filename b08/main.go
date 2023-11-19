package main

import (
	"bufio"
	"fmt"
	"os"
)

const MAT_LEN = 1502

// サンプルが通っていない
func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(reader, &n)

	mat := make([][]int, MAT_LEN)
	for i := 0; i < MAT_LEN; i++ {
		vec := make([]int, MAT_LEN)
		mat[i] = vec
	}

	var x, y int
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &x, &y)
		mat[x][y] += 1
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

	var q int
	fmt.Fscan(reader, &q)
	var a, b, c, d int
	for i := 0; i < q; i++ {
		fmt.Fscan(reader, &a, &b, &c, &d)
		fmt.Println(mat[c][d] - mat[a-1][d] - mat[c][b-1] + mat[a-1][b-1])
	}
}
