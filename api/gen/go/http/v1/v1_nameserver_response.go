// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1NameserverResponse v1 nameserver response
//
// swagger:model v1NameserverResponse
type V1NameserverResponse struct {

	// The checksum of all meaningful fields in resolv.conf file.
	// Must be used to prevent cases when client operates on outdated data.
	Checksum int64 `json:"checksum,omitempty"`

	// The requested nameserver
	Server *V1Nameserver `json:"server,omitempty"`
}

// Validate validates this v1 nameserver response
func (m *V1NameserverResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateServer(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1NameserverResponse) validateServer(formats strfmt.Registry) error {
	if swag.IsZero(m.Server) { // not required
		return nil
	}

	if m.Server != nil {
		if err := m.Server.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("server")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("server")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 nameserver response based on the context it is used
func (m *V1NameserverResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateServer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1NameserverResponse) contextValidateServer(ctx context.Context, formats strfmt.Registry) error {

	if m.Server != nil {

		if swag.IsZero(m.Server) { // not required
			return nil
		}

		if err := m.Server.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("server")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("server")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1NameserverResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1NameserverResponse) UnmarshalBinary(b []byte) error {
	var res V1NameserverResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}