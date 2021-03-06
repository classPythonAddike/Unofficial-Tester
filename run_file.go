package main

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	gwb "github.com/classPythonAddike/gowandbox"
)

func RunFile(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseMultipartForm(1 << 20) // 1 MB is the max file size file size

	file, _, err := r.FormFile("file")

	if err != nil {
		WriteMessage(
			&w,
			"Error while reading file! Please make sure you uploaded it under the name `file`!",
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	defer file.Close()

	var buf bytes.Buffer

	_, err = io.Copy(&buf, file)
	if err != nil {
		WriteMessage(
			&w,
			"Error while reading file!",
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	content := buf.String()

	buf.Reset()

	// ---- Testing the file on WandBox ----

	testerCode, err := ioutil.ReadFile("tester.py")
	if err != nil {
		WriteMessage(
			&w,
			"Whoops! Looks like the tester hasn't been uploaded! Please ping @class PythonAddike to remind him.",
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	testCases, err := ioutil.ReadFile("test_cases.py")
	if err != nil {
		WriteMessage(
			&w,
			"Whoops! Looks like the testcases haven't been uploaded! Please ping @class PythonAddike to remind him.",
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	program := gwb.NewGWBProgram()
	program.Code = string(testerCode)
	program.Options = "warning"
	program.Compiler = "cpython-3.8.0"

	program.Codes = []gwb.Program{
		{
			"to_submit_ch_" + ch_no + ".py",
			content,
		},

		{
			"test_cases_ch_" + ch_no + ".py",
			string(testCases),
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	result, err := program.Execute(ctx) // 60s timeout

	if err != nil {

		if err.Error() == "" {
			WriteMessage(
				&w,
				"Wandbox returned an ratelimit error while trying to run your code. Please try again in a few minutes!",
			)
			w.WriteHeader(http.StatusTooManyRequests)
			return
		}

		WriteMessage(
			&w,
			"Error while making request to WandBox! -"+err.Error(),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	WriteMessage(
		&w,
		"Note - this weeks testcases have been reduced to 1500 in number, from 5001, due to WandBox not allowing large files. If you want to run all 5000 cases, download the tester from  Yan's Github - https://github.com/Pomroka/TWT_Challenges_Tester \n\n",
	)

	WriteMessage(
		&w,
		result.ProgramMessage,
	)
}
