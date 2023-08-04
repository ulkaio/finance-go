package main

import (
	"fmt"
	"github.com/ulkaio/finance-go/equity"
)

func main() {
	q, err := equity.Get("AMZN")
	if err != nil {
		// Uh-oh.
		panic(err)
	}
	// Print out the quote
	fmt.Printf("%v\n", q)
}
