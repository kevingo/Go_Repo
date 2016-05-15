package main

import (
	"fmt"
)

func multiply(a, b int, reply *int) {
	*reply = a * b
}

func main() {
	n := 0
	reply := &n
	multiply(10, 5, reply)
	fmt.Println("Result = ", *reply)
}