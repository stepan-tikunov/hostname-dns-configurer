// Code generated by go-swagger; DO NOT EDIT.

package hostname_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new hostname service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new hostname service API client with basic auth credentials.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - user: user for basic authentication header.
// - password: password for basic authentication header.
func NewClientWithBasicAuth(host, basePath, scheme, user, password string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BasicAuth(user, password)
	return &Client{transport: transport, formats: strfmt.Default}
}

// New creates a new hostname service API client with a bearer token for authentication.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - bearerToken: bearer token for Bearer authentication header.
func NewClientWithBearerToken(host, basePath, scheme, bearerToken string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BearerToken(bearerToken)
	return &Client{transport: transport, formats: strfmt.Default}
}

/*
Client for hostname service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	HostnameServiceGetHostname(params *HostnameServiceGetHostnameParams, opts ...ClientOption) (*HostnameServiceGetHostnameOK, error)

	HostnameServiceSetHostname(params *HostnameServiceSetHostnameParams, opts ...ClientOption) (*HostnameServiceSetHostnameOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
HostnameServiceGetHostname gets hostname
*/
func (a *Client) HostnameServiceGetHostname(params *HostnameServiceGetHostnameParams, opts ...ClientOption) (*HostnameServiceGetHostnameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHostnameServiceGetHostnameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "HostnameService_GetHostname",
		Method:             "GET",
		PathPattern:        "/api/hostname",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &HostnameServiceGetHostnameReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*HostnameServiceGetHostnameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*HostnameServiceGetHostnameDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
HostnameServiceSetHostname sets hostname
*/
func (a *Client) HostnameServiceSetHostname(params *HostnameServiceSetHostnameParams, opts ...ClientOption) (*HostnameServiceSetHostnameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHostnameServiceSetHostnameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "HostnameService_SetHostname",
		Method:             "POST",
		PathPattern:        "/api/hostname",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &HostnameServiceSetHostnameReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*HostnameServiceSetHostnameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*HostnameServiceSetHostnameDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
