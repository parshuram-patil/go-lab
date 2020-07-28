package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type EmpStruct struct {
	EmpID     string  `json:"emp_id"`
	EmpName   string  `json:"emp_name"`
	EmpSalary float64 `json:"emp_salary"`
}

func Emphandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		io.WriteString(w, "<h1>EmpHandler GET Page</h1>")
	case "POST":
		/* io.WriteString(w, "<h1>EmpHandler POST Page</h1>")
		fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm) */
		reqBody, _ := ioutil.ReadAll(r.Body)
		var emp EmpStruct
		json.Unmarshal(reqBody, &emp)
		fmt.Println("Emp POST Req body --->", emp)
		Save(emp)
		json.NewEncoder(w).Encode(emp)

	}
}
