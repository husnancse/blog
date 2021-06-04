package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	paths := []string{"/etc", "/bin", "/boot", "/dev", "/proc"}
	for _, path := range paths {
		wg.Add(1)
		go readDir(&wg, path)
	}

	wg.Wait()
}

func readDir(wg *sync.WaitGroup, path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Folder \"%s\" berisi %d file/folder\n", path, len(files))
	wg.Done()
}
