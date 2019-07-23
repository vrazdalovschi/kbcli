// Code generated by go-swagger; DO NOT EDIT.

package kbmodel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TagDefinition tag definition
// swagger:model TagDefinition
type TagDefinition struct {

	// applicable object types
	// Unique: true
	ApplicableObjectTypes []TagDefinitionApplicableObjectTypesEnum `json:"applicableObjectTypes"`

	// audit logs
	AuditLogs []*AuditLog `json:"auditLogs"`

	// description
	// Required: true
	Description *string `json:"description"`

	// id
	ID strfmt.UUID `json:"id,omitempty"`

	// is control tag
	IsControlTag *bool `json:"isControlTag,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this tag definition
func (m *TagDefinition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApplicableObjectTypes(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateAuditLogs(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

type TagDefinitionApplicableObjectTypesEnum string

const (

	// TagDefinitionApplicableObjectTypesACCOUNT captures enum value "ACCOUNT"
	TagDefinitionApplicableObjectTypesACCOUNT TagDefinitionApplicableObjectTypesEnum = "ACCOUNT"

	// TagDefinitionApplicableObjectTypesACCOUNTEMAIL captures enum value "ACCOUNT_EMAIL"
	TagDefinitionApplicableObjectTypesACCOUNTEMAIL TagDefinitionApplicableObjectTypesEnum = "ACCOUNT_EMAIL"

	// TagDefinitionApplicableObjectTypesBLOCKINGSTATES captures enum value "BLOCKING_STATES"
	TagDefinitionApplicableObjectTypesBLOCKINGSTATES TagDefinitionApplicableObjectTypesEnum = "BLOCKING_STATES"

	// TagDefinitionApplicableObjectTypesBUNDLE captures enum value "BUNDLE"
	TagDefinitionApplicableObjectTypesBUNDLE TagDefinitionApplicableObjectTypesEnum = "BUNDLE"

	// TagDefinitionApplicableObjectTypesCUSTOMFIELD captures enum value "CUSTOM_FIELD"
	TagDefinitionApplicableObjectTypesCUSTOMFIELD TagDefinitionApplicableObjectTypesEnum = "CUSTOM_FIELD"

	// TagDefinitionApplicableObjectTypesINVOICE captures enum value "INVOICE"
	TagDefinitionApplicableObjectTypesINVOICE TagDefinitionApplicableObjectTypesEnum = "INVOICE"

	// TagDefinitionApplicableObjectTypesPAYMENT captures enum value "PAYMENT"
	TagDefinitionApplicableObjectTypesPAYMENT TagDefinitionApplicableObjectTypesEnum = "PAYMENT"

	// TagDefinitionApplicableObjectTypesTRANSACTION captures enum value "TRANSACTION"
	TagDefinitionApplicableObjectTypesTRANSACTION TagDefinitionApplicableObjectTypesEnum = "TRANSACTION"

	// TagDefinitionApplicableObjectTypesINVOICEITEM captures enum value "INVOICE_ITEM"
	TagDefinitionApplicableObjectTypesINVOICEITEM TagDefinitionApplicableObjectTypesEnum = "INVOICE_ITEM"

	// TagDefinitionApplicableObjectTypesINVOICEPAYMENT captures enum value "INVOICE_PAYMENT"
	TagDefinitionApplicableObjectTypesINVOICEPAYMENT TagDefinitionApplicableObjectTypesEnum = "INVOICE_PAYMENT"

	// TagDefinitionApplicableObjectTypesSUBSCRIPTION captures enum value "SUBSCRIPTION"
	TagDefinitionApplicableObjectTypesSUBSCRIPTION TagDefinitionApplicableObjectTypesEnum = "SUBSCRIPTION"

	// TagDefinitionApplicableObjectTypesSUBSCRIPTIONEVENT captures enum value "SUBSCRIPTION_EVENT"
	TagDefinitionApplicableObjectTypesSUBSCRIPTIONEVENT TagDefinitionApplicableObjectTypesEnum = "SUBSCRIPTION_EVENT"

	// TagDefinitionApplicableObjectTypesSERVICEBROADCAST captures enum value "SERVICE_BROADCAST"
	TagDefinitionApplicableObjectTypesSERVICEBROADCAST TagDefinitionApplicableObjectTypesEnum = "SERVICE_BROADCAST"

	// TagDefinitionApplicableObjectTypesPAYMENTATTEMPT captures enum value "PAYMENT_ATTEMPT"
	TagDefinitionApplicableObjectTypesPAYMENTATTEMPT TagDefinitionApplicableObjectTypesEnum = "PAYMENT_ATTEMPT"

	// TagDefinitionApplicableObjectTypesPAYMENTMETHOD captures enum value "PAYMENT_METHOD"
	TagDefinitionApplicableObjectTypesPAYMENTMETHOD TagDefinitionApplicableObjectTypesEnum = "PAYMENT_METHOD"

	// TagDefinitionApplicableObjectTypesTAG captures enum value "TAG"
	TagDefinitionApplicableObjectTypesTAG TagDefinitionApplicableObjectTypesEnum = "TAG"

	// TagDefinitionApplicableObjectTypesTAGDEFINITION captures enum value "TAG_DEFINITION"
	TagDefinitionApplicableObjectTypesTAGDEFINITION TagDefinitionApplicableObjectTypesEnum = "TAG_DEFINITION"

	// TagDefinitionApplicableObjectTypesTENANT captures enum value "TENANT"
	TagDefinitionApplicableObjectTypesTENANT TagDefinitionApplicableObjectTypesEnum = "TENANT"

	// TagDefinitionApplicableObjectTypesTENANTKVS captures enum value "TENANT_KVS"
	TagDefinitionApplicableObjectTypesTENANTKVS TagDefinitionApplicableObjectTypesEnum = "TENANT_KVS"
)

var TagDefinitionApplicableObjectTypesEnumValues = []string{
	"ACCOUNT",
	"ACCOUNT_EMAIL",
	"BLOCKING_STATES",
	"BUNDLE",
	"CUSTOM_FIELD",
	"INVOICE",
	"PAYMENT",
	"TRANSACTION",
	"INVOICE_ITEM",
	"INVOICE_PAYMENT",
	"SUBSCRIPTION",
	"SUBSCRIPTION_EVENT",
	"SERVICE_BROADCAST",
	"PAYMENT_ATTEMPT",
	"PAYMENT_METHOD",
	"TAG",
	"TAG_DEFINITION",
	"TENANT",
	"TENANT_KVS",
}

func (e TagDefinitionApplicableObjectTypesEnum) IsValid() bool {
	for _, v := range TagDefinitionApplicableObjectTypesEnumValues {
		if v == string(e) {
			return true
		}
	}
	return false
}

var tagDefinitionApplicableObjectTypesItemsEnum []interface{}

func init() {
	var res []TagDefinitionApplicableObjectTypesEnum
	if err := json.Unmarshal([]byte(`["ACCOUNT","ACCOUNT_EMAIL","BLOCKING_STATES","BUNDLE","CUSTOM_FIELD","INVOICE","PAYMENT","TRANSACTION","INVOICE_ITEM","INVOICE_PAYMENT","SUBSCRIPTION","SUBSCRIPTION_EVENT","SERVICE_BROADCAST","PAYMENT_ATTEMPT","PAYMENT_METHOD","TAG","TAG_DEFINITION","TENANT","TENANT_KVS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tagDefinitionApplicableObjectTypesItemsEnum = append(tagDefinitionApplicableObjectTypesItemsEnum, v)
	}
}

