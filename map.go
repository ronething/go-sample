package main

import (
	"fmt"
)

func hello() {
	nums := []int{10, 20, 30, 40}

	for num := range nums {
		fmt.Println(num)
	}
}

func main() {
	hello()
}
