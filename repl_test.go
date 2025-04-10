package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input: "Pikachu Bulbasaur Charmander",
			expected: []string{"pikachu", "bulbasaur", "charmander"},
		},
		{
			input: "Gotta catch em all",
			expected: []string{"gotta", "catch", "em", "all"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected){
			t.Errorf("cleanInput(%q) returned %d words, expected %d words", c.input, len(actual), len(c.expected))
    continue
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("cleanInput(%q)[%d] = %q, expected %q", c.input, i, actual[i], c.expected[i])
			}
		}
	}
}