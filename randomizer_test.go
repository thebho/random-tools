package random

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type fn func(int) int

func helper(f fn) []int {
	var testArray = make([]int, 10, 10)
	for i := 0; i < 10; i++ {
		testArray[i] = f(4)
	}
	return testArray
}

func TestInt(t *testing.T) {
	var testArray = helper(Int)
	for _, testCase := range testArray {
		assert.True(t, testCase <= 4)
		assert.True(t, testCase > 0)
	}
}

func TestIndex(t *testing.T) {
	var testArray = helper(Index)
	for _, testCase := range testArray {
		assert.True(t, testCase < 4)
		assert.True(t, testCase >= 0)
	}
}

func TestInFromArray(t *testing.T) {
	var testArray = make([]int, 10, 10)
	for i := 0; i < 10; i++ {
		testArray[i] = IntFromArray([]int{1, 2, 3, 4})
	}
	fmt.Print("Random Int from {1,2,3,4}: { ")
	for _, testCase := range testArray {
		fmt.Printf("%d ", testCase)
		assert.True(t, testCase < 5)
		assert.True(t, testCase > 0)
	}
	fmt.Println("}")
}

func TestArray(t *testing.T) {
	var testArray = make([][]int, 4, 4)
	for i := 0; i < 4; i++ {
		testArray[i] = Array(4)
	}
	fmt.Print("Random Arrays from Array(4): { ")
	for _, testCase := range testArray {
		fmt.Printf("%d ", testCase)
		for _, testCaseElement := range testCase {
			assert.True(t, testCaseElement <= 3)
			assert.True(t, testCaseElement >= 0)
		}
	}
	fmt.Println("}")
}
