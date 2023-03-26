# TDD in Golang

## References

### Heavily inpired from

https://quii.gitbook.io/learn-go-with-tests/

## Introduction

TDD

## Why?

### Why write tests?
- Ensure the code is doing what it's supposed to
- Catch errors/failures early
- Documentation
	- Tests can act as documentation for your code, since it ensures that you cover a variety of cases. As a reader, you can understand the different use cases for the components/functions.
- Future proof your code
	- As the code base grows, you can't test each workflow which covers every line of code
	- Writing tests ensures that even if you make a change in one component does not affect other components

	
### Why follow TDD?
- Helps us to develop the logic in our code
- Very high test-coverage
- Better structure to our code
- Write testable code instead of restructuring the code afterwards

## How? / Let's learn by scenarios

### Scenario 1 - Implement a Sum() function

Demoes the process of TDD
#### Functionality

Get the sum of an array of numbers

#### Steps in TDD
1. Write unit test
2. Write the minimal amount of code for the test to run
3. (Iteratively) Write enough code to make the test pass

#### TDD for Sum() functioanlity
1. Write unit test

	sum_test.go

	```
	package sum
	
	import "testing"
	
	func TestSum(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		expected := 15
	
		actual := Sum(numbers)
	
		if actual != expected {
			t.Errorf("input: %v; expected: %d; actual: %d\n", numbers, expected, actual)
		}
	}
	
	```
	
	Run the test
	
	```
	$ go test *
	# command-line-arguments [command-line-arguments.test]
	./sum_test.go:9:12: undefined: Sum
	FAIL    command-line-arguments [build failed]
	FAIL
	```

2. Write the minimal amount of code for the test to run

	sum.go
	
	```
	package sum
	
	func Sum(input []int) int {
		return 0
	}
	```	
	
	Run the test
	
	```
	$ go test *
	--- FAIL: TestSum (0.00s)
	    sum_test.go:12: input: [1 2 3 4 5]; expected: 15; actual: 0
	FAIL
	FAIL    command-line-arguments  0.287s
	FAIL
	```

3. Write enough code to make it pass
	
	sum.go
	
	```
	package sum
	
	func Sum(input []int) int {
		sum := 0
		for _, v := range input {
			sum += v
		}
		return sum
	}

	```
	
	Run the test
	
	```
	$ go test -v -run ^TestSum$ ./...
	=== RUN   TestSum
	--- PASS: TestSum (0.00s)
	PASS
	ok      github.com/albingeorge/tdd/1_sum        0.113s
	```
	
### Scenario 2 - Implement a countdown
Shows how TDD can help developers cut down the time it takes to code, but reducing the effort to refactor the code after completing implementation.

#### Functionality

Write a program which counts down from 3, printing each number on a new line (with a 1-second pause) and when it reaches zero it will print "Go!" and exit.

```
3
2
1
Go!
```

#### Let's implement without following TDD

countdown.go

```
package countdown

import (
	"fmt"
	"time"
)

func CountdownNonTdd() {
	num := 3
	for i := num; i > 0; i-- {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}

	fmt.Println("Go!")
}
```

countdown_test.go

```
func TestCountdownNonTdd(t *testing.T) {
	CountdownNonTdd()
}
```

Execute the test

```
$ go test -v -run ^TestCountdownNonTdd$ ./...
=== RUN   TestCountdownNonTdd
3
2
1
Go!--- PASS: TestCountdownNonTdd (3.00s)
PASS
ok      github.com/albingeorge/tdd/2_countdown  3.105s
```

#### Problems with the above implementation

1. We can't test the below functionalities
	1. Is the right data getting printed? (in this case, "3, 2, 1, Go!")
	2. Is the sleep happening in between the entries?
2. The test takes very long time to execute
3. Test prints to stdout polluting the test results

#### TDD approach 1

Use Dependency Injection design pattern to capture the things to test.

1. Write test

	countdown_test.go
	
	```
	func TestCountdown(t *testing.T) {
		// Create a []byte buffer
		buffer := &bytes.Buffer{}

		// Pass the buffer as a dependency to the implementation
		Countdown(buffer)

		// Fetch the content of the bufer to compute the output
		got := buffer.String()
		want := `3
	2
	1
	Go!`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	```
	
	Execute test
	
	```
	$ go test -run ^TestCountdown$ *.go
	# command-line-arguments [command-line-arguments.test]
	./countdown_test.go:14:2: undefined: Countdown
	FAIL    command-line-arguments [build failed]
	```

2. Write the minimal amount of code for the test to run

	countdown.go
	
	```
	func Countdown(writer io.Writer) {
	}
	```
	
	Execute the test
	
	```
	$ go test -run ^TestCountdown$ *.go
	--- FAIL: TestCountdown (0.00s)
	    countdown_test.go:21: got "" want "3\n2\n1\nGo!"
	FAIL
	```
	
