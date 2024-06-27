// Code generated by go-swagger; DO NOT EDIT.

package hostname_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/stepan-tikunov/hostname-dns-configurer/api/gen/go/http/v1"
)

// NewHostnameServiceGetHostnameParams creates a new HostnameServiceGetHostnameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewHostnameServiceGetHostnameParams() *HostnameServiceGetHostnameParams {
	return &HostnameServiceGetHostnameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewHostnameServiceGetHostnameParamsWithTimeout creates a new HostnameServiceGetHostnameParams object
// with the ability to set a timeout on a request.
func NewHostnameServiceGetHostnameParamsWithTimeout(timeout time.Duration) *HostnameServiceGetHostnameParams {
	return &HostnameServiceGetHostnameParams{
		timeout: timeout,
	}
}

// NewHostnameServiceGetHostnameParamsWithContext creates a new HostnameServiceGetHostnameParams object
// with the ability to set a context for a request.
func NewHostnameServiceGetHostnameParamsWithContext(ctx context.Context) *HostnameServiceGetHostnameParams {
	return &HostnameServiceGetHostnameParams{
		Context: ctx,
	}
}

// NewHostnameServiceGetHostnameParamsWithHTTPClient creates a new HostnameServiceGetHostnameParams object
// with the ability to set a custom HTTPClient for a request.
func NewHostnameServiceGetHostnameParamsWithHTTPClient(client *http.Client) *HostnameServiceGetHostnameParams {
	return &HostnameServiceGetHostnameParams{
		HTTPClient: client,
	}
}

/*
HostnameServiceGetHostnameParams contains all the parameters to send to the API endpoint

	for the hostname service get hostname operation.

	Typically these are written to a http.Request.
*/
type HostnameServiceGetHostnameParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the hostname service get hostname params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HostnameServiceGetHostnameParams) WithDefaults() *HostnameServiceGetHostnameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the hostname service get hostname params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HostnameServiceGetHostnameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the hostname service get hostname params
func (o *HostnameServiceGetHostnameParams) WithTimeout(timeout time.Duration) *HostnameServiceGetHostnameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the hostname service get hostname params
func (o *HostnameServiceGetHostnameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the hostname service get hostname params
func (o *HostnameServiceGetHostnameParams) WithContext(ctx context.Context) *HostnameServiceGetHostnameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the hostname service get hostname params
func (o *HostnameServiceGetHostnameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the hostname service get hostname params
func (o *HostnameServiceGetHostnameParams) WithHTTPClient(client *http.Client) *HostnameServiceGetHostnameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the hostname service get hostname params
func (o *HostnameServiceGetHostnameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *HostnameServiceGetHostnameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}