package text

type Text struct {
	text string
}

func New() *Text {
	return &Text{}
}

func (e *Text) Set(text string) *Text {
	e.text = text
	return e
}

func (e Text) Stringify() string {
	return e.text
}
