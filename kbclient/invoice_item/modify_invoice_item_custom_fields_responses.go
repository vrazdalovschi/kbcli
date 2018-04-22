// Code generated by go-swagger; DO NOT EDIT.

package invoice_item

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/killbill/kbcli/kbcommon"

	strfmt "github.com/go-openapi/strfmt"
)

// ModifyInvoiceItemCustomFieldsReader is a Reader for the ModifyInvoiceItemCustomFields structure.
type ModifyInvoiceItemCustomFieldsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ModifyInvoiceItemCustomFieldsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewModifyInvoiceItemCustomFieldsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewModifyInvoiceItemCustomFieldsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		errorResult := kbcommon.NewKillbillError(response.Code())
		if err := consumer.Consume(response.Body(), &errorResult); err != nil && err != io.EOF {
			return nil, err
		}
		return nil, errorResult
	}
}

// NewModifyInvoiceItemCustomFieldsNoContent creates a ModifyInvoiceItemCustomFieldsNoContent with default headers values
func NewModifyInvoiceItemCustomFieldsNoContent() *ModifyInvoiceItemCustomFieldsNoContent {
	return &ModifyInvoiceItemCustomFieldsNoContent{}
}

/*ModifyInvoiceItemCustomFieldsNoContent handles this case with default header values.

Successful operation
*/
type ModifyInvoiceItemCustomFieldsNoContent struct {
}

func (o *ModifyInvoiceItemCustomFieldsNoContent) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/invoiceItems/{invoiceItemId}/customFields][%d] modifyInvoiceItemCustomFieldsNoContent ", 204)
}

func (o *ModifyInvoiceItemCustomFieldsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewModifyInvoiceItemCustomFieldsBadRequest creates a ModifyInvoiceItemCustomFieldsBadRequest with default headers values
func NewModifyInvoiceItemCustomFieldsBadRequest() *ModifyInvoiceItemCustomFieldsBadRequest {
	return &ModifyInvoiceItemCustomFieldsBadRequest{}
}

/*ModifyInvoiceItemCustomFieldsBadRequest handles this case with default header values.

Invalid invoice item id supplied
*/
type ModifyInvoiceItemCustomFieldsBadRequest struct {
}

func (o *ModifyInvoiceItemCustomFieldsBadRequest) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/invoiceItems/{invoiceItemId}/customFields][%d] modifyInvoiceItemCustomFieldsBadRequest ", 400)
}

func (o *ModifyInvoiceItemCustomFieldsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
