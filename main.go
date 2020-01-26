package main 

import (
"exclaim"
"yell"
	"fmt"
)


func main() {
    fmt.Println("Sanity Check")
    y := yell.Yell("Sanity Check")
    e := exclaim.Exclaim("Sanity Check")
    fmt.Println(y)
    fmt.Println(e)
}
