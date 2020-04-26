package main

import "fmt"

func main() {
	m := make(map[string]interface{})
	m["name"] = "Tom"
	m["age"] = 18
	m["scores"] = [3]int{98, 99, 85}
	fmt.Println(m) // map[age:18 name:Tom scores:[98 99 85]]
	fmt.Println(typeof(m["scores"]))
}

func typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}
