package reflect

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

func dump(i interface{}) io.Reader {
	buf := bytes.NewBuffer([]byte{})
	enc := json.NewEncoder(buf)
	enc.Encode(i)
	return buf
}

func logError(format string, args ...interface{}) {
	Logger.Printf("ERROR: %s", fmt.Sprintf(format, args...))
}
