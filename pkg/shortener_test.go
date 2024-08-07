package pkg

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	id        uint64
	shortened string
}{
	{
		id:        12345,
		shortened: "dnh",
	},
	{
		id:        30540,
		shortened: "h6K",
	},
	{
		id:        0,
		shortened: "a",
	},
}

func TestShortener(t *testing.T) {
	for _, test := range testCases {

		expected := id2String(test.id)
		log.Println(expected)
		assert.Equal(t, expected, test.shortened)
	}
}

func TestReverseShortener(t *testing.T) {
	var expected uint64
	for _, test := range testCases {
		expected = string2Id(test.shortened)
		assert.Equal(t, test.id, expected)
	}
}
