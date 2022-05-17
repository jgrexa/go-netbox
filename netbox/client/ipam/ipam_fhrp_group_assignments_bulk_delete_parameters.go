// Code generated by go-swagger; DO NOT EDIT.

// Copyright (c) 2020 Samuel Mutel <12967891+smutel@users.noreply.github.com>
//
// Permission to use, copy, modify, and distribute this software for any
// purpose with or without fee is hereby granted, provided that the above
// copyright notice and this permission notice appear in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
// WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
// MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
// ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
// WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
// ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
// OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
//

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewIpamFhrpGroupAssignmentsBulkDeleteParams creates a new IpamFhrpGroupAssignmentsBulkDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamFhrpGroupAssignmentsBulkDeleteParams() *IpamFhrpGroupAssignmentsBulkDeleteParams {
	return &IpamFhrpGroupAssignmentsBulkDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamFhrpGroupAssignmentsBulkDeleteParamsWithTimeout creates a new IpamFhrpGroupAssignmentsBulkDeleteParams object
// with the ability to set a timeout on a request.
func NewIpamFhrpGroupAssignmentsBulkDeleteParamsWithTimeout(timeout time.Duration) *IpamFhrpGroupAssignmentsBulkDeleteParams {
	return &IpamFhrpGroupAssignmentsBulkDeleteParams{
		timeout: timeout,
	}
}

// NewIpamFhrpGroupAssignmentsBulkDeleteParamsWithContext creates a new IpamFhrpGroupAssignmentsBulkDeleteParams object
// with the ability to set a context for a request.
func NewIpamFhrpGroupAssignmentsBulkDeleteParamsWithContext(ctx context.Context) *IpamFhrpGroupAssignmentsBulkDeleteParams {
	return &IpamFhrpGroupAssignmentsBulkDeleteParams{
		Context: ctx,
	}
}

// NewIpamFhrpGroupAssignmentsBulkDeleteParamsWithHTTPClient creates a new IpamFhrpGroupAssignmentsBulkDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamFhrpGroupAssignmentsBulkDeleteParamsWithHTTPClient(client *http.Client) *IpamFhrpGroupAssignmentsBulkDeleteParams {
	return &IpamFhrpGroupAssignmentsBulkDeleteParams{
		HTTPClient: client,
	}
}

/* IpamFhrpGroupAssignmentsBulkDeleteParams contains all the parameters to send to the API endpoint
   for the ipam fhrp group assignments bulk delete operation.

   Typically these are written to a http.Request.
*/
type IpamFhrpGroupAssignmentsBulkDeleteParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam fhrp group assignments bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamFhrpGroupAssignmentsBulkDeleteParams) WithDefaults() *IpamFhrpGroupAssignmentsBulkDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam fhrp group assignments bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamFhrpGroupAssignmentsBulkDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam fhrp group assignments bulk delete params
func (o *IpamFhrpGroupAssignmentsBulkDeleteParams) WithTimeout(timeout time.Duration) *IpamFhrpGroupAssignmentsBulkDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam fhrp group assignments bulk delete params
func (o *IpamFhrpGroupAssignmentsBulkDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam fhrp group assignments bulk delete params
func (o *IpamFhrpGroupAssignmentsBulkDeleteParams) WithContext(ctx context.Context) *IpamFhrpGroupAssignmentsBulkDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam fhrp group assignments bulk delete params
func (o *IpamFhrpGroupAssignmentsBulkDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam fhrp group assignments bulk delete params
func (o *IpamFhrpGroupAssignmentsBulkDeleteParams) WithHTTPClient(client *http.Client) *IpamFhrpGroupAssignmentsBulkDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam fhrp group assignments bulk delete params
func (o *IpamFhrpGroupAssignmentsBulkDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *IpamFhrpGroupAssignmentsBulkDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
