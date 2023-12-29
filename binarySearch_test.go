package main

import (
	"fmt"
	"testing"
)

// Test binary Search for each number in array of 10
// for a valid return value.
func TestOrderNaturalList(t *testing.T) {
	testArray := []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := int64(0); i < int64(len(testArray)); i++ {
		fmt.Println("find ", i, "in", testArray)
		result := binarySort(testArray, i)
		fmt.Println("results ", result)
		if result != i {
			t.Fatalf(`binarySort failed on %d`, i)
		}
		// t.Error("Found as expected", i)

	}

}

func TestMissingEvenNumber(t *testing.T) {
	testArray := []int64{1, 3, 5, 7, 9}

	for i := int64(0); i < int64(len(testArray)); i++ {
		fmt.Println("find ", i, "in", testArray)
		result := binarySort(testArray, i)
		fmt.Println("results ", result)
		if result != i {
			if i%2 != 0 {
				t.Fatalf(`binarySort failed on %d`, i)
			}
		}
	}
}

func TestMissingOddNumber(t *testing.T) {
	testArray := []int64{0, 2, 4, 6, 8, 10}

	for i := int64(0); i < int64(len(testArray)); i++ {
		fmt.Println("find ", i, "in", testArray)
		result := binarySort(testArray, i)
		fmt.Println("results ", result)
		if result != i {
			if i%2 == 0 {
				t.Fatalf(`binarySort failed on %d`, i)
			}
		}
	}
}
