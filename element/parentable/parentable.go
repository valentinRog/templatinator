package parentable

import (
	"fmt"

	"github.com/valentinrog/templatinator/element/common/children"
	"github.com/valentinrog/templatinator/element/common/tagname"
)

type Parentable struct {
	tagname.TagName[Parentable]
	children.Children[Parentable]
}

func New(name string) *Parentable {
	e := &Parentable{}
	e.TagName = tagname.Create(e, name)
	e.Children = children.Create(e)
	return e
}

func (p Parentable) Stringify() string {
	return fmt.Sprintf("<%s>%s</%s>", p.GetTagName(), p.Children.Stringify(), p.GetTagName())
}
