package main

import "fmt"

func main() {
	fmt.Println(concat([]string{"A", "B"}, []string{"C", "D", "E"}))
	fmt.Println(concat2([]string{"A", "B"}, []string{"C", "D", "E"}))
}

func concat(s1, s2 []string) []string {
	combinedSlice := append(s1, s2...)
	return combinedSlice
}

func concat2(s1, s2 []string) []string {
	s := make([]string, len(s1)+len(s2))
	copy(s, s1)
	copy(s[len(s1):], s2)
	return s
}
