package main

import (
	"fmt"
	"log"
)

func main() {
	all, err := loadChampions()

	if err != nil {
		log.Fatalf("An error occurred loading/parsing champions, err=%v", err)
	}

	fmt.Printf("There are %d/%d champions (all): %v\n\n", len(all), cap(all), all)

	first := all[0]
	fmt.Printf("The first champion is %v\n\n", first)

	some := all[:20]
	fmt.Printf("The first %d champions are %v\n\n", len(some), some)

	// Remove an element from the slice by index.
	most := append(all[0:10], all[11:]...)

	fmt.Printf("There are %d/%d champions (most): %v\n\n", len(most), cap(most), most)
	fmt.Printf("There are %d/%d champions (all): %v\n\n", len(all), cap(all), all)

	// Remove elements by altering the slice in place.
	index := 0

	for _, champion := range all {
		if champion.hasClass("Sorcerer") {
			continue
		}

		all[index] = champion
		index++
	}

	all = all[:index]

	fmt.Printf("There are %d champions after removing Sorcerers: %v\n\n", len(all), all)
}
