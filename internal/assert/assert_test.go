package assert

import (
	"testing"
)

func TestAssert(t *testing.T) {
	actual := "test"
	want := actual

	Equal(t, actual, want)
}
