package main

import "fmt"

func main() {
	// var colors map[string]string

	colors := make(map[string]string)
	// colors := map[string]string{
	// 	"red": "#ff0000",
	// 	"green": "#f23434",
	// }

	colors["white"] = "#ffffff"

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex color for",color,"is", hex)
	}
}