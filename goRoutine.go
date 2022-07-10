package main

import (
	"time"
)

func main() {
	people := []string{"pedro", "nico"}
	c := make(chan string)
	for _, person := range people {
		go isSexy(person, c)
	}

}

func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 5)
	c <- person + "is sexy"
}
