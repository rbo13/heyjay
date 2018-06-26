package heyjay

import (
	"encoding/json"
	"io"
)

// JSON returns as json from the v interface
func JSON(w io.Writer, v interface{}) (err error) {
	return json.NewEncoder(w).Encode(v)
}
