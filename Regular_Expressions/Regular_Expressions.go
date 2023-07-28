package Regular_Expressions

import (
	"fmt"
	"regexp"
)

func Regular_Expressions() {
	printRegexOperation("a", "rawyueryuiqwery", "Should match if a is in the str", true)

	printRegexOperation("^a", "rawyueryuiqwery", "Should match only if the str starts with a (doesn't matter if the str contains any a's)", false)

	printRegexOperation("yoma$", "bank of yoma", "Should match at the end of string", true)
	printRegexOperation("yoma$", "yoma bank", "should not match", false)
}

func printRegexOperation(pattern, str, desc string, exp bool) (err error) {
	fmt.Print("#\t", desc, "\n")
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
