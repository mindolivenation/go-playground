package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#008000",
		"white": "#ffffff",
	}

	fmt.Println(colors)

	colors2 := make(map[string]string)
	colors2["white"] = "$foo"
	// delete(colors2, "white")
	fmt.Println(colors2)

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Printf("%v %v\n", color, hex)
	}
}
