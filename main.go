package main

import (
	"fmt"
	"os"

	gojsondiff "github.com/fernandoporazzi/go-json-diff/go-json-diff"
)

func main() {
	fileA, err := os.Open("a.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully opened a.json")
	defer fileA.Close()

	fileB, err := os.Open("b.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully opened b.json")
	defer fileB.Close()

	// It's show time!
	comparator := gojsondiff.NewComparator(fileA, fileB)

	fmt.Println("Are the files equal?", comparator.Compare())
}
