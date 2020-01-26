package main

import (
	"exclaim"
	"fmt"
	"token"
	"yell"
)

func main() {
	fmt.Println("Sanity Check")
	y := yell.Yell("Sanity Check")
	e := exclaim.Exclaim("Sanity Check")
	t := token.SanityCheck()
	fmt.Println(y)
	fmt.Println(e)
	fmt.Println(t)
}
