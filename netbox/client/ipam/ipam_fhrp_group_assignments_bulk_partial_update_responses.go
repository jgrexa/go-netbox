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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smutel/go-netbox/netbox/models"
)

// IpamFhrpGroupAssignmentsBulkPartialUpdateReader is a Reader for the IpamFhrpGroupAssignmentsBulkPartialUpdate structure.
type IpamFhrpGroupAssignmentsBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamFhrpGroupAssignmentsBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamFhrpGroupAssignmentsBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIpamFhrpGroupAssignmentsBulkPartialUpdateOK creates a IpamFhrpGroupAssignmentsBulkPartialUpdateOK with default headers values
func NewIpamFhrpGroupAssignmentsBulkPartialUpdateOK() *IpamFhrpGroupAssignmentsBulkPartialUpdateOK {
	return &IpamFhrpGroupAssignmentsBulkPartialUpdateOK{}
}

/* IpamFhrpGroupAssignmentsBulkPartialUpdateOK describes a response with status code 200, with default header values.

IpamFhrpGroupAssignmentsBulkPartialUpdateOK ipam fhrp group assignments bulk partial update o k
*/
type IpamFhrpGroupAssignmentsBulkPartialUpdateOK struct {
	Payload *models.FHRPGroupAssignment
}

func (o *IpamFhrpGroupAssignmentsBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /ipam/fhrp-group-assignments/][%d] ipamFhrpGroupAssignmentsBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *IpamFhrpGroupAssignmentsBulkPartialUpdateOK) GetPayload() *models.FHRPGroupAssignment {
	return o.Payload
}

func (o *IpamFhrpGroupAssignmentsBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FHRPGroupAssignment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
