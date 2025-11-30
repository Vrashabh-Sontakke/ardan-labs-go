package main

import "fmt"

func main() {
	if isPal("gaooog") == true {
		fmt.Println("Is Pal")
	} else {
		fmt.Println("Not Pal")
	}
}

func isPal(text string) bool {
	rune := []rune(text)
	textLength := len(text)
	for i := 0; i < textLength/2; i++ {
		if rune[i] != rune[textLength-1-i] {
			return false
		}
	}
	return true

}
