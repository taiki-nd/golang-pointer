package main

import "fmt"

type Vertex struct {
	X int
	Y int
	// X, Y intでも可能
	S string
}

/*
func main() {
	v := Vertex{X: 1, Y: 2}
	fmt.Println(v)
	fmt.Println(v.X, v.Y)
	v.X = 100
	fmt.Println(v.X, v.Y)

	v2 := Vertex{X: 1}
	fmt.Println(v2)

	v3 := Vertex{X: 1, Y: 1, S: "test"}
	fmt.Println(v3)

	v4 := Vertex{}
	fmt.Println(v4)
	fmt.Printf("%T %v\n", v4, v4)

	var v5 Vertex
	fmt.Println(v5) //v4 と同じ意味
	fmt.Printf("%T %v\n", v5, v5)

	v6 := new(Vertex)
	fmt.Println(v6) // =>&{0 0 }
	fmt.Printf("%T %v\n", v6, v6)

	v7 := &Vertex{}
	fmt.Println(v7)
	fmt.Printf("%T %v\n", v7, v7)

}
*/

func changeVertex(v Vertex) {
	v.X = 100
	//(*v).X = 100と本来書くべきであるがstructがやってくれる
}

func changeVertex2(v *Vertex) {
	v.X = 100
}

func main() {
	v := Vertex{1, 2, "test"}
	changeVertex(v)
	fmt.Println(v) // => {1, 2, "test"}となり変わらない、v:=で宣言しものに影響を与えない

	v2 := &Vertex{1, 2, "test"}
	changeVertex2(v2)
	fmt.Println(v2)  // => &{100 2 test}
	fmt.Println(*v2) // =>  {100 2 test}
}
