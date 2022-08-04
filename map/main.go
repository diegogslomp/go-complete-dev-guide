package main

import "fmt"

func main() {
	// var colors map[string]string
	colors := make(map[string]string)

	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#00FF00",
	// }

	colors["white"] = "#FFFFFF"
	delete(colors, "white")

	fmt.Println(colors)
}
