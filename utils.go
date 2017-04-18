package reflect

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

var (
	Logger = log.New(ioutil.Discard, "", 0)
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
// authentication token for use when authenticating clients.
//
// Here is an example of how to generate a signed authentication token.
//
//	params := []reflect.Parameter{
//	  {
//	    Field: "Field1",
//	    Op:    reflect.EqualsOperation,
//	    Value: "abc123",
//	  },
//	}
//
//	tok := reflect.GenerateToken("<Your Secret Key>", params)
func GenerateToken(secretKey string, parameters []Parameter) string {
	var params []string

	// For each of the parameters, we'll compile them in to a string with the
	// required format.
	for _, p := range parameters {
		// Put values in order since they're always ORed.
		vals := make([]string, len(p.AnyValues))
		copy(vals, p.AnyValues)
		sort.Strings(vals)

		enc := []interface{}{
			p.Field,
			p.Op,
			p.Value,
			vals,
		}

		v, err := json.Marshal(enc)
		if err != nil {
			logError("Could not encode token: %v", err)
			return ""
		}

		params = append(params, string(v))
	}

	// Now that we have the params we need to sort them all.
	sort.Strings(params)

	msg := []byte(fmt.Sprintf("V2\n%s", strings.Join(params, "\n")))

	mac := hmac.New(sha256.New, []byte(secretKey))
	mac.Write(msg)

	return fmt.Sprintf("=2=%s", base64.StdEncoding.EncodeToString(mac.Sum(nil)))
}
