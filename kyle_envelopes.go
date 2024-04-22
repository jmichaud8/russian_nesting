package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// Custom struct to hold key:value pairs
type KV struct {
	Key   int
	Value int
}

// Custom struct to represent a map of KV:list of KVs
type NestingMap map[KV][]KV

var nestingMap = make(NestingMap)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	fmt.Println("Please provide input: ")
	message, _ := reader.ReadString('\n')

	// Compile regex pattern (accounting for possible whitespace), using capture groups for the actual values
	pattern := regexp.MustCompile(`\s*\[\s*(?P<key>\d+)\s*,\s*(?P<val>\d+)\s*\]\s*`)

	// Pull key value pairs from input string using regexp
	matches := pattern.FindAllStringSubmatch(message, -1)
	result := make(map[int]int)

	// Stick the kv pairs into a map
	for _, match := range matches {
		// Convert key and val strings to integers
		key, _ := strconv.Atoi(match[1])
		value, _ := strconv.Atoi(match[2])

		// Add to map
		result[key] = value
	}

	// This println and the one on 72 are just for debugging
	//fmt.Println("Map of KV Pairs:", result)

	// Build a structure that contains a map each KV to a slice of all other KVs that would fit inside it
	// ex: [1,1]: {}, [2,2]: {[1,1]}, [3,3]: {[1,1],[2,2]}, etc
	for curr_k, curr_v := range result {
		for next_k, next_v := range result {
			if curr_k > next_k && curr_v > next_v {
				addNesting(nestingMap, KV{curr_k, curr_v}, KV{next_k, next_v})
			}
		}
	}

	var largestKey KV
	var largestSetSize int

	// using the above structure, identify the "largest" KV (ie, the KV with the greatest number of KVs that could fit inside it)
	for pair, set := range nestingMap {
		currentLargestSet := len(set)
		if currentLargestSet > largestSetSize {
			largestKey = pair
			largestSetSize = currentLargestSet
		}
	}

	// use that largest KV to kick off the recursive function
	maxNestingDepth := findMaxNestingDepth(nestingMap[largestKey], 1)

	//fmt.Println("Nesting map: ", nestingMap)
	fmt.Println("Max nesting depth:")
	fmt.Println(maxNestingDepth)

}

// This recursive function will start with the "largest" KV and drill down through KVs, selecting the next "largest" KV
// At each recursion, increment the depth of recursion, since that's all we really need to know when we're done
func findMaxNestingDepth(kvs []KV, nestingDepth int) int {
	if len(kvs) == 0 {
		return nestingDepth
	}

	// Find the KV with the greatest slice length as its value in the current array
	// Full disclosure, this is not completely deterministic. If there is a tie for slice length, it picks the next envelope at random
	// I think the best approach here would be something called a topological search, but I ran out of time :)
	var biggestSlice KV
	var biggestSliceLength int
	for _, kv := range kvs {
		currentSliceLength := len(nestingMap[kv])
		if currentSliceLength > biggestSliceLength {
			biggestSlice = kv
			biggestSliceLength = currentSliceLength
		}
	}

	// Recursively calls itself, passing the nested KV pairs associated with the current largest slice
	return findMaxNestingDepth(nestingMap[biggestSlice], nestingDepth+1)
}

// This helper function is used to build up the map, with conditional logic to see if we need to append or start a new slice
func addNesting(nestingMap map[KV][]KV, outer KV, inner KV) {
	// this is the "comma-ok idiom", something I wasn't familiar with before this challenge
	// kinda like a ternary, or a try-catch
	if existingNested, ok := nestingMap[outer]; ok {
		nestingMap[outer] = append(existingNested, inner)
	} else {
		nestingMap[outer] = []KV{inner}
	}
}
