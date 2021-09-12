package main

import (
	"log"
	"net/http"
)

var message_from_yan = `
<h1>A Quick Message From YanTovis</h1> -

How to use?
Import your solution in line 92, and run this script
If you see some weird chars instead of colors in output or don't want colors
switch COLOR_OUT to False in line 30

WARNING: My tester ignores printing in input() but official tester FAILS if you
        print something in input()
        Don't do that: input("What is the test number?")
        Use empty input: input()

Some possible errors:
    - None in "Your output": Your solution didn't print for all cases.
    - None in "Input": Your solution print more times then there is cases.
    - If you see None in "Input" or "Your output" don't check other cases until
        you fix problem with printing, cos "Input" and "Your output" are misaligned
        after first missing/extra print
    - StopIteration: Your solution try to get more input then there is test cases
`

func WriteMessage(w *http.ResponseWriter, message string, code int) {
	(*w).WriteHeader(code)
	_, err := (*w).Write([]byte(message))

	if err != nil {
		log.Printf("Error while writing message `%v` - %v\n", message, err)
	}
}