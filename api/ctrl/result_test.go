package ctrl

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSpeculate(t *testing.T) {
	ass := assert.New(t)

	detect := []string{"b", "a", "d", "e"}
	right := []string{"a", "b", "c", "d"}
	excepted := []string{"b", "a", "d", "c"}

	for i := 0; i < 10; i++ {
		final := Speculate(detect, right)
		ass.Equal(excepted, final)
	}
}

func TestGetOrder(t *testing.T) {
	ass := assert.New(t)

	detect := []string{"b", "a", "d", "c"}
	right := []string{"a", "b", "c", "d"}
	excepted := "2143"

	for i := 0; i < 10; i++ {
		final := GetOrder(detect, right)
		ass.Equal(excepted, final)
	}
}
