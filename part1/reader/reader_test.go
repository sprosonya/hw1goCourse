package reader

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFail(t *testing.T) {
	var input = ``
	in := bytes.NewBufferString(input)
	_, err := Reader(in)
	assert.NotNil(t, err)
}
