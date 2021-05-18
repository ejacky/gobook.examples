// Copyright © 2017 ChaiShushan <chaishushan{AT}gmail.com>.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// +build ignore

package main

import (
	"examples/ch3.x/globalvar"
	"fmt"
)

func main() {
	fmt.Println(globalvar.GetPkgValue())
	fmt.Println(globalvar.GetPkgInfo())
}
