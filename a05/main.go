package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, k int
	fmt.Fscan(reader, &n, &k)

	count := 0
	for a := 1; a <= n; a++ {
		for b := 1; b <= n; b++ {
			c := k - (a + b)
			if c >= 1 && c <= n {
				count++
			}
		}
	}

	fmt.Println(count)
}
