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

	color := make([]string, h+1)
	for i := 1; i <= h; i++ {
		fmt.Fscan(reader, &color[i])
	}

	routes := make([][]int, h+1)
	for i := 0; i <= h; i++ {
		routes[i] = make([]int, w+1)
	}

	routes[1][1] = 1
	for i := 1; i <= h; i++ {
		for j := 1; j <= w; j++ {
			if i == 1 && j == 1 {
				continue
			}

			if string(color[i][j-1]) == "#" {
				continue
			}

			routes[i][j] = routes[i-1][j] + routes[i][j-1]
		}
	}

	fmt.Println(routes[h][w])
}
