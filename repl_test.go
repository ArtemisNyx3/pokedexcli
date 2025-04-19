package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "pikachu Charmander BULbasUr",
			expected: []string{"pikachu", "charmander", "bulbasur"},
		},
		{
			input:    "heLlo WORLD",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice against the expected slice
		if len(c.expected) > len(actual){
			t.Errorf("Less words than expected")
			t.Fail()
		}else if len(c.expected) < len(actual) {
			t.Errorf("More words than expected")
			t.Fail()
		}
		// if they don't match, use t.Errorf to print an error message
		// and fail the test

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
			if word != expectedWord {
				t.Errorf("Incorrect word")
				t.Fail()
			}
		}
	}
	

}
