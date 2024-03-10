package template

import "github.com/valentinrog/templatinator/element/common/children"

type Template struct {
	children.Children[Template]
}

func New() *Template {
	e := &Template{}
	e.Children = children.Create(e)
	return e
}

func (p Template) Stringify() string {
	return p.Children.Stringify()
}
