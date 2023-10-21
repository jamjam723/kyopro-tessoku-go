package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, x int
	fmt.Fscan(reader, &n, &x)

	var a int
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a)
		if x == a {
			fmt.Println("Yes")
			return
		}
	}

	fmt.Println("No")
}
