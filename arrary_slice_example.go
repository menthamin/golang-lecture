package main

import "fmt"

func main() {
	s := make([]int, 3)
	var a []int

	s[0] = 10

	fmt.Println(s)
	fmt.Println(a)
	a = append(a, 10)
	fmt.Println(a)
	a = append(a, 11)
	fmt.Println(a)

	l := a[0:1]
	fmt.Println((l))
	// Arrary
	// var a [5]int
	// // var b [5]string
	// // c := [5]int{1, 2, 3, 4, 5}
	// // var d [2][3]int

	// a[0] = 100
	// a[4] = 100
	// fmt.Println(a)

	// b := [3]int{1, 2, 3}
	// fmt.Println(b)

	// var c [2][3]int
	// fmt.Println(c)
	// c[0][2] = 100
	// fmt.Println(c)

	// Array: 정확히 정해진 개수
	// Slice: Python의 리스트와 가깝다.
	// 유연한 개수의 배열
	// 배열보다 자주 사용
	// 배열의 확장판
	// Type은 배열과 다름

	// Slice
	// var a []int
	// var b []string

	// s := make([]int, 3)
	// fmt.Println(a, b, s)

	// Map: Key, Value 자료형
	// Python의 Dict와 비슷
	// Set 자료형은 없을까?

	// s := make([]int, 3)
	// m := make(map[string]int)
}