3. Write enough code to make it pass

	countdown.go
	
	```
	// Accept the io.Writer interface
	func Countdown(writer io.Writer) {
		num := 3
		for i := num; i > 0; i-- {
			// Replaces the Println with Fprintln, which accepts an io.Writer interface
			fmt.Fprintln(writer, i)
			time.Sleep(1 * time.Second)
		}

		fmt.Fprint(writer, "Go!")
	}
	```
	
	Execute the test

	```
	go test -v -run ^TestCountdown$ *.go
	=== RUN   TestCountdown
	--- PASS: TestCountdown (3.00s)
	PASS
	ok      command-line-arguments  3.123s
	```


##### Usage

main.go

```
func main() {
	countdown.Countdown(os.Stdout)
}
```

#### Output

```
$ go run main.go
3
2
1
Go!
```

##### Advantages of using Dependency Injection in this approach
1. Enables us to test the internals of a function even if the function does not return an output
2. Make the implementation more general purpose. We can now use the Countdown() function for multiple implementations of io.Writer
	1. Get the output as a string - as explained in the test case
	2. Use the function to print to standard output - as mentioned under Usage
	3. Use the function to print to http response writer, etc

#### Problem with Approach 1

- Can't test sleep functionality
- Tests are slow - bottleneck at `time.Sleep()` call

#### TDD approach 2

Use mock to test the sleep functionality

1. Write test

	countdown_test.go
	
	```
	type SleeperMock struct {
		count int
	}
	
	func (m *SleeperMock) Sleep() {
		m.count++
	}
	
	// Can also test how many sleeps are done in the function
	func TestCountdownImproved(t *testing.T) {
		buffer := &bytes.Buffer{}
		sleep := SleeperMock{}
	
		CountdownImproved(buffer, &sleep)
	
		got := buffer.String()
		want := `3
	2
	1
	Go!`
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	
		sleepWanted := 3
		if sleep.count != sleepWanted {
			t.Errorf("sleep count: got %q want %q", sleep.count, sleepWanted)
		}
	}
	```
	
	Execute the test
	
	```
	$ go test -run ^TestCountdownImproved$ *.go
	# command-line-arguments [command-line-arguments.test]
	./countdown_test.go:40:2: undefined: CountdownImproved
	FAIL    command-line-arguments [build failed]
	FAIL
	```

2. 	Write the minimal amount of code for the test to run

	countdown.go
	
	```
	type SleeperInterface interface {
		Sleep()
	}
	
	func CountdownImproved(writer io.Writer, s SleeperInterface) {
	}
	```
	
	countdown_test.go
	
	```
	go test -v -run ^TestCountdownImproved$ *.go
	=== RUN   TestCountdownImproved
	    countdown_test.go:48: got "" want "3\n2\n1\nGo!"
	    countdown_test.go:53: sleep count: got '\x00' want '\x03'
	--- FAIL: TestCountdownImproved (0.00s)
	```
	
3. Write enough code to make the test pass

	countdown.go
	
	```
	type SleeperInterface interface {
		Sleep()
	}
	
	func CountdownImproved(writer io.Writer, s SleeperInterface) {
		num := 3
		for i := num; i > 0; i-- {
			fmt.Fprintln(writer, i)
			s.Sleep()
		}
	
		fmt.Fprint(writer, "Go!")
	}
	```
	
	Run the test
	
	```
	$ go test -v -run ^TestCountdownImproved$ *.go
	=== RUN   TestCountdownImproved
	--- PASS: TestCountdownImproved (0.00s)
	PASS
	```
	
	
##### Usage of countdown()

main.go
```
type sleeperImplementation struct{}

func (s sleeperImplementation) Sleep() {
	time.Sleep(1 * time.Second)
}

func main() {
	sleeper := sleeperImplementation{}
	countdown.CountdownImproved(os.Stdout, sleeper)
}
```

Execute main

```
$ go run main.go
3
2
1
Go!
```

## Conclusion

Rundown of all the tests

```
$ go test -v ./...
?       github.com/albingeorge/tdd      [no test files]
=== RUN   TestSum
--- PASS: TestSum (0.00s)
PASS
ok      github.com/albingeorge/tdd/1_sum        0.135s
=== RUN   TestCountdownNonTdd
3
2
1
Go!--- PASS: TestCountdownNonTdd (3.00s)
=== RUN   TestCountdown
--- PASS: TestCountdown (3.00s)
=== RUN   TestCountdownImproved
--- PASS: TestCountdownImproved (0.00s)
PASS
ok      github.com/albingeorge/tdd/2_countdown  6.194s
```


## References

[Content of this talk](https://github.com/albingeorge/TDD)

[Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests/)

[Accept interfaces, return struct](https://bryanftan.medium.com/accept-interfaces-return-structs-in-go-d4cab29a301b)

[Dependency Injection design pattern](https://en.wikipedia.org/wiki/Dependency_injection)