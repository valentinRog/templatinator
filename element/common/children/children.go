package children

import (
	"github.com/valentinrog/templatinator/tag"
)

type Children[T tag.Tag] struct {
	self     T
	children []tag.Tag
}

func Create[T tag.Tag](self T) Children[T] {
	return Children[T]{self: self}
}

func (e *Children[T]) AppendChildren(children ...tag.Tag) T {
	e.children = append(e.children, children...)
	return e.self
}
