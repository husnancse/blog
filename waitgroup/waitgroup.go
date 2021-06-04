// file: waitgroup.go

package main

import (
	"fmt"
	"sync"
)

func readData(wg *sync.WaitGroup) {
	fmt.Println("Membaca beberapa data...")
	wg.Done()

}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go readData(&wg)
	wg.Wait()
}
