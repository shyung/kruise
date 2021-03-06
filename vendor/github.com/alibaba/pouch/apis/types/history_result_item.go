// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// HistoryResultItem An object containing image history at API side.
// swagger:model HistoryResultItem
type HistoryResultItem struct {

	// the author of the build point.
	// Required: true
	Author string `json:"Author"`

	// a custom message set when creating the layer.
	// Required: true
	Comment string `json:"Comment"`

	// the combined date and time at which the layer was created.
	// Required: true
	Created int64 `json:"Created"`

	// the command which created the layer.
	// Required: true
	CreatedBy string `json:"CreatedBy"`

	// mark whether the history item created a filesystem diff or not.
	// Required: true
	EmptyLayer bool `json:"EmptyLayer"`

	// ID of each layer image.
	// Required: true
	ID string `json:"ID"`

	// size of each layer image.
	// Required: true
	Size int64 `json:"Size"`
}

// Validate validates this history result item
func (m *HistoryResultItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmptyLayer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSize(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HistoryResultItem) validateAuthor(formats strfmt.Registry) error {

	if err := validate.RequiredString("Author", "body", string(m.Author)); err != nil {
		return err
	}

	return nil
}

func (m *HistoryResultItem) validateComment(formats strfmt.Registry) error {

	if err := validate.RequiredString("Comment", "body", string(m.Comment)); err != nil {
		return err
	}

	return nil
}

func (m *HistoryResultItem) validateCreated(formats strfmt.Registry) error {

	if err := validate.Required("Created", "body", int64(m.Created)); err != nil {
		return err
	}

	return nil
}

func (m *HistoryResultItem) validateCreatedBy(formats strfmt.Registry) error {

	if err := validate.RequiredString("CreatedBy", "body", string(m.CreatedBy)); err != nil {
		return err
	}

	return nil
}

func (m *HistoryResultItem) validateEmptyLayer(formats strfmt.Registry) error {

	if err := validate.Required("EmptyLayer", "body", bool(m.EmptyLayer)); err != nil {
		return err
	}

	return nil
}

func (m *HistoryResultItem) validateID(formats strfmt.Registry) error {

	if err := validate.RequiredString("ID", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *HistoryResultItem) validateSize(formats strfmt.Registry) error {

	if err := validate.Required("Size", "body", int64(m.Size)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HistoryResultItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HistoryResultItem) UnmarshalBinary(b []byte) error {
	var res HistoryResultItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
