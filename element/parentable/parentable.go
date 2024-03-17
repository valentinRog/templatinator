package parentable

import (
	"fmt"
	"github.com/valentinrog/templatinator/element/common/attributes"
	"github.com/valentinrog/templatinator/element/common/children"
	"github.com/valentinrog/templatinator/element/common/tagname"
)

type Parentable struct {
	tagname.TagName[*Parentable]
	attributes.Attributes[*Parentable]
	children.Children[*Parentable]
}

func New(name string) *Parentable {
	e := &Parentable{}
	e.TagName = tagname.Create(e, name)
	e.Attributes = attributes.Create(e)
	e.Children = children.Create(e)
	return e
}

func (e *Parentable) Stringify() string {
	return fmt.Sprintf(
		"<%s %s>%s</%s>",
		e.GetTagName(),
		attributes.CreateUtils(&e.Attributes).String(),
		children.CreateUtils(&e.Children).String(),
		e.GetTagName(),
	)
}
