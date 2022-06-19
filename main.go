package main

import (
	"fmt"

	ma "github.com/Shunpoco/go-cyclic-sample/moduleA"
	mb "github.com/Shunpoco/go-cyclic-sample/moduleB"
)

func main() {
	a := ma.A{X: 10, Y: 20}
	b := mb.B{X: 100, Y: 200}

	a.SumB(&b)

	fmt.Println(a, b)
}
