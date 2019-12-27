package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSanitizeToAlphaNumeric(t *testing.T) {
	cases := map[string]struct {
		source string
		result string
	}{
		"With symbols": {
			source: "***Lorem Ipsum",
			result: "Lorem Ipsum",
		},
		"With numbers": {
			source: "Lor3m 1psum",
			result: "Lor3m 1psum",
		},
		"Middle Simbols": {
			source: "Lorem !psum",
			result: "Lorem psum",
		},
		"Non Symbol": {
			source: "Lorem Ipsum",
			result: "Lorem Ipsum",
		},
	}

	for v, test := range cases {
		t.Run(v, func(t *testing.T) {
			res := SanitizeToAlphaNumeric(test.source)
			assert.Equal(t, test.result, res)
		})
	}
}
