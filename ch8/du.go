package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func dirents(dir string) []os.DirEntry {
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatalln(err)
	}
	return files
}

func walkDir(dir string, fileSizes chan<- int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, fileSizes)
		} else {
			info, err := entry.Info()
			if err != nil {
				log.Fatalf("Error occurred when reading file size of %s\n Error: %s", entry.Name(), err)
			}
			fileSizes <- info.Size()
		}
	}
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	fileSizes := make(chan int64)
	go func() {
		for _, root := range roots {
			walkDir(root, fileSizes)
		}
		close(fileSizes)
	}()

	/*
		The range loop on the fileSizes channel will not "miss" any values.
		It will keep receiving values as long as the producer (the Go routine) keeps sending them.
		Once the Go routine closes the channel, the loop will exit.
	*/

	var nfiles, nbytes int64
	for size := range fileSizes {
		nfiles++
		nbytes += size
	}
	printDiskUsage(nfiles, nbytes)
}
