package reflect

import (
	"fmt"
	"net/http"
)

func tabletPath(keyspace, key string) string {
	return fmt.Sprintf("/%s/keyspaces/%s/tablets/%s", defaultAPIVersion, keyspace, key)
}

// Appends records to an existing tablet.
//
// The supplied records (recs) can be either a map or a slice of maps.
func (client *Client) Append(keyspace, key string, recs interface{}) error {
	req := client.newRequest("PUT", tabletPath(keyspace, key), dump(recs))
	_, err := client.do(req, http.StatusAccepted)
	return err
}

// Replace all the records in a tablet with a new set of records.
//
// The supplied records (recs) can be either a map or a slice of maps.
func (client *Client) Replace(keyspace, key string, body interface{}) error {
	req := client.newRequest("POST", tabletPath(keyspace, key), dump(body))
	_, err := client.do(req, http.StatusAccepted)
	return err
}
