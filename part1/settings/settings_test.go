package settings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckRightUsage(t *testing.T) {
	assert.NotNil(t, checkRightUsage(Options{WriteRepeatedLines: true, CountNumOfRepeats: true, OnlyUnicLines: true, IgnoreNumOfChars: 3}))
}
func TestRightUsage(t *testing.T) {
	assert.Equal(t, rightUsage(Options{WriteRepeatedLines: true, CountNumOfRepeats: true, OnlyUnicLines: true, IgnoreNumOfChars: 3}, []bool{true, true, true}), []string{"go run main.go -d -s 3", "go run main.go -c -s 3", "go run main.go -u -s 3"}, "must be equal")
	assert.Equal(t, rightUsage(Options{WriteRepeatedLines: true, CountNumOfRepeats: true, IgnoreNumOfFields: 6}, []bool{true, true, false}), []string{"go run main.go -d -f 6", "go run main.go -c -f 6"}, "must be equal")
}
func TestCutFieldsAndChars(t *testing.T) {
	assert.Equal(t, cutFieldsAndChars("", 1, 1), "", "must be equal")
	assert.Equal(t, cutFieldsAndChars("a", 3, 0), "", "must be equal")
	assert.Equal(t, cutFieldsAndChars("aaa", 0, 100), "", "must be equal")
}
func TestEqualWithoutCase(t *testing.T) {
	assert.True(t, EqualWithoutCase("hello", "HelLo", 0, 0))
	assert.False(t, EqualWithoutCase("helo", "HelLo", 0, 0))
}
func TestEqualWithCase(t *testing.T) {
	assert.True(t, EqualWithCase("hello", "hello", 1, 0))
	assert.False(t, EqualWithCase("hello", "HelLo", 0, 0))
}
