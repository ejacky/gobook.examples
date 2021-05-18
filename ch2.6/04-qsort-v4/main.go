// Copyright © 2018 ChaiShushan <chaishushan{AT}gmail.com>.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// +build ignore

package main

import (
	qsort2 "examples/ch2.6/04-qsort-v4"
	"fmt"
)

func main() {
	values := []int64{42, 9, 101, 95, 27, 25}

	qsort2.Slice(values, func(i, j int) bool {
		return values[i] < values[j]
	})

	fmt.Println(values)
}
