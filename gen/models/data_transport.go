// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// DataTransport data transport
// swagger:model DataTransport
type DataTransport struct {

	// params
	Params map[string]string `json:"params,omitempty"`

	// transport type
	TransportType string `json:"transportType,omitempty"`
}

// Validate validates this data transport
func (m *DataTransport) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DataTransport) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DataTransport) UnmarshalBinary(b []byte) error {
	var res DataTransport
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
