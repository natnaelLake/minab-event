package utilService

import (
	"net/http"

	"github.com/hasura/go-graphql-client"
)

// headersTransport struct allows you to inject custom headers into every request made by the HTTP client
type headersTransport struct {
	headers http.Header
	base    http.RoundTripper
}

// RoundTrip method sets custom headers for each request
func (t *headersTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	for k, v := range t.headers {
		req.Header.Set(k, v[0])
	}
	return t.base.RoundTrip(req)
}

// Client function initializes and returns a GraphQL client configured with the necessary headers
func Client() *graphql.Client {	
	// Set up the HTTP client with the request headers
	headers := http.Header{}
	headers.Add("x-hasura-admin-secret", "adminsecretkey") // Adjust this as needed
	headers.Add("x-hasura-access-key", "adminaccesskey")  // Add this if required

	// Create an HTTP transport that adds headers to requests
	httpClient := &http.Client{
		Transport: &headersTransport{
			headers: headers,
			base:    http.DefaultTransport,
		},
	}

	// Set up the GraphQL client with the Hasura endpoint
	newClient := graphql.NewClient("http://hasura:8080/v1/graphql", httpClient)

	return newClient
}
