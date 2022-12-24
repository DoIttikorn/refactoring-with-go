package fizzbuzz

import "fmt"

func FizzBuzz(num int) string {

	if num%15 == 0 {
		return "FizzBuzz"
	}

	if isFizz(num) {
		return "Fizz"
	}

	if num%5 == 0 {
		return "Buzz"
	}

	return fmt.Sprintf("%d", num)
}

func isFizz(num int) bool {
	return num%3 == 0
}
