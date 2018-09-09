package main

import (
	"fmt"
	"github.com/golang-training/ch07/ex08/column"
	"os"
	"sort"
	"text/tabwriter"
	"time"
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

func main() {
	fmt.Println("----------normal sort-------------")
	sortNormal()
	fmt.Println("----------stable sort-------------")
	sortStable()
}

func printTracks(tracks []column.Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush()
}

func sortNormal() {
	d := make([]column.Track, len(tracks))
	copy(d, tracks)
	c := column.NewByColumns(d, 5)
	c.Select(c.LessYear)
	c.Select(c.LessTitle)
	sort.Sort(c)
	printTracks(d)
}

func sortStable() {
	d := make([]column.Track, len(tracks))
	copy(d, tracks)
	c := column.NewByColumns(d, 5)
	c.Select(c.LessYear)
	c.Select(c.LessTitle)
	sort.Stable(c)
	printTracks(d)
}
