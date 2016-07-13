package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/tiaguinho/test-meli-go/examples"
	"net/http"
)

var (
	m *martini.ClassicMartini
)

//init
func init() {
	//new server
	m = martini.Classic()

	m.Use(martini.Static("assets"))
	m.Use(martini.Recovery())
	m.Use(martini.Logger())

	m.Use(render.Renderer(render.Options{
		Directory:  "templates",
		Layout:     "layout",
		Extensions: []string{".tmpl", ".html"},
		Delims:     render.Delims{"{[{", "}]}"},
		Charset:    "UTF-8",
		IndentJSON: true,
		IndentXML:  true,
	}))

	m.Action(getRouter())
}

//set routers
func getRouter() martini.Handler {
	r := martini.NewRouter()

	r.Get("/", examples.Home)

	r.Get("/methods/get", examples.Get)
	r.Get("/methods/post", examples.Post)
	r.Get("/methods/put", examples.Put)
	r.Get("/methods/delete", examples.Delete)

	return r.Handle
}

func main() {
	http.ListenAndServe(":8080", m)
}
