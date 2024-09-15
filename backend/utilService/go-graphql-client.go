package utilService

import (
	"net/http"

	"github.com/hasura/go-graphql-client"
)

type headersTransport struct {
	headers http.Header
	base    http.RoundTripper
}

func (t *headersTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	for k, v := range t.headers {
		req.Header.Set(k, v[0])
	}
	return t.base.RoundTrip(req)
}

func Client() *graphql.Client {
	headers := http.Header{}
	headers.Add("x-hasura-admin-secret", "myadminsecretkey")
	headers.Add("x-hasura-access-key", "myadminsecretkey")

	httpClient := &http.Client{
		Transport: &headersTransport{
			headers: headers,
			base:    http.DefaultTransport,
		},
	}

	newClient := graphql.NewClient("http://hasura:8080/v1/graphql", httpClient)

	return newClient
}
