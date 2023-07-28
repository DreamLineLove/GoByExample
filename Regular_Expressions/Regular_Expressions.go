package Regular_Expressions

import (
	"fmt"
	"regexp"
)

func Regular_Expressions() {
	printRegexOperation("a", "rawyueryuiqwery", "Should match if a is in the str")

	printRegexOperation("^a", "rawyueryuiqwery", "Should match only if the str starts with a (doesn't matter if the str contains any a's)")
}

func printRegexOperation(exp, str, desc string) (err error) {
	fmt.Print("#\t", desc, "\n")
	fmt.Println("pattern:  ", exp)
	fmt.Println("string:   ", str)
	if matched, err := regexp.MatchString(exp, str); err != nil {
		fmt.Println()
		return fmt.Errorf("error while regex-matching: %w", err)
	} else {
		fmt.Println("-matched: ", matched)
		fmt.Println()
		return nil
	}
}
