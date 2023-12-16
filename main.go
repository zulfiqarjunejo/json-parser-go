package main

import (
	"flag"
	"fmt"
	"jp/jp"
	"os"
)

func main() {
	file_name := flag.String("file", "", "Name of the json file")
	flag.Parse()

	if *file_name == "" {
		fmt.Printf("Error: file name cannot be blank\n")
		os.Exit(1)
	}

	fmt.Println("parsing file", *file_name)

	if err := jp.Parse("{}"); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	os.Exit(0)
}
