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

package session

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openziti/edge-api/rest_model"
)

// DetailSessionReader is a Reader for the DetailSession structure.
type DetailSessionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DetailSessionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDetailSessionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDetailSessionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDetailSessionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDetailSessionOK creates a DetailSessionOK with default headers values
func NewDetailSessionOK() *DetailSessionOK {
	return &DetailSessionOK{}
}

/* DetailSessionOK describes a response with status code 200, with default header values.

A single session
*/
type DetailSessionOK struct {
	Payload *rest_model.DetailSessionEnvelope
}

func (o *DetailSessionOK) Error() string {
	return fmt.Sprintf("[GET /sessions/{id}][%d] detailSessionOK  %+v", 200, o.Payload)
}
func (o *DetailSessionOK) GetPayload() *rest_model.DetailSessionEnvelope {
	return o.Payload
}

func (o *DetailSessionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.DetailSessionEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDetailSessionUnauthorized creates a DetailSessionUnauthorized with default headers values
func NewDetailSessionUnauthorized() *DetailSessionUnauthorized {
	return &DetailSessionUnauthorized{}
}

/* DetailSessionUnauthorized describes a response with status code 401, with default header values.

The supplied session does not have the correct access rights to request this resource
*/
type DetailSessionUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DetailSessionUnauthorized) Error() string {
	return fmt.Sprintf("[GET /sessions/{id}][%d] detailSessionUnauthorized  %+v", 401, o.Payload)
}
func (o *DetailSessionUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DetailSessionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDetailSessionNotFound creates a DetailSessionNotFound with default headers values
func NewDetailSessionNotFound() *DetailSessionNotFound {
	return &DetailSessionNotFound{}
}

/* DetailSessionNotFound describes a response with status code 404, with default header values.

The requested resource does not exist
*/
type DetailSessionNotFound struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DetailSessionNotFound) Error() string {
	return fmt.Sprintf("[GET /sessions/{id}][%d] detailSessionNotFound  %+v", 404, o.Payload)
}
func (o *DetailSessionNotFound) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DetailSessionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
