package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	str := "Hello World"

	fmt.Println(len(str))

	str1 := "Hello"
	str2 := "World"
	result := str1 + " " + str2
	fmt.Println(len(result))

	fmt.Println(str1[0])

	fmt.Println(result[1 : len(result)-1])

	// String conversion
	num := 19
	str3 := strconv.Itoa(num) // From int to string
	fmt.Println(len(str3))

	// Strings splitting
	fruits := "apple, orange, banana, mango"
	parts := strings.Split(fruits, ", ")
	for _, part := range parts {
		fmt.Println(part)
	}

	// Joining strings
	countries := []string{"Mexico", "USA", "Canada"}
	joinedCountries := strings.Join(countries, ", ")
	fmt.Println(joinedCountries)

	// Substring contain
	fmt.Println(strings.Contains(joinedCountries, "Sweden"))

	// Substring replace
	replaced := strings.ReplaceAll(joinedCountries, ", ", "-")
	fmt.Println(replaced)

	// Trim
	strwspace := " Hey I have space "
	fmt.Println(strwspace)
	fmt.Println(strings.TrimSpace(strwspace))

	// Upper / lower transformation
	fmt.Println(strings.ToUpper(joinedCountries))
	fmt.Println(strings.ToLower(joinedCountries))

	// Repeat
	fmt.Println(strings.Repeat("*", len(joinedCountries)))
	fmt.Println(joinedCountries)
	fmt.Println(strings.Repeat("*", len(joinedCountries)))

	// Count runes
	fmt.Println(strings.Count("Hello World", "l"))

	// Prefix / suffix evaluation
	fmt.Println(strings.HasPrefix("doctor", "doc"))
	fmt.Println(strings.HasPrefix("doctor", "dr"))
	fmt.Println(strings.HasSuffix("German Shepherd", "Shepherd"))

	// REGEX
	str5 := "Hello, 123 Go! 88"
	re := regexp.MustCompile("\\d+")
	matches := re.FindAllString(str5, -1)
	for _, match := range matches {
		fmt.Println(match)
	}

	// UNICODE UTF-8
	str6 := "Hello, ñiño llanero"
	fmt.Println(utf8.RuneCountInString(str6))

	// STRING BUILDER
	var builder strings.Builder

	// Write some strings (More efficient and performant than concatenation)
	builder.WriteString("Hello, <UNK>")
	builder.WriteString(" World!")
	builder.WriteRune('x')

	// Convert builder into string
	finalString := builder.String()
	fmt.Println(finalString)

	// Reset
	builder.Reset()

	builder.WriteString("Starting fresh!")
	finalString = builder.String()
	fmt.Println(finalString)

}
