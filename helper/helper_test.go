package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseAndValidateInput(t *testing.T) {
	testCases := map[string]struct {
		Input   []string
		IsError bool
	}{
		"ShouldSuccess": {
			Input:   []string{"0000 0010"},
			IsError: false,
		},
		"ShouldFailedWithIncorrectTimeFormat": {
			Input:   []string{"4354 3214"},
			IsError: true,
		},
		"ShouldFailedWithIncorrectScheduleLength": {
			Input:   []string{"0010 0020 0030"},
			IsError: true,
		},
	}

	for v, test := range testCases {
		t.Run(v, func(t *testing.T) {
			_, err := ParseAndValidateInput(test.Input)

			if test.IsError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
