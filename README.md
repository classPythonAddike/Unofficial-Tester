# Unofficial Tester

A simple API to simplify unofficial testing (Idea by SKP627)

Use the following snippet to test your code -
```py
import requests

resp = requests.get(
	"https://addike-tester.herokuapp.com/run-test",
	files = {
		"file": open("golfed_solution.py", "r") # Change this to your file name
	}
)

print(resp.text)
```

Make sure you've installed the `requests` module first.
