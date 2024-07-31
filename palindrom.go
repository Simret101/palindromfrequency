package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	for {
		Palindrome()
		Frequency()

		//To Ask the person if they want to continue
		fmt.Println("Do you want to check  string again? say yes or no,Thank You!")
		reader := bufio.NewReader(os.Stdin)
		answer, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
		answer = strings.TrimSpace(answer)

		if strings.ToLower(answer) != "yes" {
			fmt.Println("Goodbye!")
			return
		}
	}
}

// To check the palindrom
func Palindrome() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome!")
	fmt.Println("Please enter a string to check if it is palindrome or not!")

	word, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	// To trim all whitespace and newlines
	word = strings.TrimSpace(word)

	// Check if the input is empty
	if word == "" {
		fmt.Println("Error: Input string cannot be empty.")
		return
	}

	// Check if the input contains numbers
	if containsNumber(word) {
		fmt.Println("Error: Input string cannot contain numbers.")
		return
	}

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

// To check frequency
func Frequency() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please Enter a String to check the frequency of the string!")

	//To trim all the new lines
	ans, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	ans = strings.TrimSpace(ans)

	//To Check if the input is empty
	if ans == "" {
		fmt.Println("Error: Input string cannot be empty.")
		return
	}

	//To Check if the input contains numbers
	if containsNumber(ans) {
		fmt.Println("Error: Input string cannot contain numbers.")
		return
	}

	// To remove all the punctuation
	ans2 := strings.ReplaceAll(ans, "!", "")
	ans2 = strings.ReplaceAll(ans2, ".", "")
	ans2 = strings.ReplaceAll(ans2, ",", "")
	ans2 = strings.ReplaceAll(ans2, "?", "")

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
		var s string = fmt.Sprintf("%c", value)
		fmt.Println(s, count)
	}
}

func containsNumber(s string) bool {
	for _, r := range s {
		if unicode.IsNumber(r) {
			return true
		}
	}
	return false
}

//Clean code style
/*Write clear comments
Use meaningful names
Avoid duplicating code
use abstraction
reuse the piece of code
functions short
Use consistent formatting*/


