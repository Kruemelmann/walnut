package main

import (
	"flag"
	"fmt"
)

func main() {
	str := flag.String("entry", "", "name of the entry to generate the repo for")

	//end of all flags
	flag.Parse()

	//TODO interpretate flags and generate corresponding code
	fmt.Printf("hello %v\n", *str)
}
