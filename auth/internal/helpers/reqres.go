package helpers

import (
	"encoding/json"
	"net/http"
)

func Body(r *http.Request, o interface{}) error {
	return json.NewDecoder(r.Body).Decode(o)
}

func Res(o interface{}) []byte {
	s, err := json.Marshal(o)
	if err != nil {
		return []byte("{}")
	}
	return s
}
