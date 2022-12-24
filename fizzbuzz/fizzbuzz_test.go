package fizzbuzz

import "testing"

func TestFizzBuzz(t *testing.T) {

	cases := []struct {
		input int
		want  string
		name  string
	}{
		{name: "input 1 should return 1", input: 1, want: "1"},
		{name: "input 2 should return 2", input: 2, want: "2"},
		{name: "input 3 should return Fizz", input: 3, want: "Fizz"},
		{name: "input 4 should return 4", input: 4, want: "4"},
		{name: "input 5 should return Buzz", input: 5, want: "Buzz"},
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
