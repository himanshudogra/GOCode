package main

import (
	"fmt"
	"strconv"
)

var i int = 42

func main() {
	fmt.Printf("%v,%T\n", i, i)

	var j string
	j = strconv.Itoa(i)
	fmt.Printf("%v,%T\n", j, j)
}
