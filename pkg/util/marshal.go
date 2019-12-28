package util

import "encoding/json"

// MustMarshal will marshal JSON or panic
func MustMarshal(v interface{}) json.RawMessage {
	data, err := json.Marshal(v)

	if err != nil {
		panic(err)
	}

	return json.RawMessage(data)
}
