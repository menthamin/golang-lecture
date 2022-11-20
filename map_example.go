package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Americano"] = 3000
	m["Cafe latte"] = 3500

	fmt.Println((m))

	delete(m, "a")
	fmt.Println(m["Americano"])

	_, exi := m["Dolce Latte"]
	fmt.Println(exi)

}
