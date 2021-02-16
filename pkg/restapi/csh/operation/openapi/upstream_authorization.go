// Code generated by go-swagger; DO NOT EDIT.

package openapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UpstreamAuthorization upstream authorization
//
// swagger:model UpstreamAuthorization
type UpstreamAuthorization struct {

	// base URL
	BaseURL string `json:"baseURL,omitempty"`

	// zcap
	Zcap string `json:"zcap,omitempty"`
}

// Validate validates this upstream authorization
func (m *UpstreamAuthorization) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this upstream authorization based on context it is used
func (m *UpstreamAuthorization) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpstreamAuthorization) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpstreamAuthorization) UnmarshalBinary(b []byte) error {
	var res UpstreamAuthorization
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
