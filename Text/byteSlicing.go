package main

import (
	"fmt"
)

func main() {

	s := "Himanshu Dogra" // bydefault it stores UTF8 character
	p := []byte(s)
	fmt.Printf("%v,%T\n", p, p)
}
