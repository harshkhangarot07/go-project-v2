package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	operation := flag.String("op", "", "operation to perform : rename , move")
	source := flag.String("src", " ", "Source file path")
	destination := flag.String("dest", "", "Destination file path")

	flag.Parse()

	if *operation == "" || *source == "" {
		fmt.Println("usage : -op <opeation> -src <source> -dest <destinantion>")
		return
	}

	switch *operation {
	case "rename":
		fmt.Printf("Renaming file from %s to %s\n", *source, *destination)
		err := os.Rename(*source, *destination)
		if err != nil {
			fmt.Printf("Error renaming file: %v\n", err)
			return
		}
		fmt.Printf("File renamed successfully from %s to %s\n", *source, *destination)

	case "move":
		err := os.Rename(*source, *destination)
		if err != nil {
			fmt.Printf("error in moving file %v\n", err)
			return
		}
		fmt.Printf("file moved from %s to %s", *source, *destination)

	default:
		fmt.Printf("unsupported opearation . use 'rename' or 'move' .")
	}
}
