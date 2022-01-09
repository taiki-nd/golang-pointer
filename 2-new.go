package main

import "fmt"

func main() {
	//何も入れない状態でポインタが入るメモリの領域を確保したい
	var p *int = new(int)
	fmt.Println(p)
	fmt.Println(*p)
	*p++
	fmt.Println(*p)

	var p2 *int
	fmt.Println(p2) //宣言はしているがメモリを確保していない
	*p2++
	fmt.Println(p2)
}
