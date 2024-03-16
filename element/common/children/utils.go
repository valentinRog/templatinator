package children

import (
	"github.com/valentinrog/templatinator/tag"
	"strings"
)

type Utils[T tag.Tag] struct {
	Children *Children[T]
}

func (e Utils[T]) String() string {
	sb := strings.Builder{}
	for _, c := range e.Children.children {
		sb.WriteString(c.Stringify())
	}
	return sb.String()
}
