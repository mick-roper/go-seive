package main

import (
	"flag"
	"log"
	"sort"
	"time"
)

var keys []int
var seive map[int]bool

func main() {
	nPtr := flag.Int("n", 100, "the upper ceiling")

	flag.Parse()

	n := *nPtr
	if n < 2 {
		log.Panic("n must be greater than or equal to 2")
	}

	seive = make(map[int]bool, n)

	log.Println("Creating seive...")

	for i := 2; i <= n; i++ {
		seive[i] = true
	}

	for i := range seive {
		keys = append(keys, i)
	}

	log.Println("sorting keys...")

	sort.Ints(keys)

	log.Println("Seiving for primes...")

	start := time.Now()

	for p := 2; p <= n; {
		for i := p * p; i <= n; i += p {
			seive[i] = false
		}

		if p*p > n { // break out of the loop here
			break
		}

		for i := range keys {
			if i > p && seive[i] {
				p = i
				break
			}
		}
	}

	elapsed := time.Since(start)

	for i := range keys {
		if seive[i] {
			log.Println(i, "is prime")
		}
	}

	log.Printf("took %v milliseconds to calculate %v primes", elapsed.Nanoseconds()/1000000, n)
}
