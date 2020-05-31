package filter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FilterNewsTickerItems(t *testing.T) {
	// given
	testdata := map[string]bool{
		"LEIPZIG":       false,
		"2. BUNDESLIGA": false,
		"SPORT":         false,
		"MOTORSPORT":    false,
		"WIRTSCHAFT":    true,
	}

	for category, expected := range testdata {
		// when
		actual := isInterestingCategory(category)

		// then
		assert.Equal(t, expected, actual, category)
	}
}
