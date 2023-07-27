// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// BGPNeigh b g p neigh
//
// swagger:model BGPNeigh
type BGPNeigh struct {

	// BGP Nieghbor IP address
	IPAddress string `json:"ipAddress,omitempty"`

	// Remote AS number
	RemoteAs int64 `json:"remoteAs,omitempty"`
}

// Validate validates this b g p neigh
func (m *BGPNeigh) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this b g p neigh based on context it is used
func (m *BGPNeigh) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BGPNeigh) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BGPNeigh) UnmarshalBinary(b []byte) error {
	var res BGPNeigh
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
