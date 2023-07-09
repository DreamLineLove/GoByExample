# Errors 

## Expectations
1. Error is a built-in type in Go. You can create custom errors by implementing the Error() function on a custom error struct (utilizing interface). Then you may use any Errors functions with your custom error type.
2. Errors.New() creates a new error with a string message.
3. fmt.Errorf() returns an error (message) with one or more errors (custom or built-in) with the required format. %w will implement unwrap which will return the wrapped error or a slice of wrapped errors. 
4. %w only accepts error values.
5. %v also accepts errors but does not implement the Unwrap() method.
6. You can include both the function call and its variables, as well as the testing logic for the returned error variable, on the same line in an if statement, separated by the ';'.

## Other Findings
1. To implement error interface on any custom type you must implement the Error() method that returns a string value! 
