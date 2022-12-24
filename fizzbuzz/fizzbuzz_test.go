package fizzbuzz

import "testing"

func TestFizzBuzz(t *testing.T) {

	cases := []struct {
		input int
		want  string
		name  string
	}{
		{name: "input 1 should be 1", input: 1, want: "1"},
		{name: "input 2 should be 2", input: 2, want: "2"},
		{name: "input 3 should be Fizz", input: 3, want: "Fizz"},
		{name: "input 4 should be 4", input: 4, want: "4"},
		{name: "input 5 should be Buzz", input: 5, want: "Buzz"},
		{name: "input 6 should be Fizz", input: 6, want: "Fizz"},
		{name: "input 7 should be 7", input: 7, want: "7"},
		{name: "input 8 should be 8", input: 8, want: "8"},
		{name: "input 9 should be Fizz", input: 9, want: "Fizz"},
		{name: "input 10 should be Buzz", input: 10, want: "Buzz"},
		{name: "input 11 should be 11", input: 11, want: "11"},
		{name: "input 12 should be Fizz", input: 12, want: "Fizz"},
		{name: "input 13 should be 13", input: 13, want: "13"},
		{name: "input 14 should be 14", input: 14, want: "14"},
		{name: "input 15 should be FizzBuzz", input: 15, want: "FizzBuzz"},
		{name: "input 30 should be FizzBuzz", input: 30, want: "FizzBuzz"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := FizzBuzz(tc.input)

			if got != tc.want {
				t.Errorf("FizzBuzz(%d) = %q, want %q", tc.input, got, tc.want)
			}
		})
	}
}
