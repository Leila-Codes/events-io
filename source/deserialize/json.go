package deserialize

import (
	"encoding/json"
)

func Json[OUT interface{}](d []byte) OUT {
	v := new(OUT)
	err := json.Unmarshal(d, v)
	if err != nil {
		panic(err)
	}

	return *v
}
