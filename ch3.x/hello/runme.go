// Copyright © 2017 ChaiShushan <chaishushan{AT}gmail.com>.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// +build ignore

package main

import (
	"examples/ch3.x/hello"
	"fmt"
)

func main() {
	s := "你好, 世界!\n"
	fmt.Printf("%d: %x\n", len(s), s)
	hello.PrintHelloWorld()
	hello.PrintHelloWorld_zh()
	hello.PrintHelloWorld_var()
}
