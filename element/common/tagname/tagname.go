package tagname

import "github.com/valentinrog/templatinator/tag"

type TagName[T tag.Tag] struct {
	self T
	name string
}

func Create[T tag.Tag](self T, name string) TagName[T] {
	return TagName[T]{self: self, name: name}
}

func (e *TagName[T]) GetTagName() string {
	return e.name
}
