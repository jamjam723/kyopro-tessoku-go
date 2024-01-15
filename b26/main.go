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

	prime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		prime[i] = true
	}

	for i := 2; i <= n; i++ {
		if !prime[i] {
			continue
		}
		prime[i] = true
		for j := i * 2; j <= n; j += i {
			prime[j] = false
		}
	}

	for i := 2; i <= n; i++ {
		if prime[i] {
			fmt.Println(i)
		}
	}
}
