package request

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequestLineParser(t *testing.T) {
	assert.Equal(t, "this is my test bro", "this is my test bro")
}
