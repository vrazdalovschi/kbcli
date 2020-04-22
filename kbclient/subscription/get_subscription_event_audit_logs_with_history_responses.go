// Code generated by go-swagger; DO NOT EDIT.

package subscription

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/killbill/kbcli/v2/kbcommon"

	strfmt "github.com/go-openapi/strfmt"

	kbmodel "github.com/killbill/kbcli/v2/kbmodel"
)

// GetSubscriptionEventAuditLogsWithHistoryReader is a Reader for the GetSubscriptionEventAuditLogsWithHistory structure.
type GetSubscriptionEventAuditLogsWithHistoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSubscriptionEventAuditLogsWithHistoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSubscriptionEventAuditLogsWithHistoryOK()
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

// NewGetSubscriptionEventAuditLogsWithHistoryOK creates a GetSubscriptionEventAuditLogsWithHistoryOK with default headers values
func NewGetSubscriptionEventAuditLogsWithHistoryOK() *GetSubscriptionEventAuditLogsWithHistoryOK {
	return &GetSubscriptionEventAuditLogsWithHistoryOK{}
}

/*GetSubscriptionEventAuditLogsWithHistoryOK handles this case with default header values.

successful operation
*/
type GetSubscriptionEventAuditLogsWithHistoryOK struct {
	Payload []*kbmodel.AuditLog

	HttpResponse runtime.ClientResponse
}

func (o *GetSubscriptionEventAuditLogsWithHistoryOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/subscriptions/events/{eventId}/auditLogsWithHistory][%d] getSubscriptionEventAuditLogsWithHistoryOK  %+v", 200, o.Payload)
}

func (o *GetSubscriptionEventAuditLogsWithHistoryOK) GetPayload() []*kbmodel.AuditLog {
	return o.Payload
}

func (o *GetSubscriptionEventAuditLogsWithHistoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSubscriptionEventAuditLogsWithHistoryNotFound creates a GetSubscriptionEventAuditLogsWithHistoryNotFound with default headers values
func NewGetSubscriptionEventAuditLogsWithHistoryNotFound() *GetSubscriptionEventAuditLogsWithHistoryNotFound {
	return &GetSubscriptionEventAuditLogsWithHistoryNotFound{}
}

/*GetSubscriptionEventAuditLogsWithHistoryNotFound handles this case with default header values.

Subscription event not found
*/
type GetSubscriptionEventAuditLogsWithHistoryNotFound struct {
	HttpResponse runtime.ClientResponse
}

func (o *GetSubscriptionEventAuditLogsWithHistoryNotFound) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/subscriptions/events/{eventId}/auditLogsWithHistory][%d] getSubscriptionEventAuditLogsWithHistoryNotFound ", 404)
}

func (o *GetSubscriptionEventAuditLogsWithHistoryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
