package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var a, b int
	fmt.Fscan(reader, &a, &b)

	for i := a; i <= b; i++ {
		if 100%i == 0 {
			fmt.Println("Yes")
			return
		}
	}

	fmt.Println("No")
}
