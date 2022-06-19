package moduleA

import (
	mb "github.com/Shunpoco/go-cyclic-sample/moduleB"
)

type A struct {
	X int
	Y int
}

func (this *A) SumB(b *mb.B) {
	this.X += b.X
	this.Y += b.Y
}
