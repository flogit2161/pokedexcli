package main 

import "testing"

func TestCleanInput(t *testing.T){
	cases := []struct{
		input string
		expected []string
	}{
		{
			input:     " hello    world",
			expected: []string{"hello", "world"},
		},
		{
			input:"hello WOrLd",
			expected: []string{"hello", "world"},
		},
		{
			input:"hellO wOrLD",
			expected: []string{"hello", "world"},
		},
		
	}

	for _, c := range cases{
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected){
			t.Errorf("Test Failed, length of slice %d, expected length %d", len(actual), len(c.expected))
			continue
		}
		
		for i := range actual{
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord{
				t.Errorf("Test Failed, word mismatch, actual: %v, expected: %v", word, expectedWord)
			}
		}
	}
}