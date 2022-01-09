package main

import "fmt"

func main() {
	/*
	   Q1. 以下の実行結果をコードを書かずに答えてください。エラーがおきますか？正しく表示されるとすると何が表示字されますか？

	   package main

	   import (
	       "fmt"
	   )

	   func main() {
	       var i int = 10
	       var p *int
	       p = &i
	       var j int
	       j = *p
	       fmt.Println(j)
	   }
	*/
	var i int = 10
	var p *int
	p = &i
	fmt.Println(p) //メモリのアドレス
	var j int
	j = *p //中身を見にいく
	fmt.Println(j)

	/*
		Q2. 以下の実行結果をコードを書かずに答えてください。エラーがおきますか？正しく表示されるとすると何が表示されますか？

		package main

		import (
		    "fmt"
		)

		func main() {
		    var k int = 100
		    var l int = 200
		    var p1 *int
		    var p2 *int
		    p1 = &k
		    p2 = &l
		    k = *p1 + *p2
		    p2 = p1
		    l = *p2 + k
		    fmt.Println(j)
		}
	*/
	var k int = 100
	var l int = 200
	var p1 *int
	var p2 *int
	p1 = &k
	p2 = &l
	fmt.Println(*p1, *p2) //100 200
	k = *p1 + *p2
	fmt.Println(k)        //300  p1 = &kで*p1が300
	fmt.Println(*p1, *p2) //300
	p2 = p1
	l = *p2 + k
	fmt.Println(l)
}
