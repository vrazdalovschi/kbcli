// Code generated by go-swagger; DO NOT EDIT.

package invoice

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

// GetInvoiceByNumberReader is a Reader for the GetInvoiceByNumber structure.
type GetInvoiceByNumberReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInvoiceByNumberReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetInvoiceByNumberOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetInvoiceByNumberNotFound()
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

// NewGetInvoiceByNumberOK creates a GetInvoiceByNumberOK with default headers values
func NewGetInvoiceByNumberOK() *GetInvoiceByNumberOK {
	return &GetInvoiceByNumberOK{}
}

/*GetInvoiceByNumberOK handles this case with default header values.

successful operation
*/
type GetInvoiceByNumberOK struct {
	Payload *kbmodel.Invoice
}

func (o *GetInvoiceByNumberOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/invoices/{invoiceNumber}][%d] getInvoiceByNumberOK  %+v", 200, o.Payload)
}

func (o *GetInvoiceByNumberOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.Invoice)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInvoiceByNumberNotFound creates a GetInvoiceByNumberNotFound with default headers values
func NewGetInvoiceByNumberNotFound() *GetInvoiceByNumberNotFound {
	return &GetInvoiceByNumberNotFound{}
}

/*GetInvoiceByNumberNotFound handles this case with default header values.

Invoice not found
*/
type GetInvoiceByNumberNotFound struct {
}

func (o *GetInvoiceByNumberNotFound) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/invoices/{invoiceNumber}][%d] getInvoiceByNumberNotFound ", 404)
}

func (o *GetInvoiceByNumberNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
