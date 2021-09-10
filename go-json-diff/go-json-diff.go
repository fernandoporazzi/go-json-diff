package gojsondiff

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
)

type Comparator struct {
	fileA *os.File
	fileB *os.File
}

// NewComparator takes in two files as arguments and returns a Comparator struct
func NewComparator(a, b *os.File) *Comparator {
	return &Comparator{
		fileA: a,
		fileB: b,
	}
}

// Compare checks if fileA and fileB have the same content
func (c *Comparator) Compare() bool {
	// Check if files are the same size
	// This is the cheapest check possible, thus the first
	if getFileSize(c.fileA) != getFileSize(c.fileB) {
		fmt.Println("Files have different size")
		return false
	}

	/**
	* Convert bytes to an interface{} we can deal with.
	*
	* Since the shape of our JSON inputs is unknown, we need
	* to use the interface{} type to fully be able to loop
	* through deeply nested objects.
	*
	* Recursion is needed to loop through nested objects.
	* This is yet to be implemented.
	 */
	ia := getInterface(c.fileA)
	ib := getInterface(c.fileB)

	return isEqual(ia, ib)

}

func getFileSize(f *os.File) int64 {
	stat, err := f.Stat()
	if err != nil {
		fmt.Println("Error getting Stat for JSON A")
		fmt.Println(err)
	}

	return stat.Size()
}

func getInterface(f *os.File) interface{} {
	byteValue, _ := ioutil.ReadAll(f)

	var result interface{}

	err := json.Unmarshal([]byte(byteValue), &result)
	if err != nil {
		fmt.Println("Error unmarshalling the JSON A")
		fmt.Println(err)
	}

	return result
}

func isEqual(a, b interface{}) bool {
	/**
	* Get type of first item in the JSON.
	*
	* Everything stems from this initial type.
	 */
	kindA := reflect.TypeOf(a).Kind()

	switch kindA {
	case reflect.Slice:
		sliceA := a.([]interface{})
		sliceB := b.([]interface{})

		equal := true

		if len(sliceA) != len(sliceB) {
			fmt.Println("Length of inner slices is different")
			return false
		}

		for _, v := range sliceA {
			if !contains(sliceB, v) {
				equal = false
				break
			}
		}

		return equal

	case reflect.Map:
		mapA := a.(map[string]interface{})
		mapB := b.(map[string]interface{})

		equal := false

		for _, v := range mapA {
			for _, v2 := range mapB {
				if reflect.DeepEqual(v, v2) {
					return true
				}
			}
		}

		return equal

	}
	return false
}

func contains(slice []interface{}, value interface{}) bool {
	for _, v := range slice {
		if reflect.DeepEqual(v, value) {
			return true
		}
	}

	return false
}
