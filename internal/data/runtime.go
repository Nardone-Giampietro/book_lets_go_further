package data

import (
	"fmt"
	"strconv"
)

type Runtime int32

// MarshalJSON Since pointer methods can only be invoked on pointers, we use the value method so we can call it for value also
func (r Runtime) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d mins", r)

	// Write it in double quotes. This in order to be a JSON string
	quotedJSONValue := strconv.Quote(jsonValue)

	return []byte(quotedJSONValue), nil
}
