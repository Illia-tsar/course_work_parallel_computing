package course_work_parallel_computing

import (
	"flag"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var nGoRoutines int
var parallelBuild bool
var printTime bool
var searchWord bool
var word string
var endTime time.Duration
var invIdx InvertedIndex
var doc int

func main() {
	flag.IntVar(&nGoRoutines, "n", 1, "Define the number of GoRoutines")
	flag.BoolVar(&parallelBuild, "pb", true, "Build Index in parallel")
	flag.BoolVar(&printTime, "t", true, "Print Index build time")
	flag.BoolVar(&searchWord, "s", false, "Search word in index or not")
	flag.StringVar(&word, "sw", "", "Word to search")
	flag.IntVar(&doc, "d", 1, "Document titles to show")
	flag.Parse()

	if parallelBuild {
		wg.Add(nGoRoutines)
		startTime := time.Now()
		invIdx = ParallelBuild(nGoRoutines, &wg)
		endTime = time.Since(startTime)
		wg.Wait()
	} else {
		startTime := time.Now()
		invIdx = SequentialBuild()
		endTime = time.Since(startTime)
	}
	if printTime {
		fmt.Printf("Index build time: %s\n", endTime)
	}
	if searchWord {
		Find(invIdx, word)
	}
}
