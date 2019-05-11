package main

import "fmt"
import "math/rand"
import "time"

func main() {

	rand.Seed(time.Now().UnixNano())
	var a [10]int
	for date := 0; date < 10; date++ {
		a[date] = rand.Intn(55)
	}

	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-1-i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]

			}

		}

	}
	fmt.Println(a)
	for i := 0; i < len(a); i++ {
		fmt.Printf("%d,", a[i])
	}
	fmt.Printf("\n")
}
