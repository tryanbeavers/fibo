package main

import "fmt"
import "os"
import "strconv"

func main() {
	iter, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(iter, " : ITERATIONS")

	iterRes := fiboIter(iter)
	fmt.Println(iterRes, " : ITERATIVE COMPUTED RESULT")

	recurRes := fiboRecur(iter)
	fmt.Println(recurRes, " : RECURSIVE COMPUTED RESULT")
}

func fiboIter(iterations int) int {
	prev := 0
	curr := 1

	for i := 1; i <= iterations; i++ {
		swap := prev
		prev = curr
		curr = swap + prev
	}

	return prev
}

func fiboRecur(iterations int) int {
	if iterations == 0 {
		return 0
	} else if iterations == 1 {
		return 1
	} else {
		return fiboRecur(iterations-2) + fiboRecur(iterations-1)
	}

}
