package reflect

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type KeyspaceField struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	FieldType   int       `json:"type"`
	ColumnType  string    `json:"column_type"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Keyspace struct {
	Name          string           `json:"name"`
	Slug          string           `json:"slug"`
	Status        string           `json:"status,omitempty"`
	StatisticsKey string           `json:"statistics_key"`
	Description   string           `json:"description"`
	CreatedAt     time.Time        `json:"created_at"`
	UpdatedAt     time.Time        `json:"updated_at"`
	Fields        []*KeyspaceField `json:"fields,omitempty"`
}

func (client *Client) Keyspaces() ([]*Keyspace, error) {
	req := client.newRequest("GET", fmt.Sprintf("/%s/keyspaces", defaultAPIVersion), nil)
	bytes, err := client.do(req, http.StatusOK)

	if err != nil {
		return nil, err
	}

	var keyspaces []*Keyspace

	err = json.Unmarshal(bytes, &keyspaces)

	// TODO: Handle these failure modes better. There might be errors in here...!
	if err != nil {
		return nil, err
	}

	return keyspaces, nil
}
