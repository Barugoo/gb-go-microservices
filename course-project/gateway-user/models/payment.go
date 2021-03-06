// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Payment payment
//
// swagger:model Payment
type Payment struct {

	// amount
	Amount float64 `json:"Amount,omitempty"`

	// created at
	CreatedAt int64 `json:"CreatedAt,omitempty"`

	// ID
	ID int32 `json:"ID,omitempty"`

	// status
	Status int64 `json:"Status,omitempty"`

	// transaction ID
	TransactionID int32 `json:"TransactionID,omitempty"`

	// updated at
	UpdatedAt int64 `json:"UpdatedAt,omitempty"`
}

// Validate validates this payment
func (m *Payment) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Payment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Payment) UnmarshalBinary(b []byte) error {
	var res Payment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
