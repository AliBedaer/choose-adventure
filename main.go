package main

import (
	"cyoaweb/cyoa"
	"flag"
	"fmt"
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

	fmt.Println(story)
}
