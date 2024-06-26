// Code generated by go-swagger; DO NOT EDIT.

package dns_service

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
	"github.com/go-openapi/swag"
)

// NewDNSServiceDeleteNameserverParams creates a new DNSServiceDeleteNameserverParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDNSServiceDeleteNameserverParams() *DNSServiceDeleteNameserverParams {
	return &DNSServiceDeleteNameserverParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDNSServiceDeleteNameserverParamsWithTimeout creates a new DNSServiceDeleteNameserverParams object
// with the ability to set a timeout on a request.
func NewDNSServiceDeleteNameserverParamsWithTimeout(timeout time.Duration) *DNSServiceDeleteNameserverParams {
	return &DNSServiceDeleteNameserverParams{
		timeout: timeout,
	}
}

// NewDNSServiceDeleteNameserverParamsWithContext creates a new DNSServiceDeleteNameserverParams object
// with the ability to set a context for a request.
func NewDNSServiceDeleteNameserverParamsWithContext(ctx context.Context) *DNSServiceDeleteNameserverParams {
	return &DNSServiceDeleteNameserverParams{
		Context: ctx,
	}
}

// NewDNSServiceDeleteNameserverParamsWithHTTPClient creates a new DNSServiceDeleteNameserverParams object
// with the ability to set a custom HTTPClient for a request.
func NewDNSServiceDeleteNameserverParamsWithHTTPClient(client *http.Client) *DNSServiceDeleteNameserverParams {
	return &DNSServiceDeleteNameserverParams{
		HTTPClient: client,
	}
}

/*
DNSServiceDeleteNameserverParams contains all the parameters to send to the API endpoint

	for the Dns service delete nameserver operation.

	Typically these are written to a http.Request.
*/
type DNSServiceDeleteNameserverParams struct {

	/* Checksum.

	   The last received checksum from GetNameserverList().

	   Format: int64
	*/
	Checksum *int64

	/* Index.

	   Index of the nameserver to delete (can be received from GetNameserverList()).

	   Format: int32
	*/
	Index int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the Dns service delete nameserver params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DNSServiceDeleteNameserverParams) WithDefaults() *DNSServiceDeleteNameserverParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the Dns service delete nameserver params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DNSServiceDeleteNameserverParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the Dns service delete nameserver params
func (o *DNSServiceDeleteNameserverParams) WithTimeout(timeout time.Duration) *DNSServiceDeleteNameserverParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the Dns service delete nameserver params
func (o *DNSServiceDeleteNameserverParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the Dns service delete nameserver params
func (o *DNSServiceDeleteNameserverParams) WithContext(ctx context.Context) *DNSServiceDeleteNameserverParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the Dns service delete nameserver params
func (o *DNSServiceDeleteNameserverParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the Dns service delete nameserver params
func (o *DNSServiceDeleteNameserverParams) WithHTTPClient(client *http.Client) *DNSServiceDeleteNameserverParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the Dns service delete nameserver params
func (o *DNSServiceDeleteNameserverParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithChecksum adds the checksum to the Dns service delete nameserver params
func (o *DNSServiceDeleteNameserverParams) WithChecksum(checksum *int64) *DNSServiceDeleteNameserverParams {
	o.SetChecksum(checksum)
	return o
}

// SetChecksum adds the checksum to the Dns service delete nameserver params
func (o *DNSServiceDeleteNameserverParams) SetChecksum(checksum *int64) {
	o.Checksum = checksum
}

// WithIndex adds the index to the Dns service delete nameserver params
func (o *DNSServiceDeleteNameserverParams) WithIndex(index int32) *DNSServiceDeleteNameserverParams {
	o.SetIndex(index)
	return o
}

// SetIndex adds the index to the Dns service delete nameserver params
func (o *DNSServiceDeleteNameserverParams) SetIndex(index int32) {
	o.Index = index
}

// WriteToRequest writes these params to a swagger request
func (o *DNSServiceDeleteNameserverParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Checksum != nil {

		// query param checksum
		var qrChecksum int64

		if o.Checksum != nil {
			qrChecksum = *o.Checksum
		}
		qChecksum := swag.FormatInt64(qrChecksum)
		if qChecksum != "" {

			if err := r.SetQueryParam("checksum", qChecksum); err != nil {
				return err
			}
		}
	}

	// path param index
	if err := r.SetPathParam("index", swag.FormatInt32(o.Index)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
