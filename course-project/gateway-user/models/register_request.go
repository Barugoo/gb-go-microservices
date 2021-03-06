// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RegisterRequest register request
//
// swagger:model RegisterRequest
type RegisterRequest struct {

	// age
	Age int64 `json:"Age,omitempty"`

	// display name
	DisplayName string `json:"DisplayName,omitempty"`

	// email
	Email string `json:"Email,omitempty"`

	// password
	Password string `json:"Password,omitempty"`

	// phone
	Phone string `json:"Phone,omitempty"`
}

// Validate validates this register request
func (m *RegisterRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RegisterRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RegisterRequest) UnmarshalBinary(b []byte) error {
	var res RegisterRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
