// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2019 HAProxy Technologies
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package service_discovery

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/haproxytech/models/v2"
)

// GetConsulsHandlerFunc turns a function with the right signature into a get consuls handler
type GetConsulsHandlerFunc func(GetConsulsParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetConsulsHandlerFunc) Handle(params GetConsulsParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetConsulsHandler interface for that can handle valid get consuls params
type GetConsulsHandler interface {
	Handle(GetConsulsParams, interface{}) middleware.Responder
}

// NewGetConsuls creates a new http.Handler for the get consuls operation
func NewGetConsuls(ctx *middleware.Context, handler GetConsulsHandler) *GetConsuls {
	return &GetConsuls{Context: ctx, Handler: handler}
}

/*GetConsuls swagger:route GET /service_discovery/consul ServiceDiscovery getConsuls

Return an array of all configured Consul servers

Returns all configured Consul servers.

*/
type GetConsuls struct {
	Context *middleware.Context
	Handler GetConsulsHandler
}

func (o *GetConsuls) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetConsulsParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetConsulsOKBody get consuls o k body
//
// swagger:model GetConsulsOKBody
type GetConsulsOKBody struct {

	// data
	// Required: true
	Data models.Consuls `json:"data"`
}

// Validate validates this get consuls o k body
func (o *GetConsulsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetConsulsOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getConsulsOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if err := o.Data.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("getConsulsOK" + "." + "data")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetConsulsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetConsulsOKBody) UnmarshalBinary(b []byte) error {
	var res GetConsulsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
