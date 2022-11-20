package main

import "fmt"

func main() {
	fmt.Println(add(100, 11))

	first := add(1, 2)
	second := add_all(1, 2, 3, 4, 5)

	fmt.Println(first, second)
}

func add(a, b int) int {
	return a + b
}

func add_all(li ...int) int {
	result := 0
	for _, n := range li {
		result += n
	}
	return result
}
