package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"

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
			"test_cases_" + ch_no + ".py",
			string(testCases),
		},
	}

	result, err := program.Execute(120000) // 120s timeout

	if err != nil {
		WriteMessage(
			&w,
			"Error while making request to WandBox! -"+err.Error(),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	WriteMessage(
		&w,
		result.ProgramMessage,
	)
}
