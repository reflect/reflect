package reflect

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"sort"
	"strings"
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

// Given a secret key and a set of parameters, generates a new signed
// authentication token for using when authenticing clients.
func GenerateToken(secretKey string, parameters []Parameter) string {
	var params []string

	// For each of the parameters, we'll compile them in to a string with the
	// required format.
	for _, p := range parameters {
		params = append(params, fmt.Sprintf("%s %s %s", p.Field, p.Op, p.Value))
	}

	// Now that we have the params we need to sort them all.
	sort.Strings(params)

	mac := hmac.New(sha256.New, []byte(secretKey))
	mac.Write([]byte(strings.Join(params, "\n")))

	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}
