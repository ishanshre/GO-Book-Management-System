package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

/*
Marshel and Unmarshell Json
	- convert to json and vice versa
*/

func ParseBody(r *http.Request, x interface{}) {
	if body, err := io.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}
