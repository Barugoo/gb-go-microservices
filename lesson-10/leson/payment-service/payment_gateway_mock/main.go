package main

import (
	"encoding/json"
	"html/template"
	"net/http"
)

type InitBody struct {
	SuccessURL string `json:"success_url"`
	FailureURL string `json:"failure_url"`
}

var tmp = template.Must(template.New("payment").Parse(
	`<div><a href="{{.SuccessURL}}">Accept payment</a></div>
	<div><a href="{{.FailureURL}}">Decline payment</a></div>`))

func main() {
	http.HandleFunc("/init", func(w http.ResponseWriter, r *http.Request) {
		var body InitBody
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if err := tmp.Execute(w, &body); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
	http.ListenAndServe(":8077", nil)
}
