// Code generated by go-swagger; DO NOT EDIT.

package security

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

// GetCurrentUserSubjectReader is a Reader for the GetCurrentUserSubject structure.
type GetCurrentUserSubjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCurrentUserSubjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCurrentUserSubjectOK()
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

// NewGetCurrentUserSubjectOK creates a GetCurrentUserSubjectOK with default headers values
func NewGetCurrentUserSubjectOK() *GetCurrentUserSubjectOK {
	return &GetCurrentUserSubjectOK{}
}

/*GetCurrentUserSubjectOK handles this case with default header values.

successful operation
*/
type GetCurrentUserSubjectOK struct {
	Payload *kbmodel.Subject
}

func (o *GetCurrentUserSubjectOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/security/subject][%d] getCurrentUserSubjectOK  %+v", 200, o.Payload)
}

func (o *GetCurrentUserSubjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.Subject)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
