package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(reader, &n)

	var op string
	var a int
	acc := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &op)
		fmt.Fscan(reader, &a)

		r := a % 10000
		if op == "+" {
			acc += r
		} else if op == "-" {
			acc -= r
			if acc < 0 {
				acc += 10000
			}
		} else {
			acc *= r
		}

		acc %= 10000
		fmt.Println(acc)
	}
}
