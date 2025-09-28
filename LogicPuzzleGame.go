package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Logic Puzzle Game")
	fmt.Println("Three friends: Alice, Bob, Carol each have a different pet: Cat, Dog, Bird.")
	fmt.Println("Clues:")
	fmt.Println("1. Alice does not have the Dog.")
	fmt.Println("2. Bob does not have the Cat.")
	fmt.Println("3. Carol has the Bird.")
	fmt.Println("Who owns which pet?")
	solutions := []struct{ A, B, C string }{
		{"Cat", "Dog", "Bird"},
		{"Dog", "Cat", "Bird"},
		{"Cat", "Bird", "Dog"},
	}
	for {
		fmt.Print("Enter Alice's pet: ")
		var alice string
		fmt.Scanln(&alice)
		fmt.Print("Enter Bob's pet: ")
		var bob string
		fmt.Scanln(&bob)
		fmt.Print("Enter Carol's pet: ")
		var carol string
		fmt.Scanln(&carol)
		alice, bob, carol = strings.Title(alice), strings.Title(bob), strings.Title(carol)
		if alice==bob || bob==carol || alice==carol {
			fmt.Println("Each pet must be unique!")
			continue
		}
		ok := false
		for _, s := range solutions {
			if s.A==alice && s.B==bob && s.C==carol {
				ok = true
				break
			}
		}
		if ok && alice!="Dog" && bob!="Cat" && carol=="Bird" {
			fmt.Println("Correct! Alice has the Cat, Bob the Dog, Carol the Bird.")
			break
		} else {
			fmt.Println("Incorrect. Try again.")
		}
	}
}