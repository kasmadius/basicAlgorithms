package basicAlgorithms

import (
	"log"
	"testing"
)

func TestBinSearch(t *testing.T) {
	arr := []int{1, 10, 20, 47, 59, 63, 75, 88, 99, 107, 120, 133, 155, 162, 176, 188, 199, 200, 210, 222}
	testInts := []int{1, 25, 99, 222, 166, 120}
	for i := 0; i < len(testInts); i++ {
		log.Printf("Value %v at index %v", testInts[i], binSearch(arr, testInts[i]))
	}
}
