package transport

import "net/http"

type HTTPTransport struct {
	Client *http.Client
}

func NewHTTPTransport(client *http.Client) HTTPTransport {
	if client == nil {
		client = http.DefaultClient
	}
	return HTTPTransport{Client: client}
}
