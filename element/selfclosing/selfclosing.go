package selfclosing

import (
	"fmt"
	"github.com/valentinrog/templatinator/element/common/attributes"
	"github.com/valentinrog/templatinator/element/common/tagname"
)

type SelfClosing struct {
	tagname.TagName[*SelfClosing]
	attributes.Attributes[*SelfClosing]
}

func New(name string) *SelfClosing {
	e := &SelfClosing{}
	e.TagName = tagname.Create(e, name)
	e.Attributes = attributes.Create(e)
	return e
}

func (e *SelfClosing) Stringify() string {
	return fmt.Sprintf("<%s %s />", e.GetTagName(), attributes.CreateUtils(&e.Attributes).String())
}
