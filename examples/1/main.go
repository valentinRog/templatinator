package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"github.com/valentinrog/templatinator/element/factory"
	"github.com/valentinrog/templatinator/element/tag"
)

var html factory.Factory

func gen() string {
	title := html.Div().Put(
		html.Span().Put(
			html.Text().Set("CECI EST"),
			html.Strong().Put(html.Text().Set("MON")),
			html.Text().Set("TITRE"),
		),
		html.Span().Put(html.Text().Set("la suite du titre")),
		html.Div().Put(
			html.Strong().Put(
				html.Text().Set("strong"),
			),
		),
	)

	return html.Div().Put(
		title,
		html.Text().Set("yo dans la div 1"),
		html.Div(),
		html.Div().Put(
			func() []tag.Tag {
				var a []tag.Tag
				for i := 0; i < 10; i += 1 {
					a = append(a, html.Div().Put(
						html.Text().Set(fmt.Sprint(i)),
					))
				}
				return a
			}()...,
		),

		html.Ul().Put(
			html.Li().Put(html.Text().Set("girafe")),
			html.Li().Put(html.Text().Set("elephant")),
			html.Li().Put(html.Text().Set("tortue")),

			html.Template().Put(
				html.Li().Put(html.Text().Set("poisson")),
				html.Li().Put(html.Text().Set("requin")),
				html.Li().Put(
					html.Text().Set(func() string {
						if rand.Intn(2) == 0 {
							return "huitre"
						}
						return "chevre"
					}()),
				),
			),
		),
	).Stringify()
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, err := fmt.Fprint(w, gen())
		if err != nil {
			log.Fatal(err)
		}
	})

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}
