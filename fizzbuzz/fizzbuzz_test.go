package fizzbuzz

import "testing"

func TestFizzBuzz(t *testing.T) {
	want := "1"
	input := 1
	got := FizzBuzz(input)

	if got != want {
		t.Errorf("FizzBuzz(%d) = %q, want %q", input, got, want)
	}
}

func TestInput2ShouldReturn2FizzBuzz(t *testing.T) {
	want := "2"
	input := 2

	got := FizzBuzz(input)

	if got != want {
		t.Errorf("FizzBuzz(%d) = %q, want %q", input, got, want)
	}
}

func TestInput3ShouldReturn2FizzBuzz(t *testing.T) {
	want := "Fizz"
	input := 3
	got := FizzBuzz(input)

	if got != want {
		t.Errorf("FizzBuzz(%d) = %q, want %q", input, got, want)
	}
}
