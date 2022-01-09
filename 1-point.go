package main

import "fmt"

/*
func main() {
	var n int = 100
	fmt.Println(n)
	fmt.Println(&n) // 100を入れたメモリのアドレスが表示される

	var p *int = &n //intのポイント型
	fmt.Println(p)
	fmt.Println(*p) //アドレスが差しているメモリの中身が表示される
}
*/

func one(x *int) {
	*x = 1
}

func main() {
	var n int = 100
	one(&n)
	fmt.Println(n)
}
