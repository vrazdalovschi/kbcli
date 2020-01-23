// Code generated by go-swagger; DO NOT EDIT.

package kbmodel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Overdue overdue
// swagger:model Overdue
type Overdue struct {

	// initial reevaluation interval
	InitialReevaluationInterval int32 `json:"initialReevaluationInterval,omitempty"`

	// overdue states
	OverdueStates []*OverdueStateConfig `json:"overdueStates"`
}

// Validate validates this overdue
func (m *Overdue) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOverdueStates(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Overdue) validateOverdueStates(formats strfmt.Registry) error {

	if swag.IsZero(m.OverdueStates) { // not required
		return nil
	}

	for i := 0; i < len(m.OverdueStates); i++ {
		if swag.IsZero(m.OverdueStates[i]) { // not required
			continue
		}

		if m.OverdueStates[i] != nil {
			if err := m.OverdueStates[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("overdueStates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Overdue) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Overdue) UnmarshalBinary(b []byte) error {
	var res Overdue
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
