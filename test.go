package main

import (
	"fmt"
)

func main() {
	str := make([]float32, 1)
	fmt.Println(len(str), cap(str))
	str = append(str, 1, 2, 3, 4, 5)
	fmt.Println(len(str), cap(str))
}
