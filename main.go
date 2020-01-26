package main

import (
	"exclaim"
	"fmt"
	"scanner"
	"token"
	"yell"
)

func main() {
	fmt.Println("Sanity Check")
	y := yell.Yell("Sanity Check")
	e := exclaim.Exclaim("Sanity Check")
	t := token.SanityCheck()
	s := scanner.New("Some input")
	fmt.Println(y)
	fmt.Println(e)
	fmt.Println(t)
	fmt.Println(s.Source())
}
