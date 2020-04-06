package arcade

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountMinimum(t *testing.T) {
	testCases := map[string]struct {
		schedules      [][]float64
		expectedResult float64
	}{
		"none": {
			schedules:      [][]float64{},
			expectedResult: 0,
		},
		"should": {
			schedules:      [][]float64{{1, 3}},
			expectedResult: 1,
		},
		"ShouldGetOneWithTwoData": {
			schedules:      [][]float64{{1, 2}, {3, 4}},
			expectedResult: 1,
		},
		"ShouldGetTwoWithThreeData": {
			schedules:      [][]float64{{1, 2}, {2, 7}, {3, 8}},
			expectedResult: 2,
		},
		"ShouldGet4WithFourSchedule": {
			schedules:      [][]float64{{1, 8}, {2, 5}, {5, 6}, {3, 7}},
			expectedResult: 4,
		},
	}

	for v, test := range testCases {
		t.Run(v, func(t *testing.T) {
			s := NewSchedule(test.schedules)
			res := s.CountMinimum()
			assert.Equal(t, test.expectedResult, res)
		})
	}
}
