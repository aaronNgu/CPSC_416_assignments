package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Post struct {
	Post string
}

type TemplateInput struct {
	CreatePostResponse string
	CharCredit         int
	Posts              []string
}

func templateHandler(w http.ResponseWriter, r *http.Request) {
	//tmplt := template.N("someName")
	//relative path from where command is executed
	tmplt, _ := template.ParseFiles("cmd/server/static/home.html")

	//tmplt, _ = tmplt.Parse("<button>click me </button>Top Student: {{.NoName}} - {{.Name}}!")
	content := TemplateInput{CharCredit: 0}
	tmplt.Execute(w, content)
}

func createPostHandler(w http.ResponseWriter, r *http.Request) {
	tmplt, _ := template.ParseFiles("cmd/server/static/home.html")
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	//fmt.Fprintf(w, "POST request successful")
	name := r.FormValue("post")
	content := TemplateInput{CreatePostResponse: name}

	tmplt.Execute(w, content)
}

func charCreditHandler(w http.ResponseWriter, r *http.Request) {
	tmplt, _ := template.ParseFiles("cmd/server/static/home.html")
	dummyCharCredit := 10
	content := TemplateInput{CharCredit: dummyCharCredit}
	tmplt.Execute(w, content)
}

func getPostsHandler(w http.ResponseWriter, r *http.Request) {
	dummyPosts := []string{"first post", "second post", "third post"}
	content := TemplateInput{Posts: dummyPosts}
	tmplt, _ := template.ParseFiles("cmd/server/static/home.html")
	tmplt.Execute(w, content)
}

func main() {
	//tmpl := template.Must(template.ParseFiles("./cmd/server/home.html"))
	// relative to runnable
	//http.Handle("/", http.FileServer(http.Dir("cmd/server/static/home.html")))
	http.HandleFunc("/", templateHandler)
	//http.HandleFunc("/form", formHandler)

	http.HandleFunc("/getcharcredit", charCreditHandler)
	http.HandleFunc("/createpost", createPostHandler)
	http.HandleFunc("/getposts", getPostsHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
