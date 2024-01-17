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
	mul := a * b

	for a > 0 && b > 0 {
		if a >= b {
			a = a % b
		} else {
			b = b % a
		}
	}

	var g int
	if a == 0 {
		g = b
	} else {
		g = a
	}

	fmt.Println(mul / g)
}
