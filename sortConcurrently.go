package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// This assignment makes you make your own partitions instead of using the classic
// merge sort algorithm for the recursive divide and conquer strategy
func merge(a []int, b []int) []int {
	output := []int{}
	aIndex := 0
	bIndex := 0

	//compare elements in each sorted array to append to output in order for merge
	for aIndex < len(a) && bIndex < len(b) {

		if a[aIndex] < b[bIndex] {
			output = append(output, a[aIndex])
			aIndex++
		} else {
			output = append(output, b[bIndex])
			bIndex++
		}
	}

	// need to add what's remaining for a or b array to output because
	// if you already appended all the elements from on array you can add the rest
	// and it will be in sorted order
	for aIndex < len(a) {
		output = append(output, a[aIndex])
		aIndex++

	}

	for bIndex < len(b) {
		output = append(output, b[bIndex])
		bIndex++
	}

	return output

}

// / function decorator that allows you to dynamically partion a slice base on how many chunks you want
func PartionArray(chunks int) func([]int, chan []int) {

	return func(numArray []int, val chan []int) {

		/// anonymous go routine that sends a piece of the slice sorted to then get handled
		/// by the merge function from the receiving channel in main
		go func() {
			defer close(val)
			for i := 0; i < chunks; i++ {

				min := (i * len(numArray) / chunks)
				max := (i + 1) * len(numArray) / chunks

				temp := numArray[min:max]
				sort.Ints(temp)

				val <- temp

			}

		}()

	}

}

func main() {
	intSlice := []int{}
	var numArray []int
	var output []int
	chunkChannel := make(chan []int)

	fmt.Println("Please enter integers. Enter X to close program")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		fmt.Println(err)

	}

	for _, str := range strings.Fields(scanner.Text()) {
		if num, err := strconv.Atoi(str); err != nil {
			fmt.Println(err)
		} else {

			numArray = append(numArray, num)
		}

	}

	intSlice = append(intSlice, numArray...)

	chunks := 4
	splitArray := PartionArray(chunks)
	splitArray(intSlice, chunkChannel)

	for i := range chunkChannel {

		//this is the part in main after the go routine sends a piece of the sorted slice to the channel
		// that then gets merged
		output = merge(output, i)
	}

	fmt.Println("Final Output ", output)

}
