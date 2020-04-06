package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/fzlrhmn/DefenderArcade/arcade"
	"github.com/fzlrhmn/DefenderArcade/helper"
)

// main function for the program
func main() {
	// Convert the list of time table to slice of string
	s, err := readInputFile()
	if err != nil {
		panic(err)
	}

	// check if the length is over than 100
	// return error if the data is more than 100
	if len(s) > 100 {
		panic("Time schedule cannot more than 100")
	}

	input, err := helper.ParseAndValidateInput(s)
	if err != nil {
		panic(err)
	}

	a := arcade.NewSchedule(input)
	fmt.Println(a.CountMinimum())
}

// readInputFile is used for read the input file
// and return it as slice of string
func readInputFile() ([]string, error) {
	// read input file of time schedule periods
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		// throw error if it is not exist
		return nil, fmt.Errorf("Cannot read input.txt")
	}

	// return as slice of string by newline
	return strings.Split(string(data), "\n"), nil
}
