package reflect

import (
	"fmt"
	"net/http"
)

const (
	CriteriaHeaderName = "X-Criteria"
	UpsertHeaderName   = "X-Insert-Missing"
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

// Update a set of records in a tablet using criteria to match records between
// the supplied records and the records that exist in the tablet. Any supplied
// records that don't match will be dropped.
//
// The supplied records (recs) can be either a map or a slice of maps.
func (client *Client) Patch(keyspace, key string, criteria []string, body interface{}) error {
	req := client.newRequest("PATCH", tabletPath(keyspace, key), dump(body))
	req.Header.Add(CriteriaHeaderName, strings.Join(criteria, ", "))
	_, err := client.do(req, http.StatusAccepted)
	return err
}

// Update or insert a set of records in a tablet using criteria to match
// records between the supplied records and the records that exist in the
// tablet. Any records that don't match will be appended.
//
// The supplied records (recs) can be either a map or a slice of maps.
func (client *Client) Upsert(keyspace, key string, criteria []string, body interface{}) error {
	req := client.newRequest("PATCH", tabletPath(keyspace, key), dump(body))
	req.Header.Add(CriteriaHeaderName, strings.Join(criteria, ", "))
	req.Header.Add(UpsertHeaderName, true)

	_, err := client.do(req, http.StatusAccepted)
	return err
}
