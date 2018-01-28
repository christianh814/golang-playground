package main

import "fmt"

func main() {
	colors := map[string]string {
		"red": "#ff0000",
		"green": "#74500",
	}
	colors["white"] = "#ffffff"

	//delete(colors, "white")

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("For the color", color, "the hex value is", hex)
	}
}
