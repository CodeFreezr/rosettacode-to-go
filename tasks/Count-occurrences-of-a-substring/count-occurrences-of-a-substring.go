package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Count("the three truths", "th")) // says: 3
	fmt.Println(strings.Count("ababababab", "abab"))     // says: 2
}

//\Count-occurrences-of-a-substring\count-occurrences-of-a-substring.go
