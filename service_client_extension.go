package gophercloud

import (
	"net/http"
)

type ServiceClientExtension struct {

	// ServiceClient is a reference to the ServiceClient.
	*ServiceClient

	// ProjectID is the ID of project to which User is authorized.
	ProjectID string
}

// Delete calls `Request` with the "DELETE" HTTP verb.
func (client *ServiceClient) Delete2(url string, JSONResponse interface{}, opts *RequestOpts) (*http.Response, error) {
	if opts == nil {
		opts = new(RequestOpts)
	}
	client.initReqOpts(url, nil, JSONResponse, opts)
	return client.Request("DELETE", url, opts)
}

// Delete calls `Request` with the "DELETE" HTTP verb.
func (client *ServiceClient) Delete3(url string, JSONBody interface{}, JSONResponse interface{}, opts *RequestOpts) (*http.Response, error) {
	if opts == nil {
		opts = new(RequestOpts)
	}
	client.initReqOpts(url, JSONBody, JSONResponse, opts)
	return client.Request("DELETE", url, opts)
}
