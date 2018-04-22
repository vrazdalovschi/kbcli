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

	kbmodel "github.com/killbill/kbcli/kbmodel"
)

// GetInvoiceItemTagsReader is a Reader for the GetInvoiceItemTags structure.
type GetInvoiceItemTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInvoiceItemTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetInvoiceItemTagsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetInvoiceItemTagsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetInvoiceItemTagsNotFound()
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

// NewGetInvoiceItemTagsOK creates a GetInvoiceItemTagsOK with default headers values
func NewGetInvoiceItemTagsOK() *GetInvoiceItemTagsOK {
	return &GetInvoiceItemTagsOK{}
}

/*GetInvoiceItemTagsOK handles this case with default header values.

successful operation
*/
type GetInvoiceItemTagsOK struct {
	Payload []*kbmodel.Tag
}

func (o *GetInvoiceItemTagsOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/invoiceItems/{invoiceItemId}/tags][%d] getInvoiceItemTagsOK  %+v", 200, o.Payload)
}

func (o *GetInvoiceItemTagsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInvoiceItemTagsBadRequest creates a GetInvoiceItemTagsBadRequest with default headers values
func NewGetInvoiceItemTagsBadRequest() *GetInvoiceItemTagsBadRequest {
	return &GetInvoiceItemTagsBadRequest{}
}

/*GetInvoiceItemTagsBadRequest handles this case with default header values.

Invalid invoice item id supplied
*/
type GetInvoiceItemTagsBadRequest struct {
}

func (o *GetInvoiceItemTagsBadRequest) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/invoiceItems/{invoiceItemId}/tags][%d] getInvoiceItemTagsBadRequest ", 400)
}

func (o *GetInvoiceItemTagsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInvoiceItemTagsNotFound creates a GetInvoiceItemTagsNotFound with default headers values
func NewGetInvoiceItemTagsNotFound() *GetInvoiceItemTagsNotFound {
	return &GetInvoiceItemTagsNotFound{}
}

/*GetInvoiceItemTagsNotFound handles this case with default header values.

Account not found
*/
type GetInvoiceItemTagsNotFound struct {
}

func (o *GetInvoiceItemTagsNotFound) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/invoiceItems/{invoiceItemId}/tags][%d] getInvoiceItemTagsNotFound ", 404)
}

func (o *GetInvoiceItemTagsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
