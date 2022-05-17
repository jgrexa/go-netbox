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

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WritableDeviceWithConfigContext writable device with config context
//
// swagger:model WritableDeviceWithConfigContext
type WritableDeviceWithConfigContext struct {

	// Airflow
	// Enum: [front-to-rear rear-to-front left-to-right right-to-left side-to-rear passive]
	Airflow string `json:"airflow,omitempty"`

	// Asset tag
	//
	// A unique tag used to identify this device
	// Max Length: 50
	AssetTag *string `json:"asset_tag,omitempty"`

	// Cluster
	Cluster *int64 `json:"cluster,omitempty"`

	// Comments
	Comments string `json:"comments,omitempty"`

	// Config context
	// Read Only: true
	ConfigContext interface{} `json:"config_context,omitempty"`

	// Created
	// Read Only: true
	// Format: date
	Created strfmt.Date `json:"created,omitempty"`

	// Custom fields
	CustomFields interface{} `json:"custom_fields,omitempty"`

	// Device role
	// Required: true
	DeviceRole *int64 `json:"device_role"`

	// Device type
	// Required: true
	DeviceType *int64 `json:"device_type"`

	// Display
	// Read Only: true
	Display string `json:"display,omitempty"`

	// Rack face
	// Required: true
	// Enum: [front rear]
	Face *string `json:"face"`

	// Id
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Last updated
	// Read Only: true
	// Format: date-time
	LastUpdated strfmt.DateTime `json:"last_updated,omitempty"`

	// Local context data
	LocalContextData interface{} `json:"local_context_data,omitempty"`

	// Location
	Location *int64 `json:"location,omitempty"`

	// Name
	// Required: true
	// Max Length: 64
	Name *string `json:"name"`

	// parent device
	ParentDevice *NestedDevice `json:"parent_device,omitempty"`

	// Platform
	Platform *int64 `json:"platform,omitempty"`

	// Position (U)
	// Minimum: 1
	Position *int64 `json:"position,omitempty"`

	// Primary ip
	// Read Only: true
	PrimaryIP string `json:"primary_ip,omitempty"`

	// Primary IPv4
	PrimaryIp4 *int64 `json:"primary_ip4,omitempty"`

	// Primary IPv6
	PrimaryIp6 *int64 `json:"primary_ip6,omitempty"`

	// Rack
	// Required: true
	Rack *int64 `json:"rack"`

	// Serial number
	// Max Length: 50
	Serial string `json:"serial,omitempty"`

	// Site
	// Required: true
	Site *int64 `json:"site"`

	// Status
	// Enum: [offline active planned staged failed inventory decommissioning]
	Status string `json:"status,omitempty"`

	// tags
	Tags []*NestedTag `json:"tags"`

	// Tenant
	// Required: true
	Tenant *int64 `json:"tenant"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`

	// Vc position
	// Maximum: 255
	// Minimum: 0
	VcPosition *int64 `json:"vc_position,omitempty"`

	// Vc priority
	// Maximum: 255
	// Minimum: 0
	VcPriority *int64 `json:"vc_priority,omitempty"`

	// Virtual chassis
	// Required: true
	VirtualChassis *int64 `json:"virtual_chassis"`
}

