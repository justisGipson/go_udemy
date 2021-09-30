package main

import "fmt"

// 3 different ways to create maps

func main() {
	// map w/ keys of type "string"
	// and values of type "string"
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	// var colors map[string]string
	// colors := make(map[string]string)

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is:", hex)
	}
}
