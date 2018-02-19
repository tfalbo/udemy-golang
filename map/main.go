package main

import (
	"fmt"
)

func main() {

	// var colours map[string]string

	// colours := make(map[string]string)

	colours := map[string]string{
		"red":   "#ff0000",
		"green": "#020390",
		"black": "#000000",
		"white": "#ffffff",
	}

	printMap(colours)
}

func printMap(c map[string]string) {
	for colour, hex := range c {
		fmt.Println("Hex code for", colour, "is", hex)
	}
}
