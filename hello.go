package main

import "fmt"

func init() {
	fmt.Println("init!")
}

//init関数は最初に出力される
func bazz() {
	fmt.Println("Bazz")
}

func main() {
	bazz()
	fmt.Println("Hello world", "test")
}
