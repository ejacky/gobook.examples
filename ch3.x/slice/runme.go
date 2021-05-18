// Copyright © 2017 ChaiShushan <chaishushan{AT}gmail.com>.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// +build ignore

package main

import (
	"examples/ch3.x/slice"
)

func main() {
	println("SumIntSlice([]int{1,2,3}) =", slice.SumIntSlice([]int{1, 2, 3}))
	println("AsmSumIntSlice([]int{1,2,3}) =", slice.AsmSumIntSlice([]int{1, 2, 3}))
}
