package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Palindrome() {

	//To enter a string
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome!")
	fmt.Println("Please enter a string to check if it is palindrom or not!")

	// To trim all whitespace and newlines
	word, _ := reader.ReadString('\n')
	word = strings.TrimSpace(word)
	word = strings.ToLower(word)
	wordSplited := strings.Split(word, " ")
	wordTrimmed := strings.Join(wordSplited, "")

	// To do the reverse string
	reversedWord := ""
	for i := len(wordTrimmed) - 1; i >= 0; i-- {
		reversedWord += string(wordTrimmed[i])
	}

	//To compare the the reversed and the orignal
	if reversedWord == wordTrimmed {
		fmt.Println("The string is Palindrome.")
	} else {
		fmt.Println("The string is not a Palindrome.")
	}

}
func frequecy() {

	//To  enter an input string
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please Enter a String to check the frequency of the string!")

	//To trim all the new lines
	ans, _ := reader.ReadString('\n')
	ans = strings.TrimSpace(ans)

	// To remove all the punctuation
	ans2 := strings.ReplaceAll(ans, "!", "")
	ans2 = strings.ReplaceAll(ans, ".", "")
	ans2 = strings.ReplaceAll(ans, ",", "")
	ans2 = strings.ReplaceAll(ans, "?", "")

	//To remove all the whitespace and convert it to lowercase
	ans2 = strings.ToLower(ans2)
	ans3 := strings.Split(ans2, " ")
	ans4 := strings.Join(ans3, "")

	//rune:- represent Unicode code point
	//Made a map to store the string and calculate it's frequency
	res := make(map[rune]int)

	for _, char := range ans4 {

		res[char]++

	}

	//To display all the frequecy of the string
	for value, count := range res {

		//QuoteRune:-returns a single-quoted Go string literal
		var s string = strconv.QuoteRune(value)
		fmt.Println(s, count)
	}

}
func main() {

	Palindrome()
	frequecy()
}

//Clean code style
/*Write clear comments
Use meaningful names
Avoid duplicating code
use abstraction
reuse the piece of code
functions short
Use consistent formatting*/

/*
func checker(a string, b string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}*/
