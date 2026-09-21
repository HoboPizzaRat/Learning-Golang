package main

// complete reformat function
// it takes
// message string and formatter function as input

// 1. Apply the given formatter 3 times to the message
// 2. Add to prefix of TEXTIO: to the result
// 3. Return the final string

func reformat(message string, formatter func(string) string) string {
	var output string = ""

	output = formatter(message)
	output = formatter(output)
	output = formatter(output)

	return "TEXTIO: " + output
}
