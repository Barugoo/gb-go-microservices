// Code generated by go-swagger; DO NOT EDIT.

package movies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewRentMovieParams creates a new RentMovieParams object
// no default values defined in spec.
func NewRentMovieParams() RentMovieParams {

	return RentMovieParams{}
}

// RentMovieParams contains all the bound params for the rent movie operation
// typically these are obtained from a http.Request
//
// swagger:parameters RentMovie
type RentMovieParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: path
	*/
	MovieID int32
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewRentMovieParams() beforehand.
func (o *RentMovieParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rMovieID, rhkMovieID, _ := route.Params.GetOK("movieId")
	if err := o.bindMovieID(rMovieID, rhkMovieID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindMovieID binds and validates parameter MovieID from path.
func (o *RentMovieParams) bindMovieID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt32(raw)
	if err != nil {
		return errors.InvalidType("movieId", "path", "int32", raw)
	}
	o.MovieID = value

	return nil
}