// Validate validates this writable device with config context
func (m *WritableDeviceWithConfigContext) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAirflow(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAssetTag(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceRole(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParentDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePosition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRack(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSerial(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSite(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTenant(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVcPosition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVcPriority(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVirtualChassis(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var writableDeviceWithConfigContextTypeAirflowPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["front-to-rear","rear-to-front","left-to-right","right-to-left","side-to-rear","passive"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableDeviceWithConfigContextTypeAirflowPropEnum = append(writableDeviceWithConfigContextTypeAirflowPropEnum, v)
	}
}

const (

	// WritableDeviceWithConfigContextAirflowFrontDashToDashRear captures enum value "front-to-rear"
	WritableDeviceWithConfigContextAirflowFrontDashToDashRear string = "front-to-rear"

	// WritableDeviceWithConfigContextAirflowRearDashToDashFront captures enum value "rear-to-front"
	WritableDeviceWithConfigContextAirflowRearDashToDashFront string = "rear-to-front"

	// WritableDeviceWithConfigContextAirflowLeftDashToDashRight captures enum value "left-to-right"
	WritableDeviceWithConfigContextAirflowLeftDashToDashRight string = "left-to-right"

	// WritableDeviceWithConfigContextAirflowRightDashToDashLeft captures enum value "right-to-left"
	WritableDeviceWithConfigContextAirflowRightDashToDashLeft string = "right-to-left"

	// WritableDeviceWithConfigContextAirflowSideDashToDashRear captures enum value "side-to-rear"
	WritableDeviceWithConfigContextAirflowSideDashToDashRear string = "side-to-rear"

	// WritableDeviceWithConfigContextAirflowPassive captures enum value "passive"
	WritableDeviceWithConfigContextAirflowPassive string = "passive"
)

// prop value enum
func (m *WritableDeviceWithConfigContext) validateAirflowEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, writableDeviceWithConfigContextTypeAirflowPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WritableDeviceWithConfigContext) validateAirflow(formats strfmt.Registry) error {
	if swag.IsZero(m.Airflow) { // not required
		return nil
	}

	// value enum
	if err := m.validateAirflowEnum("airflow", "body", m.Airflow); err != nil {
		return err
	}

	return nil
}

func (m *WritableDeviceWithConfigContext) validateAssetTag(formats strfmt.Registry) error {
	if swag.IsZero(m.AssetTag) { // not required
		return nil
	}

	if err := validate.MaxLength("asset_tag", "body", *m.AssetTag, 50); err != nil {
		return err
	}

	return nil
}

func (m *WritableDeviceWithConfigContext) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableDeviceWithConfigContext) validateDeviceRole(formats strfmt.Registry) error {

	if err := validate.Required("device_role", "body", m.DeviceRole); err != nil {
		return err
	}

	return nil
}

func (m *WritableDeviceWithConfigContext) validateDeviceType(formats strfmt.Registry) error {

	if err := validate.Required("device_type", "body", m.DeviceType); err != nil {
		return err
	}

	return nil
}

var writableDeviceWithConfigContextTypeFacePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["front","rear"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableDeviceWithConfigContextTypeFacePropEnum = append(writableDeviceWithConfigContextTypeFacePropEnum, v)
	}
}

const (

	// WritableDeviceWithConfigContextFaceFront captures enum value "front"
	WritableDeviceWithConfigContextFaceFront string = "front"

	// WritableDeviceWithConfigContextFaceRear captures enum value "rear"
	WritableDeviceWithConfigContextFaceRear string = "rear"
)

// prop value enum
func (m *WritableDeviceWithConfigContext) validateFaceEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, writableDeviceWithConfigContextTypeFacePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WritableDeviceWithConfigContext) validateFace(formats strfmt.Registry) error {

	if err := validate.Required("face", "body", m.Face); err != nil {
		return err
	}

	// value enum
	if err := m.validateFaceEnum("face", "body", *m.Face); err != nil {
		return err
	}

	return nil
}

func (m *WritableDeviceWithConfigContext) validateLastUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableDeviceWithConfigContext) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 64); err != nil {
		return err
	}

	return nil
}

func (m *WritableDeviceWithConfigContext) validateParentDevice(formats strfmt.Registry) error {
	if swag.IsZero(m.ParentDevice) { // not required
		return nil
	}

	if m.ParentDevice != nil {
		if err := m.ParentDevice.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parent_device")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("parent_device")
			}
			return err
		}
	}

	return nil
}

func (m *WritableDeviceWithConfigContext) validatePosition(formats strfmt.Registry) error {
	if swag.IsZero(m.Position) { // not required
		return nil
	}

	if err := validate.MinimumInt("position", "body", *m.Position, 1, false); err != nil {
		return err
	}

	return nil
}

func (m *WritableDeviceWithConfigContext) validateRack(formats strfmt.Registry) error {

	if err := validate.Required("rack", "body", m.Rack); err != nil {
		return err
	}

	return nil
}

func (m *WritableDeviceWithConfigContext) validateSerial(formats strfmt.Registry) error {
	if swag.IsZero(m.Serial) { // not required
		return nil
	}

	if err := validate.MaxLength("serial", "body", m.Serial, 50); err != nil {
		return err
	}

	return nil
}

func (m *WritableDeviceWithConfigContext) validateSite(formats strfmt.Registry) error {

	if err := validate.Required("site", "body", m.Site); err != nil {
		return err
	}

	return nil
}

var writableDeviceWithConfigContextTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["offline","active","planned","staged","failed","inventory","decommissioning"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableDeviceWithConfigContextTypeStatusPropEnum = append(writableDeviceWithConfigContextTypeStatusPropEnum, v)
	}
}

const (

	// WritableDeviceWithConfigContextStatusOffline captures enum value "offline"
	WritableDeviceWithConfigContextStatusOffline string = "offline"

	// WritableDeviceWithConfigContextStatusActive captures enum value "active"
	WritableDeviceWithConfigContextStatusActive string = "active"

	// WritableDeviceWithConfigContextStatusPlanned captures enum value "planned"
	WritableDeviceWithConfigContextStatusPlanned string = "planned"

	// WritableDeviceWithConfigContextStatusStaged captures enum value "staged"
	WritableDeviceWithConfigContextStatusStaged string = "staged"

	// WritableDeviceWithConfigContextStatusFailed captures enum value "failed"
	WritableDeviceWithConfigContextStatusFailed string = "failed"

	// WritableDeviceWithConfigContextStatusInventory captures enum value "inventory"
	WritableDeviceWithConfigContextStatusInventory string = "inventory"

	// WritableDeviceWithConfigContextStatusDecommissioning captures enum value "decommissioning"
	WritableDeviceWithConfigContextStatusDecommissioning string = "decommissioning"
)

// prop value enum
func (m *WritableDeviceWithConfigContext) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, writableDeviceWithConfigContextTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WritableDeviceWithConfigContext) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *WritableDeviceWithConfigContext) validateTags(formats strfmt.Registry) error {
	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags); i++ {
		if swag.IsZero(m.Tags[i]) { // not required
			continue
		}

		if m.Tags[i] != nil {
			if err := m.Tags[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *WritableDeviceWithConfigContext) validateTenant(formats strfmt.Registry) error {

	if err := validate.Required("tenant", "body", m.Tenant); err != nil {
		return err
	}

	return nil
}

func (m *WritableDeviceWithConfigContext) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableDeviceWithConfigContext) validateVcPosition(formats strfmt.Registry) error {
	if swag.IsZero(m.VcPosition) { // not required
		return nil
	}

	if err := validate.MinimumInt("vc_position", "body", *m.VcPosition, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("vc_position", "body", *m.VcPosition, 255, false); err != nil {
		return err
	}

	return nil
}

func (m *WritableDeviceWithConfigContext) validateVcPriority(formats strfmt.Registry) error {
	if swag.IsZero(m.VcPriority) { // not required
		return nil
	}

	if err := validate.MinimumInt("vc_priority", "body", *m.VcPriority, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("vc_priority", "body", *m.VcPriority, 255, false); err != nil {
		return err
	}

	return nil
}

func (m *WritableDeviceWithConfigContext) validateVirtualChassis(formats strfmt.Registry) error {

	if err := validate.Required("virtual_chassis", "body", m.VirtualChassis); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this writable device with config context based on the context it is used
func (m *WritableDeviceWithConfigContext) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDisplay(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastUpdated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateParentDevice(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePrimaryIP(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WritableDeviceWithConfigContext) contextValidateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created", "body", strfmt.Date(m.Created)); err != nil {
		return err
	}

	return nil
}

func (m *WritableDeviceWithConfigContext) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *WritableDeviceWithConfigContext) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *WritableDeviceWithConfigContext) contextValidateLastUpdated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "last_updated", "body", strfmt.DateTime(m.LastUpdated)); err != nil {
		return err
	}

	return nil
}

func (m *WritableDeviceWithConfigContext) contextValidateParentDevice(ctx context.Context, formats strfmt.Registry) error {

	if m.ParentDevice != nil {
		if err := m.ParentDevice.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parent_device")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("parent_device")
			}
			return err
		}
	}

	return nil
}

func (m *WritableDeviceWithConfigContext) contextValidatePrimaryIP(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "primary_ip", "body", string(m.PrimaryIP)); err != nil {
		return err
	}

	return nil
}

func (m *WritableDeviceWithConfigContext) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Tags); i++ {

		if m.Tags[i] != nil {
			if err := m.Tags[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *WritableDeviceWithConfigContext) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WritableDeviceWithConfigContext) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WritableDeviceWithConfigContext) UnmarshalBinary(b []byte) error {
	var res WritableDeviceWithConfigContext
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
