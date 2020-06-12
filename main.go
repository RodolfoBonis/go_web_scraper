package main

import (
	"fmt"
	"os"
	"github.com/anaskhan96/soup"
)

func main() {

	var url string

	if len(os.Args) > 1 {
		url = "https://en.wikipedia.org/wiki/" + os.Args[1]
	} else {
		url = "https://en.wikipedia.org/wiki/Go_(programming_language)"
	}

	resp, err := soup.Get(url)

	if err != nil {
		os.Exit(1)
	}

	doc := soup.HTMLParse(resp)

	sections := doc.FindAll("h2")

	for _, section := range sections {
		child := section.Children()[0]
		if child.Attrs()["class"] == "mw-headline" {
			fmt.Println(section.Children()[0].Text())
		}
	}

}
