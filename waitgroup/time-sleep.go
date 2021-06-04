// file: time-sleep.go

package main

import (
	"fmt"
	"time"
)

func readData() {
	fmt.Println("Membaca beberapa data...")
}

func main() {
	go readData()
	time.Sleep(1 * time.Second)
}
