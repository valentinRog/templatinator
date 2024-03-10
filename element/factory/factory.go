package factory

import (
	"github.com/valentinrog/templatinator/element/parentable"
	"github.com/valentinrog/templatinator/element/text"
	"github.com/valentinrog/templatinator/element/template"
)

type Factory struct {
}

//custom elements

func (Factory) Text() *text.Text {
	return text.New()
}

func (Factory) Template() *template.Template {
	return template.New()
}

//standard elements

func (Factory) Div() *parentable.Parentable {
	return parentable.New("div")
}

func (Factory) Strong() *parentable.Parentable {
	return parentable.New("strong")
}

func (Factory) Span() *parentable.Parentable {
	return parentable.New("span")
}

func (Factory) Ul() *parentable.Parentable {
	return parentable.New("ul")
}

func (Factory) Li() *parentable.Parentable {
	return parentable.New("li")
}
