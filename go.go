package main

import "fmt"

func main() {
	var name string
	fmt.Println("안녕하세요! 별다방입니다.")
	fmt.Println("이름이 무엇인가요?")
	fmt.Scanln(&name)
	fmt.Println("이름이 ", name, "이군요! 예쁜 이름이네요.")
}
