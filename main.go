package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	for number := 1; number <= 10000000; number++ {
		if number%3 == 0 {
			fmt.Printf("%v. Fizz\n", number)

		} else if number%5 == 0 {
			fmt.Printf("%v. Buzz\n", number)

		} else if number%3 == 0 && number%5 == 0 {
			fmt.Printf("%v. FizzBuzz\n", number)

		} else {
			fmt.Printf("%v\n", number)
		}
	}

	timeElapsed := time.Since(start)

	fmt.Printf("Total time for execution: %s", timeElapsed)
}
