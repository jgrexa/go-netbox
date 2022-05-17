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

package wireless

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// WirelessWirelessLanGroupsDeleteReader is a Reader for the WirelessWirelessLanGroupsDelete structure.
type WirelessWirelessLanGroupsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WirelessWirelessLanGroupsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewWirelessWirelessLanGroupsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewWirelessWirelessLanGroupsDeleteNoContent creates a WirelessWirelessLanGroupsDeleteNoContent with default headers values
func NewWirelessWirelessLanGroupsDeleteNoContent() *WirelessWirelessLanGroupsDeleteNoContent {
	return &WirelessWirelessLanGroupsDeleteNoContent{}
}

/* WirelessWirelessLanGroupsDeleteNoContent describes a response with status code 204, with default header values.

WirelessWirelessLanGroupsDeleteNoContent wireless wireless lan groups delete no content
*/
type WirelessWirelessLanGroupsDeleteNoContent struct {
}

func (o *WirelessWirelessLanGroupsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /wireless/wireless-lan-groups/{id}/][%d] wirelessWirelessLanGroupsDeleteNoContent ", 204)
}

func (o *WirelessWirelessLanGroupsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
