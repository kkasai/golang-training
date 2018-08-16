package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

type Comic struct {
	Num        int
	Year       string
	Month      string
	Day        string
	Title      string
	Link       string
	News       string
	SafeTitle  string `json:"safe_title"`
	Transcript string
	Alt        string
	Img        string
}

const xkcdURL = "https://xkcd.com/"
const infoJson = "info.0.json"

func main() {
	//===================indexing start===================================
	indexdir := "index"
	if _, err := os.Stat(indexdir); os.IsNotExist(err) {
		if err := os.Mkdir(indexdir, 0777); err != nil {
			log.Fatal(err)
		}
	}

	resp, err := http.Get(xkcdURL + infoJson)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("%s", resp.Status)
		os.Exit(1)
	}
	var comicInfo struct {
		Num int
	}
	jsonDecoder := json.NewDecoder(resp.Body)
	if err := jsonDecoder.Decode(&comicInfo); err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

	for i := 1; i <= comicInfo.Num; i++ {
		path := indexdir + "/" + strconv.Itoa(i)
		if _, err := os.Stat(path); os.IsNotExist(err) {
			f, err := os.Create(path)
			if err != nil {
				fmt.Printf("%s", err)
				os.Exit(1)
			}
			url := xkcdURL + "/" + strconv.Itoa(i) + "/" + infoJson
			resp, err := http.Get(url)
			if err != nil {
				fmt.Printf("%s", err)
				os.Exit(1)
			}
			if resp.StatusCode != http.StatusOK {
				fmt.Printf("%s", resp.Status)
				os.Exit(1)
			}
			bytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Printf("%s", err)
				os.Exit(1)
			}
			if _, err = f.Write(bytes); err != nil {
				resp.Body.Close()
				f.Close()
				fmt.Printf("%s", err)
			}
			resp.Body.Close()
			f.Close()
		}
	}
	//===================indexing end===================================

	//============================
	if len(os.Args) > 1 {
		keyword := os.Args[1]
		for i := 1; i < comicInfo.Num; i++ {
			indexFile := indexdir + "/" + strconv.Itoa(i)
			f, err := os.Open(indexFile)
			if err != nil {
				fmt.Printf("%s", err)
				os.Exit(1)
			}
			var comic Comic
			json.NewDecoder(f).Decode(&comic)
			if match, _ := regexp.MatchString(keyword, comic.Transcript); match {
				fmt.Println("\n==== Matched =====")
				fmt.Printf("URL: %s\n", xkcdURL+strconv.Itoa(i))
				fmt.Printf("Transcript: %s\n", comic.Transcript)
			}
		}
	}
	//============================
}
