package main

import (
	"fmt"
	"math/rand"
	"time"
)

import (
	"github.com/chiku/linkedlists-vs-arrays/arraylist"
	"github.com/chiku/linkedlists-vs-arrays/linkedlist"
)

const (
	MAX     = 20000
	REPEATS = 100
)

func main() {
	for insertions := 0; insertions <= MAX; insertions += 100 {
		sumOfTimesForArray := 0.0
		sumOfTimesForLinkedList := 0.0

		for iterations := 0; iterations < REPEATS; iterations++ {
			list := arraylist.NewArraylist(MAX)
			start := time.Now()

			for i := 0; i < insertions; i++ {
				list.Insert(rand.Int(), rand.Intn(MAX))
			}

			elapsed := time.Since(start)
			sumOfTimesForArray += (1000.0 * elapsed.Seconds())
		}

		for iterations := 0; iterations < REPEATS; iterations++ {
			list := linkedlist.NewLinkedlist()
			start := time.Now()

			for i := 0; i < insertions; i++ {
				list.Insert(rand.Int(), rand.Intn(MAX))
			}

			elapsed := time.Since(start)
			sumOfTimesForLinkedList += (1000.0 * elapsed.Seconds())
		}

		fmt.Printf("%d,%f,%f\n", insertions, sumOfTimesForArray/REPEATS, sumOfTimesForLinkedList/REPEATS)
	}
}
