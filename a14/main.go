package main

import (
	"bufio"
	"fmt"
	"os"

	"golang.org/x/exp/slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, k int
	fmt.Fscan(reader, &n, &k)

	boxA := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &boxA[i])
	}
	boxB := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &boxB[i])
	}
	boxC := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &boxC[i])
	}
	boxD := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &boxD[i])
	}

	aPlusB := make([]int, n*n)
	cPlusD := make([]int, n*n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			aPlusB[i*n+j] = boxA[i] + boxB[j]
			cPlusD[i*n+j] = boxC[i] + boxD[j]
		}
	}

	slices.Sort(cPlusD)

	ok := false
	for i := 0; i < n*n; i++ {
		_, ok = slices.BinarySearch(cPlusD, k-aPlusB[i])
		if ok {
			break
		}
	}

	if ok {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
