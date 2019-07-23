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

// DeleteInvoiceItemTagsReader is a Reader for the DeleteInvoiceItemTags structure.
type DeleteInvoiceItemTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteInvoiceItemTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteInvoiceItemTagsNoContent()
		result.HttpResponse = response
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		errorResult := kbcommon.NewKillbillError(response.Code())
		if err := consumer.Consume(response.Body(), &errorResult); err != nil && err != io.EOF {
			return nil, err
		}
		return nil, errorResult
	}
}

// NewDeleteInvoiceItemTagsNoContent creates a DeleteInvoiceItemTagsNoContent with default headers values
func NewDeleteInvoiceItemTagsNoContent() *DeleteInvoiceItemTagsNoContent {
	return &DeleteInvoiceItemTagsNoContent{}
}

/*DeleteInvoiceItemTagsNoContent handles this case with default header values.

Successful operation
*/
type DeleteInvoiceItemTagsNoContent struct {
	HttpResponse runtime.ClientResponse
}

func (o *DeleteInvoiceItemTagsNoContent) Error() string {
	return fmt.Sprintf("[DELETE /1.0/kb/invoiceItems/{invoiceItemId}/tags][%d] deleteInvoiceItemTagsNoContent ", 204)
}

func (o *DeleteInvoiceItemTagsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteInvoiceItemTagsBadRequest creates a DeleteInvoiceItemTagsBadRequest with default headers values
func NewDeleteInvoiceItemTagsBadRequest() *DeleteInvoiceItemTagsBadRequest {
	return &DeleteInvoiceItemTagsBadRequest{}
}

/*DeleteInvoiceItemTagsBadRequest handles this case with default header values.

Invalid invoice item id supplied
*/
type DeleteInvoiceItemTagsBadRequest struct {
	HttpResponse runtime.ClientResponse
}

func (o *DeleteInvoiceItemTagsBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /1.0/kb/invoiceItems/{invoiceItemId}/tags][%d] deleteInvoiceItemTagsBadRequest ", 400)
}

func (o *DeleteInvoiceItemTagsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}