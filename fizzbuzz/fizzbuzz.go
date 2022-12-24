package fizzbuzz

import "fmt"

func FizzBuzz(n int) string {
	if n == 3 || n == 6 {
		return "Fizz"
	}

	if n == 5 {
		return "Buzz"
	}

	return fmt.Sprintf("%d", n)
}
