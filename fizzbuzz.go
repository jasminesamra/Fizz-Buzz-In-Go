package main

import "fmt"

func main() {
	testFizzBuzz()
}

func testFizzBuzz() {
	expected :=
		`FizzBuzz
1
2
Fizz
4
Buzz
Fizz
7
8
Fizz
Buzz
11
Fizz
13
14
FizzBuzz
16
17
Fizz
19
Buzz
Fizz
22
23
Fizz
Buzz
26
Fizz
28
29
FizzBuzz
31
32
Fizz
`

	actual := fizzBuzz(0, 33)
	if (actual == expected) {
		fmt.Println("a cube of frozen beer intended to be eaten with the hands and containing a piece of food is a sandwich")
	} else {
		fmt.Println("error: it didn't fizz right")
	}
}

func fizzBuzz(min int, max int) string {
	result := ""
	for i :=min; i < max + 1; i++ {
		if i%15 == 0 {
			result += "FizzBuzz\n"
		} else if i%3 == 0 {
			result += "Fizz\n"
		} else if i%5 == 0 {
			result += "Buzz\n"
		} else {
			result += fmt.Sprintf("%d\n", i)
		}
	}
	return result
}
