package main

import (
	"html/template"
	"log"
	"net/http"
	"sort"
	"time"

	"github.com/golang-training/ch07/ex08/column"
)

var tracks = []column.Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

var html = template.Must(template.New("tracks").Parse(`
<html>
<body>
<table>
	<tr>
		<th><a href="?sort=title">Title</a></th>
		<th><a href="?sort=artist">Artist</a></th>
		<th><a href="?sort=album">Album</a></th>
		<th><a href="?sort=year">Year</a></th>
		<th><a href="?sort=length">Length</a></th>
	</tr>
{{range .}}
	<tr>
		<td>{{.Title}}</td>
		<td>{{.Artist}}</td>
		<td>{{.Album}}</td>
		<td>{{.Year}}</td>
		<td>{{.Length}}</td>
	</td>
{{end}}
</body>
</html>
`))

func main() {
	c := column.NewByColumns(tracks, 5)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.FormValue("sort") {
		case "title":
			c.Select(c.LessTitle)
		case "artist":
			c.Select(c.LessArtist)
		case "album":
			c.Select(c.LessAlbum)
		case "year":
			c.Select(c.LessYear)
		case "length":
			c.Select(c.LessLength)
		}
		sort.Sort(c)
		err := html.Execute(w, tracks)
		if err != nil {
			log.Printf("template error: %s", err)
		}
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}
