package random

import (
	"math/rand"
	"time"
)

// RandInt Depracted: Use Int instead
func RandInt(high int) int {
	return randInt(1, high)
}

// Int Returns a number betweek 1 and the high argument
func Int(high int) int {
	return randInt(1, high+1)
}

// RandIndex Depracted: Use Index instead
func RandIndex(count int) int {
	return randInt(0, count)
}

// Index returns an index between 0 and the count (non-inclusive)
func Index(count int) int {
	return randInt(0, count)
}

// RandIntFromArray Depracted: Use IntFromArray instead
func RandIntFromArray(array []int) int {
	index := randInt(0, len(array))
	return array[index]
}

// IntFromArray returns a random string from the array passed
func IntFromArray(array []int) int {
	index := randInt(0, len(array))
	return array[index]
}

// Array returns a random ordered array of `count` length
func Array(count int) []int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Perm(count)
}

func randInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}
