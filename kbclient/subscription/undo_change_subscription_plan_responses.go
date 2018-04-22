// Code generated by go-swagger; DO NOT EDIT.

package subscription

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/killbill/kbcli/kbcommon"

	strfmt "github.com/go-openapi/strfmt"
)

// UndoChangeSubscriptionPlanReader is a Reader for the UndoChangeSubscriptionPlan structure.
type UndoChangeSubscriptionPlanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UndoChangeSubscriptionPlanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewUndoChangeSubscriptionPlanNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUndoChangeSubscriptionPlanBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUndoChangeSubscriptionPlanNotFound()
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

// NewUndoChangeSubscriptionPlanNoContent creates a UndoChangeSubscriptionPlanNoContent with default headers values
func NewUndoChangeSubscriptionPlanNoContent() *UndoChangeSubscriptionPlanNoContent {
	return &UndoChangeSubscriptionPlanNoContent{}
}

/*UndoChangeSubscriptionPlanNoContent handles this case with default header values.

Successful operation
*/
type UndoChangeSubscriptionPlanNoContent struct {
}

func (o *UndoChangeSubscriptionPlanNoContent) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/subscriptions/{subscriptionId}/undoChangePlan][%d] undoChangeSubscriptionPlanNoContent ", 204)
}

func (o *UndoChangeSubscriptionPlanNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUndoChangeSubscriptionPlanBadRequest creates a UndoChangeSubscriptionPlanBadRequest with default headers values
func NewUndoChangeSubscriptionPlanBadRequest() *UndoChangeSubscriptionPlanBadRequest {
	return &UndoChangeSubscriptionPlanBadRequest{}
}

/*UndoChangeSubscriptionPlanBadRequest handles this case with default header values.

Invalid subscription id supplied
*/
type UndoChangeSubscriptionPlanBadRequest struct {
}

func (o *UndoChangeSubscriptionPlanBadRequest) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/subscriptions/{subscriptionId}/undoChangePlan][%d] undoChangeSubscriptionPlanBadRequest ", 400)
}

func (o *UndoChangeSubscriptionPlanBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUndoChangeSubscriptionPlanNotFound creates a UndoChangeSubscriptionPlanNotFound with default headers values
func NewUndoChangeSubscriptionPlanNotFound() *UndoChangeSubscriptionPlanNotFound {
	return &UndoChangeSubscriptionPlanNotFound{}
}

/*UndoChangeSubscriptionPlanNotFound handles this case with default header values.

Entitlement not found
*/
type UndoChangeSubscriptionPlanNotFound struct {
}

func (o *UndoChangeSubscriptionPlanNotFound) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/subscriptions/{subscriptionId}/undoChangePlan][%d] undoChangeSubscriptionPlanNotFound ", 404)
}

func (o *UndoChangeSubscriptionPlanNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
