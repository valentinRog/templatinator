package attributes

import (
	"fmt"
	"strings"

	"github.com/valentinrog/templatinator/tag"
)

type Utils[T tag.Tag] struct {
	attributes *Attributes[T]
}

func CreateUtils[T tag.Tag](attributes *Attributes[T]) Utils[T] {
	return Utils[T]{attributes: attributes}
}

func (e Utils[T]) String() string {
	sb := strings.Builder{}
	for k, v := range e.attributes.attributes {
		sb.WriteRune(' ')
		sb.WriteString(fmt.Sprintf("%s=\"%s\"", k, v))
	}
	return sb.String()
}
