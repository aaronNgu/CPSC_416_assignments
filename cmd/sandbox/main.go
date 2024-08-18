package main

import (
	"fmt"
)

type Some struct {
	key string
}

func main() {
	var a Some
	fmt.Println("a is ")
	fmt.Println(a)
	//fmt.Println(a == {})
}
