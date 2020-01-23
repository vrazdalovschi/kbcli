// Code generated by go-swagger; DO NOT EDIT.

package bundle

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

// GetBundleAuditLogsWithHistoryReader is a Reader for the GetBundleAuditLogsWithHistory structure.
type GetBundleAuditLogsWithHistoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBundleAuditLogsWithHistoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetBundleAuditLogsWithHistoryOK()
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

// NewGetBundleAuditLogsWithHistoryOK creates a GetBundleAuditLogsWithHistoryOK with default headers values
func NewGetBundleAuditLogsWithHistoryOK() *GetBundleAuditLogsWithHistoryOK {
	return &GetBundleAuditLogsWithHistoryOK{}
}

/*GetBundleAuditLogsWithHistoryOK handles this case with default header values.

successful operation
*/
type GetBundleAuditLogsWithHistoryOK struct {
	Payload []*kbmodel.AuditLog

	HttpResponse runtime.ClientResponse
}

func (o *GetBundleAuditLogsWithHistoryOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/bundles/{bundleId}/auditLogsWithHistory][%d] getBundleAuditLogsWithHistoryOK  %+v", 200, o.Payload)
}

func (o *GetBundleAuditLogsWithHistoryOK) GetPayload() []*kbmodel.AuditLog {
	return o.Payload
}

func (o *GetBundleAuditLogsWithHistoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBundleAuditLogsWithHistoryNotFound creates a GetBundleAuditLogsWithHistoryNotFound with default headers values
func NewGetBundleAuditLogsWithHistoryNotFound() *GetBundleAuditLogsWithHistoryNotFound {
	return &GetBundleAuditLogsWithHistoryNotFound{}
}

/*GetBundleAuditLogsWithHistoryNotFound handles this case with default header values.

Subscription bundle not found
*/
type GetBundleAuditLogsWithHistoryNotFound struct {
	HttpResponse runtime.ClientResponse
}

func (o *GetBundleAuditLogsWithHistoryNotFound) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/bundles/{bundleId}/auditLogsWithHistory][%d] getBundleAuditLogsWithHistoryNotFound ", 404)
}

func (o *GetBundleAuditLogsWithHistoryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
