package strategy_test

import (
	"strategy"
)

func ExampleSorter() {
	exampleSorter := strategy.NewSorter()
	exampleSorter.SortFile("abc.txt", 5*strategy.GB)
	exampleSorter.SortFile("def.txt", 10*strategy.GB)
	exampleSorter.SortFile("def.txt", 101*strategy.GB)
	// output:
	// abc.txt is sorted by QuickSort
	// def.txt is sorted by ExternalSort
	// No available sorting algorithm
}
