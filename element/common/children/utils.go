package children

import (
	"github.com/valentinrog/templatinator/tag"
	"strings"
)

type Utils[T tag.Tag] struct {
	children *Children[T]
}

func CreateUtils[T tag.Tag](children *Children[T]) Utils[T] {
	return Utils[T]{children: children}
}

func (e Utils[T]) String() string {
	sb := strings.Builder{}
	for _, c := range e.children.children {
		sb.WriteString(c.Stringify())
	}
	return sb.String()
}
