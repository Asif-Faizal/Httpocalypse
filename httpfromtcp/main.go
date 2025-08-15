package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("message.txt")
	if err != nil {
		log.Fatal(err)
	}
	for {
		buf := make([]byte, 8)
		n, err := f.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		if n == 0 {
			break
		}
		fmt.Printf("read: %s\n", string(buf[:n]))
	}
}
