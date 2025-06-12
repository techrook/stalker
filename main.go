package main

import (
	"flag"
	"log"

	"github.com/techrook/stalker/watcher"
)

func getPathFromFlags() string {
	path := flag.String("path", ".", "Directory to watch")
	flag.Parse()
	return *path
}

func main() {
	path := getPathFromFlags()
	err := watcher.Watch(path)
	if err != nil {
		log.Fatal(err)
	}
}
