package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Scanner for user input
	inputScanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter text (Ctrl+z , press enter  to stop):")

	for inputScanner.Scan() {
		fmt.Println("You typed:", inputScanner.Text())
	}

	if err := inputScanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}

	// Open the file to be read
	file, err := os.Open("harsh.txt")
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer file.Close()

	// Scanner for file reading
	fileScanner := bufio.NewScanner(file)

	// Read each line from the file
	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())
	}

	if err := fileScanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
