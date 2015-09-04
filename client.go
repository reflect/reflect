package reflect

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	DefaultHost = "https://api.reflect.io"

	// Version of the API to use.
	defaultAPIVersion = "v1"
)

var (
	DefaultUserAgent = fmt.Sprintf("Reflect Golang Client v%s", Version)

	// Transport used by all HTTP requests. It is not advised that you override
	// this unless you really have to.
	DefaultTransport = http.Transport{
		DisableCompression: false,
	}

	Logger = log.New(ioutil.Discard, "", 0)
)

type Client struct {
	// The host to use when connecting to the service.
	Host string

	// HTTP client used for connecting to the service.
	HTTPClient *http.Client

	// API token used for authentication.
	token string
}

// Create a new HTTP request and set the appropriate authentication parameters.
func (client *Client) newRequest(verb, path string, body io.Reader) *http.Request {
	req, _ := http.NewRequest(verb, fmt.Sprintf("%s%s", client.Host, path), body)
	req.SetBasicAuth("", client.token)
	req.Header.Set("User-Agent", DefaultUserAgent)
	req.Header.Set("Content-Type", "application/json")
	return req
}

func (client *Client) do(req *http.Request, expect int) ([]byte, error) {
	res, err := client.HTTPClient.Do(req)

	// Pretty sure the only error that could occur here would be if we failed to
	// connect to the host for some reason.
	if err != nil {
		logError("Error connecting to Reflect API. %v", err)
		return []byte{}, ErrConnectionFailed
	}

	defer res.Body.Close()

	switch res.StatusCode {
	case expect:
		return ioutil.ReadAll(res.Body)
	case http.StatusInternalServerError, http.StatusBadGateway, http.StatusServiceUnavailable:
		logError("(Status %d) Internal service error on Reflect's end.", res.StatusCode)
		return []byte{}, ErrInternal
	case http.StatusUnauthorized:
		logError("(Status %d) Authentication failed.", res.StatusCode)
		return []byte{}, ErrNotAuthenticated
	}

	// If we get here it means we have no idea what happened. Unfortunately.
	re := NewErrorFromResponse(res)
	return []byte{}, re
}

// Create a new instance of a client for talking to Reflect. The connection and
// authentication information isn't actually validated until the first request.
func NewClient(tok string) *Client {
	client := new(Client)
	client.token = tok
	client.HTTPClient = &http.Client{Transport: &DefaultTransport}
	client.Host = DefaultHost
	return client
}
