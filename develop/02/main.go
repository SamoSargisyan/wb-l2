package main

import "fmt"

func main() {
	str, err := Unpack(`qwe4\5`)
	fmt.Println(str, err)
}
