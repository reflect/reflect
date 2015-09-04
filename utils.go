package reflect

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"regexp"
)

var (
	// This expression represents a valid keyspace. It can be used to validate
	// that a given string is indeed a valid keyspace name. Functions are
	// provided that wrap this functionality.
	ValidKeyspaceExp = regexp.MustCompile(`^[a-zA-Z0-9]+([\-\.=_]*[a-zA-Z0-9]+)*$`)

	// This expression represents a valid key.
	ValidKeyExp = ValidKeyspaceExp
)

func dump(i interface{}) io.Reader {
	buf := bytes.NewBuffer([]byte{})
	enc := json.NewEncoder(buf)
	enc.Encode(i)
	return buf
}

func IsValidKey(str string) bool {
	return ValidKeyExp.MatchString(str)
}

func IsValidKeyspace(str string) bool {
	return ValidKeyspaceExp.MatchString(str)
}

func logError(format string, args ...interface{}) {
	Logger.Printf("ERROR: %s", fmt.Sprintf(format, args...))
}
