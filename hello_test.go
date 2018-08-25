package hello

import "testing"

func TestHi(t *testing.T) {
	expected := "Hello, world"
	actual := SayHi()
	if expected != actual {
		t.Errorf("Hi was incorrect, got: %s, want: %s.", actual, expected)
	}
}
