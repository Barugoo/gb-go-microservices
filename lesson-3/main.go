package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func sum(a, b int) (result int, err error) {
	m := map[string]interface{}{
		"a": a,
		"b": b,
	}
	bytes, err := json.Marshal(m)
	if err != nil {
		return 0, err
	}
	req, err := http.NewRequest(http.MethodPost, "/sum", bytes.NewReader(bytes))
	if err != nil {
		return 0, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}

	if resp.StatusCode > 399 {
		return 0, mapStatusCodeToError(resp.StatusCode)
	}

	s := struct {
		Result int `json:"result"`
	}{}
	if err := json.NewDecoder(resp.Body).Decode(&s); err != nil {
		return 0, err
	}
	return s.Result, nil
}

func concat(a, b string) (result string, err error) {

	http.NewRequest(http.MethodPost, "/div", bytes.NewReader())
}
