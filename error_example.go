package main

import (
	"errors"
	"fmt"
	"os"
)

func a(b int) (int, error) {
	if b >= 10 {
		// 10보다 크면 -1과 error message 전달
		return -1, errors.New("don't exists")
	}
	return b, nil
}

func main() {
	// error
	// fmt.Println(a(10))
	// fmt.Println(a(5))

	f, err := os.Create("data/defer.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	fmt.Fprintln(f, "Something")
	fmt.Fprintln(f, "Something")
	fmt.Fprintln(f, "Something")
	fmt.Printf("f: %v\n", f)
}