func (m *TagDefinition) validateApplicableObjectTypesItemsEnum(path, location string, value TagDefinitionApplicableObjectTypesEnum) error {
	if err := validate.Enum(path, location, value, tagDefinitionApplicableObjectTypesItemsEnum); err != nil {
		return err
	}
	return nil
}

func (m *TagDefinition) validateApplicableObjectTypes(formats strfmt.Registry) error {

	if swag.IsZero(m.ApplicableObjectTypes) { // not required
		return nil
	}

	if err := validate.UniqueItems("applicableObjectTypes", "body", m.ApplicableObjectTypes); err != nil {
		return err
	}

	for i := 0; i < len(m.ApplicableObjectTypes); i++ {

		// value enum
		if err := m.validateApplicableObjectTypesItemsEnum("applicableObjectTypes"+"."+strconv.Itoa(i), "body", m.ApplicableObjectTypes[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *TagDefinition) validateAuditLogs(formats strfmt.Registry) error {

	if swag.IsZero(m.AuditLogs) { // not required
		return nil
	}

	for i := 0; i < len(m.AuditLogs); i++ {

		if swag.IsZero(m.AuditLogs[i]) { // not required
			continue
		}

		if m.AuditLogs[i] != nil {

			if err := m.AuditLogs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("auditLogs" + "." + strconv.Itoa(i))
				}
				return err
			}

		}

	}

	return nil
}

func (m *TagDefinition) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *TagDefinition) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TagDefinition) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TagDefinition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TagDefinition) UnmarshalBinary(b []byte) error {
	var res TagDefinition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}