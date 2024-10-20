package main

import (
	"fmt"
	"log"
	"os"
)

func readdir(dir string) []os.DirEntry {
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("File\n", files)
	return files
}

func main() {
	readdir("./../ch8")
}
