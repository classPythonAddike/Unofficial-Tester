package main

import (
	"log"
	"net/http"
	"strings"
	"time"
)

var ch_no = "71"

var message_from_yan = `<head>
	<title>Unofficial Tester</title>
	<script src="https://cdnjs.cloudflare.com/ajax/libs/prism/1.24.1/components/prism-python.min.js" integrity="sha512-nvWJ2DdGeQzxIYP5eo2mqC+kXLYlH4QZ/AWYZ/yDc5EqM74jiC5lxJ+8d+6zI/H9MlsIIjrJsaRTgLtvo+Jy6A==" crossorigin="anonymous" referrerpolicy="no-referrer"></script>
</head>` +
	"<h1>Challenge " + ch_no + "</h1>" + strings.Replace(
	`
<h2>Usage:</h2>
Use the following code to test your solution -
<pre><code class="language-python">
import requests

resp = requests.get(
	"https://addike-tester.herokuapp.com/run-test",
	files = {
		"file": open("solution.py", "r") # Change this to your file name
	}
)

print(resp.text)
</code></pre>

<h2>A Quick Message From YanTovis</h2>

<strong>WARNING</strong>: My tester ignores printing in input() but official tester FAILS if you
        print something in input()
        Don't do that: input("What is the test number?")
        Use empty input: input()

Some possible errors:
<ul>
    <li> None in "Your output": Your solution didn't print for all cases.</li>
    <li> None in "Input": Your solution print more times then there is cases.</li>
    <li> If you see None in "Input" or "Your output" don't check other cases until.
        you fix problem with printing, cos "Input" and "Your output" are misaligned
        after first missing/extra print.</li>
    <li> StopIteration: Your solution try to get more input then there is test cases.</li>
</ul>
`,
	"\n",
	"<br>",
	-1,
)

func WriteMessage(w *http.ResponseWriter, message string) {
	_, err := (*w).Write([]byte(message))

	if err != nil {
		log.Printf("Error while writing message `%v` - %v\n", message, err)
	}
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			t1 := time.Now()
			next.ServeHTTP(w, r)
			t2 := time.Now()
			log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), t2.Sub(t1))
		},
	)
}
