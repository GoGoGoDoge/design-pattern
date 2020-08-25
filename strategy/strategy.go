package strategy

import "fmt"

var sortAlgFactory map[string]ISortAlg

func init() {
	sortAlgFactory = make(map[string]ISortAlg)
	sortAlgFactory["QuickSort"] = QuickSort{}
	sortAlgFactory["ExternalSort"] = ExternalSort{}
}

type Sorter struct {
	algs []AlgRange
}

const GB uint64 = 1000 * 1000 * 1000

func NewSorter() Sorter {
	sorter := Sorter{}
	sorter.algs = append(sorter.algs, AlgRange{
		start: 0,
		end:   6 * GB,
		alg:   sortAlgFactory["QuickSort"],
	})
	sorter.algs = append(sorter.algs, AlgRange{
		start: 6 * GB,
		end:   100 * GB,
		alg:   sortAlgFactory["ExternalSort"],
	})
	return sorter
}

func (sorter Sorter) SortFile(filePath string, size uint64) {
	var sortAlg ISortAlg

	// Get rid of if else here
	for _, alg := range sorter.algs {
		if alg.inRange(size) {
			sortAlg = alg.getAlg()
			break
		}
	}

	if sortAlg == nil {
		fmt.Println("No available sorting algorithm")
		return
	}

	sortAlg.sort(filePath)
}

type ISortAlg interface {
	sort(filePath string)
}

type QuickSort struct{}

func (sort QuickSort) sort(filePath string) {
	fmt.Println(filePath + " is sorted by QuickSort")
}

type ExternalSort struct{}

func (sort ExternalSort) sort(filePath string) {
	fmt.Println(filePath + " is sorted by ExternalSort")
}

type AlgRange struct {
	start uint64
	end   uint64
	alg   ISortAlg
}

func (algRange AlgRange) inRange(size uint64) bool {
	return size >= algRange.start && size <= algRange.end
}

func (algRange AlgRange) getAlg() ISortAlg {
	return algRange.alg
}
