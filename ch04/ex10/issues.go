package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"./github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	dayOfOneMonthAgo := time.Now().AddDate(0, -1, 0)
	dayOfOneYearsAgo := time.Now().AddDate(-1, 0, 0)

	var lessMonthResults []*github.Issue
	for _, v := range result.Items {
		if v.CreatedAt.After(dayOfOneMonthAgo) {
			lessMonthResults = append(lessMonthResults, v)
		}
	}
	fmt.Println("--- less than one month ---")
	for _, item := range lessMonthResults {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}

	var lessYearsResults []*github.Issue
	for _, v := range result.Items {
		if v.CreatedAt.After(dayOfOneYearsAgo) {
			lessYearsResults = append(lessYearsResults, v)
		}
	}
	fmt.Println("--- less than one year ---")
	for _, item := range lessYearsResults {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}

	var olderYearsResults []*github.Issue
	for _, v := range result.Items {
		if v.CreatedAt.Before(dayOfOneYearsAgo) {
			olderYearsResults = append(olderYearsResults, v)
		}
	}
	fmt.Println("--- older than one year ---")
	for _, item := range olderYearsResults {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}
