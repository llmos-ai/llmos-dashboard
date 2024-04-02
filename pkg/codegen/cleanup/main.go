package main

import (
	"log"
	"os"
)

func main() {
	if err := os.RemoveAll("./pkg/generated"); err != nil {
		log.Fatal(err)
	}
}
