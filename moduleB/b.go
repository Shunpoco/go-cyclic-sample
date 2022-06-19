package moduleB

import (
	ma "github.com/Shunpoco/go-cyclic-sample/moduleA"
)

type B struct {
	X int
	Y int
}

func (this *B) SumA(a *ma.A) {
	this.X += a.X
	this.Y += a.Y
}
