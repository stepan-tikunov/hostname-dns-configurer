// Code generated by go-swagger; DO NOT EDIT.

package dns_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	v1 "github.com/stepan-tikunov/hostname-dns-configurer/api/gen/go/http/v1"
)

// DNSServiceGetNameserverAtReader is a Reader for the DNSServiceGetNameserverAt structure.
type DNSServiceGetNameserverAtReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DNSServiceGetNameserverAtReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDNSServiceGetNameserverAtOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDNSServiceGetNameserverAtDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDNSServiceGetNameserverAtOK creates a DNSServiceGetNameserverAtOK with default headers values
func NewDNSServiceGetNameserverAtOK() *DNSServiceGetNameserverAtOK {
	return &DNSServiceGetNameserverAtOK{}
}

/*
DNSServiceGetNameserverAtOK describes a response with status code 200, with default header values.

A successful response.
*/
type DNSServiceGetNameserverAtOK struct {
	Payload *v1.V1NameserverResponse
}

// IsSuccess returns true when this dns service get nameserver at o k response has a 2xx status code
func (o *DNSServiceGetNameserverAtOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dns service get nameserver at o k response has a 3xx status code
func (o *DNSServiceGetNameserverAtOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dns service get nameserver at o k response has a 4xx status code
func (o *DNSServiceGetNameserverAtOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dns service get nameserver at o k response has a 5xx status code
func (o *DNSServiceGetNameserverAtOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dns service get nameserver at o k response a status code equal to that given
func (o *DNSServiceGetNameserverAtOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dns service get nameserver at o k response
func (o *DNSServiceGetNameserverAtOK) Code() int {
	return 200
}

func (o *DNSServiceGetNameserverAtOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/dns/{index}][%d] dnsServiceGetNameserverAtOK %s", 200, payload)
}

func (o *DNSServiceGetNameserverAtOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/dns/{index}][%d] dnsServiceGetNameserverAtOK %s", 200, payload)
}

func (o *DNSServiceGetNameserverAtOK) GetPayload() *v1.V1NameserverResponse {
	return o.Payload
}

func (o *DNSServiceGetNameserverAtOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.V1NameserverResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDNSServiceGetNameserverAtDefault creates a DNSServiceGetNameserverAtDefault with default headers values
func NewDNSServiceGetNameserverAtDefault(code int) *DNSServiceGetNameserverAtDefault {
	return &DNSServiceGetNameserverAtDefault{
		_statusCode: code,
	}
}

/*
DNSServiceGetNameserverAtDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type DNSServiceGetNameserverAtDefault struct {
	_statusCode int

	Payload *v1.RPCStatus
}

// IsSuccess returns true when this Dns service get nameserver at default response has a 2xx status code
func (o *DNSServiceGetNameserverAtDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this Dns service get nameserver at default response has a 3xx status code
func (o *DNSServiceGetNameserverAtDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this Dns service get nameserver at default response has a 4xx status code
func (o *DNSServiceGetNameserverAtDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this Dns service get nameserver at default response has a 5xx status code
func (o *DNSServiceGetNameserverAtDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this Dns service get nameserver at default response a status code equal to that given
func (o *DNSServiceGetNameserverAtDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the Dns service get nameserver at default response
func (o *DNSServiceGetNameserverAtDefault) Code() int {
	return o._statusCode
}

func (o *DNSServiceGetNameserverAtDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/dns/{index}][%d] DnsService_GetNameserverAt default %s", o._statusCode, payload)
}

func (o *DNSServiceGetNameserverAtDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/dns/{index}][%d] DnsService_GetNameserverAt default %s", o._statusCode, payload)
}

func (o *DNSServiceGetNameserverAtDefault) GetPayload() *v1.RPCStatus {
	return o.Payload
}

func (o *DNSServiceGetNameserverAtDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}