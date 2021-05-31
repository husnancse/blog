package main

import (
	"fmt"
	"time"
)

func initApp() {
	fmt.Println("Memulai aplikasi...")
}

func timeOut(sec time.Duration) {
	fmt.Println(int(sec), "detik sebelum akhir batas waktu.")
	time.Sleep(sec * time.Second)
}

func loadHTML() {
	fmt.Println("Memasang tags HTML...")
	time.Sleep(25 * time.Millisecond)
}

func loadImages() {
	types := []string{"jpg", "png", "gif"}

	for _, v := range types {
		go func(t string) {
			fmt.Println("Memasang gambar tipe:", t)
			time.Sleep(100 * time.Millisecond)
		}(v)
	}
}

func loadScripts() {
	scripts := []string{"CSS", "JavaScript"}

	for _, v := range scripts {
		fmt.Println("Memasang script", v)
		time.Sleep(50 * time.Millisecond)
	}
}

func main() {
	initApp()

	go loadImages()
	go loadScripts()
	go loadHTML()

	timeOut(3)
}
