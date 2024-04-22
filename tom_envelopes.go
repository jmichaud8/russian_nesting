package main

/*
	Envelopes

	In this coding challenge, you are tasked with solving the "Russian Doll Envelopes" problem.
	Given a 2D array of envelopes, where envelopes = [wi, hi] represent the width and height of an envelope,
	your goal is to determine the maximum number of envelopes you can nest inside each other.
	An envelope can fit into another only if both its width and height are strictly greater than
	those of the envelope it fits into; you cannot rotate the envelopes.

	Method:
	- I would like to only read through the input array once
	- Therefore, need to track all the envelopes in each valid nesting list that is being built according to the nesting rules
	- Track each envelope in the list and also track the length of each list
	- Parse stdin into an array of Envelope structs.  There are up to N envelopes.
	- As the input is read and parsed.  For each new envelope:
		a. Go through each list that exists so far.  For each list check for these cases:
			1. The new envelope fits in the first envelope in the list:  add the new envelope to the front of the list
			2. The last element in the list fits within the new envelope:  add the new envelope to the end of the list
			3. The new envelope fits into the interior of the list: add the new envelope to the middle of the list
		b. Create a new list with the new envelope
	- Find the max list length
	- Print result
*/

import (
	"fmt"
)

type Envelope struct {
	w int
	h int
}

func main() {
	var max int           // max list size
	var numLists int      // current number of lists
	var length []int      // size of each list
	var list [][]Envelope // the list of envelopes themselves
	var newEnvelope Envelope
	var done bool = false
	var i int
	var j, k int

	// Read stdin and parse, looking for an envelope.  Exact format example: [[5,4],[6,4],[6,7],[2,3]]
	fmt.Println("Input envelope string then <enter>: ")
	fmt.Scanf("[")

	for !done {
		// Read an envelope
		fmt.Scanf("[%d,%d]", &newEnvelope.w, &newEnvelope.h)
		//fmt.Println()            // Debug
		//fmt.Println(newEnvelope) // Debug

		// Go through each list that exists so far.  For each list check for the three cases
		for i = 0; i < numLists; i++ {
			// Look for one of the three cases

			// Case 1. The new envelope fits in the first envelope in the list?
			if newEnvelope.w < list[i][0].w && newEnvelope.h < list[i][0].h {
				// Add the new envelope to the front of the list
				// Append column to the slice
				list[i] = append(list[i], Envelope{0, 0})
				for j = length[i]; j > 0; j-- {
					list[i][j] = list[i][j-1]
				}
				list[i][0] = newEnvelope
				length[i]++
			}

			// Case 2. The last element in the list fits within the new envelope?
			if list[i][length[i]-1].w < newEnvelope.w && list[i][length[i]-1].h < newEnvelope.h {
				// Add the new envelope to the end of the list
				// Append column to the slice
				list[i] = append(list[i], Envelope{0, 0})
				list[i][length[i]] = newEnvelope
				length[i]++
			}

			// Case 3. The new envelope fits into the interior of the list
			for j = 0; j < length[i]-1; j++ {
				// Criteria.  New envelope able to fit current element in it and fit into next element
				if newEnvelope.w > list[i][j].w && newEnvelope.h > list[i][j].h && newEnvelope.w < list[i][j+1].w && newEnvelope.h < list[i][j+1].h {
					// Add the new envelope to the middle of the list
					// Append column to the slice
					list[i] = append(list[i], Envelope{0, 0})
					for k = length[i]; k > j+1; k-- {
						list[i][k] = list[i][k-1]
					}
					list[i][j+1] = newEnvelope
					length[i]++
				}
			}

		} // end of looping through existing lists

		// Create a new list with the new envelope (always do this)
		// Append rows to the slices
		list = append(list, []Envelope{newEnvelope})
		length = append(length, 1)
		numLists++

		//fmt.Println(list)   // Debug
		//fmt.Println(length) // Debug

		// Read the comma to get to next element.
		// If there is an error that means it is the end of the list
		n, err := fmt.Scanf(",")
		if err != nil {
			done = true
			if n > 0 {
				// do nothing
			}
		}
	}
	// Find the max list length
	for i = 0; i < numLists; i++ {
		if length[i] > max {
			max = length[i]
		}
	}

	// Print result
	fmt.Println(max)
}
