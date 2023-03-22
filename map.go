package main

import "fmt"

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for color", color, "is", hex)
	}
}

func mapSandbox() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#234dfe",
		"white": "#ffffff",
	}

	printMap(colors)
}
