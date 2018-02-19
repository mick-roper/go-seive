# go-seive
sieve of Eratosthenes written in go!

## Requirements:
Go 1.9.3

## Running the app:

1: Run ```go build ./app/main.go``` at the root of the project
2: Run ```./main -n=x``` where x is the upper boundary of the range of primes you want to generate
3: Watch the console as the app chews on your request and prints the output

## More stuff to do:

* Speed up the app by calculating asynchronously
* Add a mechanism to test if a value is prime