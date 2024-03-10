package children

import (
	"strings"

	"github.com/valentinrog/templatinator/element/tag"
)

type Children[T tag.Tag] struct {
	self     *T
	children []tag.Tag
}

func Create[T tag.Tag](self *T) Children[T] {
	return Children[T]{
		self: self,
	}
}

func (e *Children[T]) Put(children ...tag.Tag) *T {
	e.children = append(e.children, children...)
	return e.self
}

func (e *Children[T]) Stringify() string {
	sb := strings.Builder{}
	for _, c := range e.children {
		sb.WriteString(c.Stringify())
	}
	return sb.String()
}
