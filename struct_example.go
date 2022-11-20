package main

import "fmt"

type person struct {
	name string
	age  int
	job  string
}

func main() {
	a1 := person{"Austin", 30, "programmer"}
	a2 := person{name: "Violet", age: 28, job: "Designer"}
	a3 := person{name: "Sopia", age: 30, job: "Product Owner"}

	fmt.Println(a1, a2, a3)
	fmt.Println(a1.name, a2.age, a3.job)

	a1.age = 31
	fmt.Println(a1.age)
}
