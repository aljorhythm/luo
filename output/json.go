package output

import "encoding/json"

func toJson(i interface{}) []byte {
	out, err := json.Marshal(i)

	if err == nil {
		return out
	}

	panic(err)
}
