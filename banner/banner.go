package main

import "fmt"

func main() {
	banner(10, "ardan labs")
}

func banner(bannerWidth int, text string) {
	margin := (bannerWidth - len(text)) / 2
	if margin == 0 {
		margin = 2
	}
	for i := 0; i < margin; i++ {
		fmt.Print(" ")
	}
	fmt.Println(text)
	for i := 0; i < bannerWidth; i++ {
		fmt.Print("-")
	}
}
