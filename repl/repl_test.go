package repl

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		{
			input: " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input: " hi lovely goose",
			expected: []string{"hi", "lovely", "goose"},
		},
		{
			input: " little bean is so fluffy",
			expected: []string{"little", "bean", "is", "so", "fluffy"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice
		// if they don't match, use t.Errorf and continue to next case
		if len(actual) != len(c.expected) {
			t.Errorf("actual length of %d is not expected length of %d", len(actual), len(c.expected))
		}
		for i:= range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("At %d place in slice, %s does not match %s", i, word, expectedWord)
			}
		}
	}
}