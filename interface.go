package main

import "fmt"

type Person interface {
	getName() string
}

type Teen interface {
	getAge() int
}

type Student struct {
	name string
	age  int
}

func (stu *Student) getName() string {
	return stu.name
}

func (stu *Student) getAge() int {
	return stu.age
}

func main() {
	s := Student{
		name: "Tom",
		age:  18,
	}

	var p Person = &s
	var t Teen = &s

	fmt.Println(p.getName()) // Tom
	fmt.Println(t.getAge())  // 18
}
