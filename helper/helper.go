package helper

import (
	"fmt"
	"strconv"
	"strings"
)

// ParseAndValidateInput is used for parse and validate input
func ParseAndValidateInput(s []string) ([][]float64, error) {
	var input [][]float64

	for _, v := range s {
		// initiate slice variable to store it as set of start time and end time
		var i []float64

		// split the text into slice
		s := strings.Split(v, " ")

		// throw error if given time schedule is not 2 timestamps
		if len(s) != 2 {
			return input, fmt.Errorf("Invalid time schedule")
		}

		// convert first timestamp (start time) index to float64 type
		startPlay, err := strconv.ParseFloat(s[0], 64)
		if err != nil {
			return input, fmt.Errorf("Start play time is in not correct format")
		}

		// show error if timestamp is not in 24 hour range
		if startPlay < 0 || startPlay > 2359 {
			return input, fmt.Errorf("Start play has to be within 0000 to 2359")
		}

		// convert second timestamp (end time) index to float64 type
		endPlay, err := strconv.ParseFloat(s[1], 64)
		if err != nil {
			return input, fmt.Errorf("End play time is in not correct format")
		}

		// show error if timestamp is not in 24 hour range
		if endPlay < 0 || endPlay > 2359 {
			return input, fmt.Errorf("End play has to be within 0000 to 2359")
		}

		i = append(i, startPlay, endPlay)
		input = append(input, i)
	}

	return input, nil
}
