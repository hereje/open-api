// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// FunctionConfig function config
//
// swagger:model functionConfig
type FunctionConfig struct {

	// build data
	BuildData interface{} `json:"build_data,omitempty"`

	// display name
	DisplayName string `json:"display_name,omitempty"`

	// generator
	Generator string `json:"generator,omitempty"`

	// routes
	Routes []*FunctionRoute `json:"routes"`
}

// Validate validates this function config
func (m *FunctionConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRoutes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FunctionConfig) validateRoutes(formats strfmt.Registry) error {

	if swag.IsZero(m.Routes) { // not required
		return nil
	}

	for i := 0; i < len(m.Routes); i++ {
		if swag.IsZero(m.Routes[i]) { // not required
			continue
		}

		if m.Routes[i] != nil {
			if err := m.Routes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("routes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *FunctionConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FunctionConfig) UnmarshalBinary(b []byte) error {
	var res FunctionConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
