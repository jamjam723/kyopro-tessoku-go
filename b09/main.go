package main

import (
	"bufio"
	"fmt"
	"os"
)

const MAT_LEN = 1503

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(reader, &n)

	mat := make([][]int, MAT_LEN)
	for i := 0; i < MAT_LEN; i++ {
		vec := make([]int, MAT_LEN)
		mat[i] = vec
	}

	var a, b, c, d int
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a, &b, &c, &d)
		mat[a][b] += 1
		mat[a][d] -= 1
		mat[c][b] -= 1
		mat[c][d] += 1
	}

	for i := 0; i < MAT_LEN; i++ {
		for j := 1; j < MAT_LEN; j++ {
			mat[i][j] += mat[i][j-1]
		}
	}

	for i := 0; i < MAT_LEN; i++ {
		for j := 1; j < MAT_LEN; j++ {
			mat[j][i] += mat[j-1][i]
		}
	}

	total := 0
	for i := 0; i < MAT_LEN; i++ {
		for j := 0; j < MAT_LEN; j++ {
			if mat[i][j] > 0 {
				total += 1
			}
		}
	}
	fmt.Println(total)
}
