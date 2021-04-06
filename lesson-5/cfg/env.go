package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	line, ok := os.LookupEnv("LINE")
	if !ok {
		log.Fatal("LINE env isn't set")
	}
	fmt.Println(line)
}
