// Copyright © 2018 ChaiShushan <chaishushan{AT}gmail.com>.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// +build ignore

package main

//extern int go_qsort_compare(void* a, void* b);
import "C"

import (
	qsort2 "examples/ch2.6/02-qsort-v2"
	"fmt"
	"unsafe"
)

func main() {
	values := []int32{42, 9, 101, 95, 27, 25}

	qsort2.Sort(unsafe.Pointer(&values[0]),
		len(values), int(unsafe.Sizeof(values[0])),
		qsort2.CompareFunc(C.go_qsort_compare),
	)
	fmt.Println(values)
}

//export go_qsort_compare
func go_qsort_compare(a, b unsafe.Pointer) C.int {
	pa, pb := (*C.int)(a), (*C.int)(b)
	return C.int(*pa - *pb)
}
