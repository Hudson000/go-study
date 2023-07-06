package main

import (
	"fmt"
	"os"
	"strconv"
	"test/popcount"
)

func test07() {
	x := os.Args[1]
	y, err := strconv.ParseUint(x, 10, 64)
	if err != nil {
		return
	}
	fmt.Println(popcount.PopCount(y))
}
