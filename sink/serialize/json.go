package serialize

import "encoding/json"

func Json[IN interface{}](in IN) []byte {
	d, _ := json.Marshal(in)
	return d
}
