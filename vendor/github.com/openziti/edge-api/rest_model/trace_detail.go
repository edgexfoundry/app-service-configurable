// Code generated by go-swagger; DO NOT EDIT.

//
// Copyright NetFoundry Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// __          __              _
// \ \        / /             (_)
//  \ \  /\  / /_ _ _ __ _ __  _ _ __   __ _
//   \ \/  \/ / _` | '__| '_ \| | '_ \ / _` |
//    \  /\  / (_| | |  | | | | | | | | (_| | : This file is generated, do not edit it.
//     \/  \/ \__,_|_|  |_| |_|_|_| |_|\__, |
//                                      __/ |
//                                     |___/

package rest_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TraceDetail trace detail
//
// swagger:model traceDetail
type TraceDetail struct {

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// trace Id
	TraceID string `json:"traceId,omitempty"`

	// until
	// Format: date-time
	Until strfmt.DateTime `json:"until,omitempty"`
}

// Validate validates this trace detail
func (m *TraceDetail) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUntil(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TraceDetail) validateUntil(formats strfmt.Registry) error {
	if swag.IsZero(m.Until) { // not required
		return nil
	}

	if err := validate.FormatOf("until", "body", "date-time", m.Until.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this trace detail based on context it is used
func (m *TraceDetail) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TraceDetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TraceDetail) UnmarshalBinary(b []byte) error {
	var res TraceDetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
