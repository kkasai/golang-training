package main

import (
	"./github"
	"./template"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/search", handleSearch)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "example request: http://localhost:8080/search?q=keyword1+keyword2+...")
}

func handleSearch(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	v := r.Form.Get("q")
	keywords := strings.Split(v, " ")
	result, err := github.SearchIssues(keywords)
	if err != nil {
		fmt.Print("")
	}
	template.OutputHTML(w, result)
}
