package attributes

import (
	"fmt"
	"github.com/valentinrog/templatinator/tag"
	"strings"
)

type Attributes[T tag.Tag] struct {
	self       T
	attributes map[string]string
}

func Create[T tag.Tag](self T) Attributes[T] {
	return Attributes[T]{self: self, attributes: map[string]string{}}
}

func (e *Attributes[T]) SetAttr(k, v string) T {
	e.attributes[k] = v
	return e.self
}

func (e *Attributes[T]) String() string {
	sb := strings.Builder{}
	for k, v := range e.attributes {
		sb.WriteRune(' ')
		sb.WriteString(fmt.Sprintf("%s=\"%s\"", k, v))
	}
	return sb.String()
}
