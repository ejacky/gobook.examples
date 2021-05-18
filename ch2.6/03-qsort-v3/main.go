// Copyright © 2018 ChaiShushan <chaishushan{AT}gmail.com>.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// +build ignore

package main

import (
	qsort2 "examples/ch2.6/03-qsort-v3"
	"fmt"
	"unsafe"
)

func main() {
	values := []int32{42, 9, 101, 95, 27, 25}

	qsort2.Sort(unsafe.Pointer(&values[0]), len(values), int(unsafe.Sizeof(values[0])),
		func(a, b unsafe.Pointer) int {
			pa, pb := (*int32)(a), (*int32)(b)
			return int(*pa - *pb)
		},
	)

	fmt.Println(values)
}
