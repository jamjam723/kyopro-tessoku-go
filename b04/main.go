package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n string
	fmt.Fscan(reader, &n)

	ans := 0
	for i := 0; i < len(n); i++ {
		if string(n[i]) == "1" {
			ans += int(math.Pow(2, float64(len(n)-1-i)))
		}
	}
	fmt.Println(ans)
}
