// Find the top K most common words in a text document.
// Input path: location of the document, K top words
// Output: Slice of top K words
// For this excercise, word is defined as characters separated by a whitespace

// Note: You should use `checkError` to handle potential errors.

package textproc

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func topWords(path string, K int) []WordCount {
	//declaring structure to be used
	var wc []WordCount
	var index int = 0

	//reads file from computer and outputs data into string array that is is split up by each individual word
	myString, err := os.ReadFile(path)
	checkError(err)
	out := strings.FieldsFunc(string(myString), Split)

	//created a map with a string in boolean in order to check if a word in the string array was already counted
	visited := make(map[string]bool, 0)

	//made the structure the size of the array to ensure there was no out of range exceptions
	wc = make([]WordCount, len(out))

	//for loop that runs through the string array holding each individual word from the file
	for i := 0; i < len(out); i++ {
		//boolean checks to see if the word at a specifix index in the array was already counted or not
		if !visited[out[i]] {
			//if the word has not been counted then the word and the number of times it appears are stored in structure variable
			wc[index] = WordCount{Word: out[i], Count: findFrequency(out, out[i])}

			//index is increased for the struct variable and map is set to true for the word so it is not repeated
			index++
			visited[out[i]] = true
		}
	}

	//calls helper function to sort the struct
	sortWordCounts(wc)

	//creates another struct variable with the first K index
	wcOut := wc[:K]

	return wcOut //returns the struct variable
}

func findFrequency(arr []string, str string) int {
	//function used to find the number of times the word in the file
	//was repeated. Done by simply interating through array and comparing
	//the specific word. If found in the array the counter increases.

	count := 0
	for _, item := range arr {
		if item == str {
			count++
		}
	}

	return count
}

func Split(r rune) bool {
	//function used in order to split the string file at mulitple characters instead of just one.

	return r == ' ' || r == '\n'
}

//--------------- DO NOT MODIFY----------------!

// A struct that represents how many times a word is observed in a document
type WordCount struct {
	Word  string
	Count int
}

// Method to convert struct to string format
func (wc WordCount) String() string {
	return fmt.Sprintf("%v: %v", wc.Word, wc.Count)
}

// Helper function to sort a list of word counts in place.
// This sorts by the count in decreasing order, breaking ties using the word.

func sortWordCounts(wordCounts []WordCount) {
	sort.Slice(wordCounts, func(i, j int) bool {
		wc1 := wordCounts[i]
		wc2 := wordCounts[j]
		if wc1.Count == wc2.Count {
			return wc1.Word < wc2.Word
		}
		return wc1.Count > wc2.Count
	})
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
