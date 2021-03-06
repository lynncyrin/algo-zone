package algozone

import (
	"testing"
)

///////////////////////
// sort script start //
///////////////////////

// selection sort!
//
// docs: https://en.wikipedia.org/wiki/Selection_sort
//
// Selection sort looks through every element an input list, and finds the
// smallest element. That element is then appended to the end of an output list.
// When it reaches the end of the input list, all of the output elements will
// be sorted.

func selectionSort(inputList []string) (sortedList []string) {
	unsortedList := inputList
	inputListLength := len(inputList)

	for i := 0; i < inputListLength; i++ {
		smallestIndex := findSmallestIndex(unsortedList)
		smallestElement := unsortedList[smallestIndex]
		sortedList = append(sortedList, smallestElement)
		unsortedList = removeIndex(unsortedList, smallestIndex)
	}

	return sortedList
}

func findSmallestIndex(inputList []string) (smallestIndex int) {
	for index, element := range inputList {
		if element < inputList[smallestIndex] {
			smallestIndex = index
		}
	}

	return smallestIndex
}

func removeIndex(inputList []string, index int) (outputList []string) {
	outputList = append(inputList[:index], inputList[index+1:]...)
	return outputList
}

/////////////////////
// sort script end //
/////////////////////

func TestSelectionSort(t *testing.T) {
	testData := []struct {
		name     string
		sortFunc sortFunc
	}{
		{
			name:     "sorted_by_go_sort_selection_sort_test",
			sortFunc: selectionSort,
		},
	}
	for _, test := range testData {
		t.Run(test.name, func(t *testing.T) {
			// setup
			inputList, err := getInputList()
			if err != nil {
				t.Error(err.Error())
			}

			// logic under test
			outputList := test.sortFunc(inputList)

			// assertions
			err = writeAndCompareOutputList(outputList, test.name)
			if err != nil {
				t.Error(err.Error())
			}
		})
	}
}

// ☝🏽 per-script helpers
