package main

import (
	"fmt"
	"strings"
)


func main(){
	fmt.Println("Hello, World!")
}

func cleanInput(text string) []string {
	var cleaned []string
	words := strings.Fields(strings.ToLower(text))
	cleaned = append(cleaned, words...)
	return cleaned
}
