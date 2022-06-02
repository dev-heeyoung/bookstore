package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	body, err1 := ioutil.ReadAll(r.Body)
	if err1 == nil {
		err2 := json.Unmarshal([]byte(body), x)
		if err2 != nil {
			return
		}
	}
}
