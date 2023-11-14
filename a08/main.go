package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var h, w int
	fmt.Fscan(reader, &h, &w)

	mat := make([][]int, h+1)
	for i := 0; i <= h; i++ {
		vec := make([]int, w+1)
		mat[i] = vec
	}

	for i := 1; i <= h; i++ {
		for j := 1; j <= w; j++ {
			fmt.Fscan(reader, &mat[i][j])
		}
	}

	for i := 1; i <= h; i++ {
		for j := 1; j <= w; j++ {
			mat[i][j] += mat[i][j-1]
		}
	}

	for i := 1; i <= w; i++ {
		for j := 1; j <= h; j++ {
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
