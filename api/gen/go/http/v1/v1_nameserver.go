// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1Nameserver v1 nameserver
//
// swagger:model v1Nameserver
type V1Nameserver struct {

	// Internet address of the name server, either IPv4 or IPv6.
	Address string `json:"address,omitempty"`

	// index
	Index int32 `json:"index,omitempty"`
}

// Validate validates this v1 nameserver
func (m *V1Nameserver) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 nameserver based on context it is used
func (m *V1Nameserver) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1Nameserver) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1Nameserver) UnmarshalBinary(b []byte) error {
	var res V1Nameserver
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
