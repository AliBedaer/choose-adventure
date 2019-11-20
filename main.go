package main

import (
	"cyoaweb/cyoa"
	"flag"
	"log"
	"os"
)

func main() {
	fileName := flag.String("file", "story.json", "file that contain story")
	flag.Parse()

	file, err := os.Open(*fileName)

	if err != nil {
		//panic(err)
		log.Fatal(err)
	}

	story,err := cyoa.ReadJson(file)

	if err != nil {
		panic(err)
	}
	tpl := "templates/stories.html"

	cyoa.ParseTemplate(tpl,story)

}
