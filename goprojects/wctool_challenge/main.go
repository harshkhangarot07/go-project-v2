package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	argswithprog := os.Args

	argswithoutprog := os.Args[1:]

	//to target a specific argument
	arg := os.Args[2]
	fmt.Println(arg, "\n ", "\n", argswithoutprog, argswithprog)

	wordptr := flag.String("word", "foo", "a string")

	numptr := flag.Int("number", 42, "an integer")

	forkptr := flag.Bool("fork", false, "a boolean value")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string variable")

	flag.Parse()

	fmt.Println("word:", *wordptr)
	fmt.Println("numb:", *numptr)
	fmt.Println("fork:", *forkptr)
}
