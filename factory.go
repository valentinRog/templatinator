package templatinator

import (
	"github.com/valentinrog/templatinator/element/parentable"
	"github.com/valentinrog/templatinator/element/selfclosing"
	"github.com/valentinrog/templatinator/element/template"
	"github.com/valentinrog/templatinator/element/text"
)

type Factory struct{}

//custom elements

func (Factory) Text() *text.Text {
	return text.New()
}

func (Factory) Template() *template.Template {
	return template.New()
}

//standard html selfclosing elements

func (Factory) Link() *selfclosing.SelfClosing {
	return selfclosing.New("link")
}

func (Factory) Input() *selfclosing.SelfClosing {
	return selfclosing.New("input")
}

//standard html parentable elements

func (Factory) Div() *parentable.Parentable {
	return parentable.New("div")
}

func (Factory) Ul() *parentable.Parentable {
	return parentable.New("ul")
}

func (Factory) Li() *parentable.Parentable {
	return parentable.New("li")
}

func (Factory) P() *parentable.Parentable {
	return parentable.New("p")
}

func (Factory) Span() *parentable.Parentable {
	return parentable.New("span")
}

func (Factory) Strong() *parentable.Parentable {
	return parentable.New("strong")
}

func (Factory) Script() *parentable.Parentable {
	return parentable.New("script")
}

func (Factory) Html() *parentable.Parentable {
	return parentable.New("html")
}

func (Factory) Head() *parentable.Parentable {
	return parentable.New("head")
}

func (Factory) Body() *parentable.Parentable {
	return parentable.New("body")
}

func (Factory) Title() *parentable.Parentable {
	return parentable.New("title")
}

func (Factory) Button() *parentable.Parentable {
	return parentable.New("button")
}

func (Factory) Form() *parentable.Parentable {
	return parentable.New("form")
}

func (Factory) H1() *parentable.Parentable {
	return parentable.New("h1")
}
