package Regular_Expressions

import (
	"fmt"
	"regexp"
)

func Regular_Expressions() {
	// exp "a" will match any str that has a in it
	printRegexOperation("a", "rawyueryuiqwery", true)
	printRegexOperation("bear", "raw b jdi ea ghal r", false)
	printRegexOperation("bear", "the bear should match", true)
	printRegexOperation("bear", "the Bear should not match", true)

	// exp "^a" will match any str that has begins with a
	printRegexOperation("^a", "rawyueryuiqwery", false)

	// exp "yoma$" will match any str that ends with yoma
	printRegexOperation("yoma$", "bank of yoma", true)
	printRegexOperation("yoma$", "yoma bank", false)

	// Quantifiers
	// * ? + {} {,}
	printRegexOperation("abc*", "ab ab ab ab", true)
	printRegexOperation("abc?", "ab ab ab ab", true)
	printRegexOperation("abc+", "ab ab ab ab", false)
	printRegexOperation("ab{2}c", "abbc ccasdf", true)
	printRegexOperation("https?", "https://google.com", true)
	printRegexOperation("https?", "http://google.com", true)
	printRegexOperation("https*", "http://google.com", true)
	printRegexOperation("https+", "http://google.com", false)
	printRegexOperation("ht{3}ps*", "http://google.com", false)
	printRegexOperation("xyz{2,}", "asdfhk-xyzzzzzz2304", true)
	printRegexOperation("xyz{2}", "asdfhk-xyzzzzzz2304", false)
	printRegexOperation("xyz{2, 2}", "asdfhk-xyzzzzzz2304", false)

	// Or operator
	printRegexOperation("ht{2}p[s|]", "https://google.com", true)
	printRegexOperation("ht{2}p[t|]", "https://google.com", false)

	// Flags
	printRegexOperation("(?i)bear", "asdfjasdf. asdklfkl.w. The Bear should match this time.", true)
	printRegexOperation("(?i)USA", "usa", true)

	// Brackets
	printRegexOperation("a|b|c", "b ghhgsdjfgdf", true)
	printRegexOperation("a|b|c", "tuertuoe", false)
	printRegexOperation("[abc]", "tuertuoe", false)
	printRegexOperation("[abc]", "b ghhgsdjfgdf", true)
}

func printRegexOperation(pattern, str string, exp bool) (err error) {
	fmt.Println("pattern:  ", pattern)
	fmt.Println("string:   ", str)
	fmt.Println("- expected output:   ", exp)
	if matched, err := regexp.MatchString(pattern, str); err != nil {
		fmt.Println()
		return fmt.Errorf("error while regex-matching: %w", err)
	} else {
		fmt.Println("\t- matched:   ", matched)
		fmt.Println()
		return nil
	}
}
