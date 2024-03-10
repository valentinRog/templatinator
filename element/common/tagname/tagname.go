package tagname

import "github.com/valentinrog/templatinator/element/tag"

type TagName[T tag.Tag] struct {
	self *T
	name string
}

func Create[T tag.Tag](self *T, name string) TagName[T] {
	return TagName[T]{
		self: self,
		name: name,
	}
}

func (t *TagName[T]) GetTagName() string {
	return t.name
}
