package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var q int
	fmt.Fscan(reader, &q)

	var x int
	for i := 0; i < q; i++ {
		fmt.Fscan(reader, &x)

		prime := true
		for k := 2; k <= int(math.Sqrt(float64(x))); k++ {
			if x%k == 0 {
				prime = false
				break
			}
		}
		if prime {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
