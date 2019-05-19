package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/russross/blackfriday.v2"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := httprouter.New()
	r.NotFound = http.FileServer(http.Dir("public"))
	r.GET("/markdown", GenerateMarkdown)

	r.GET("/posts", PostsIndexHandler)
	r.POST("/posts", PostsCreateHandler)

	http.ListenAndServe(":"+port, r)
}

// GenerateMarkdown creates the markdown
func GenerateMarkdown(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	markdown := blackfriday.Run([]byte(r.FormValue("body")))
	rw.Write(markdown)
}

// PostsIndexHandler prints a simple message
func PostsIndexHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "posts index")
}

// PostsCreateHandler prints a simple message
func PostsCreateHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "posts create")
}
