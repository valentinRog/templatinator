package main

import (
	"fmt"
	"log"
	"strconv"

	"net/http"

	"github.com/valentinrog/templatinator"
	"github.com/valentinrog/templatinator/tag"
)

var html templatinator.Factory

var gid uint64 = 0
var todos = map[uint64]string{}

func addTodo(todo string) {
	gid += 1
	todos[gid] = todo
}

func renderTodos() tag.Tag {
	return html.Template().AppendChildren(
		func() []tag.Tag {
			var a []tag.Tag
			for id, todo := range todos {
				e := html.Li().AppendChildren(
					html.Text().Set(todo),
					html.Button().
						SetAttr("hx-post", fmt.Sprintf("/delete?id=%d", id)).
						SetAttr("hx-target", "#todos-list").
						SetAttr("hx-swap", "innerHTML").
						AppendChildren(html.Text().Set("delete")),
				)
				a = append(a, e)
			}
			return a
		}()...,
	)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		content := html.Html().AppendChildren(
			html.Head().AppendChildren(
				html.Title().AppendChildren(html.Text().Set("My todo-list")),
				html.Link().SetAttr("rel", "stylesheet").SetAttr("href", "https://cdn.jsdelivr.net/npm/@picocss/pico@2/css/pico.min.css"),
				html.Script().SetAttr("src", "https://unpkg.com/htmx.org@1.9.11"),
			),
			html.Body().AppendChildren(
				html.Ul().SetAttr("id", "todos-list").AppendChildren(renderTodos()),
				html.Form().
					SetAttr("hx-post", "/add").
					SetAttr("hx-target", "#todos-list").
					SetAttr("hx-swap", "innerHTML").
					AppendChildren(
						html.Input().SetAttr("type", "text").SetAttr("required", "").SetAttr("name", "todo"),
						html.Input().SetAttr("type", "submit").SetAttr("value", "add todo"),
					),
			),
		)

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, err := fmt.Fprint(w, "<!DOCTYPE html>"+content.Stringify())
		if err != nil {
			log.Fatal(err)
		}
	})

	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		addTodo(r.FormValue("todo"))
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, err := fmt.Fprint(w, renderTodos().Stringify())
		if err != nil {
			log.Fatal(err)
		}
	})

	http.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.ParseUint(r.URL.Query().Get("id"), 10, 64)
		delete(todos, id)
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, err := fmt.Fprint(w, renderTodos().Stringify())
		if err != nil {
			log.Fatal(err)
		}
	})

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}
