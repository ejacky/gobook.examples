// Copyright © 2017 ChaiShushan <chaishushan{AT}gmail.com>.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// +build ignore

package main

import (
	"examples/ch3.x/ifelse"
)

func main() {
	println("If(true, 1, 2) =", ifelse.If(true, 1, 2))
	println("If(false, 1, 2) =", ifelse.If(false, 1, 2))
	println("AsmIf(true, 1, 2) =", ifelse.AsmIf(true, 1, 2))
	println("AsmIf(false, 1, 2) =", ifelse.AsmIf(false, 1, 2))
	println("AsmIf(false, 2, 1) =", ifelse.AsmIf(false, 2, 1))
}
