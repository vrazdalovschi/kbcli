// Code generated by go-swagger; DO NOT EDIT.

package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewTransferChildCreditToParentParams creates a new TransferChildCreditToParentParams object
// with the default values initialized.
func NewTransferChildCreditToParentParams() *TransferChildCreditToParentParams {
	var ()
	return &TransferChildCreditToParentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTransferChildCreditToParentParamsWithTimeout creates a new TransferChildCreditToParentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTransferChildCreditToParentParamsWithTimeout(timeout time.Duration) *TransferChildCreditToParentParams {
	var ()
	return &TransferChildCreditToParentParams{

		timeout: timeout,
	}
}

// NewTransferChildCreditToParentParamsWithContext creates a new TransferChildCreditToParentParams object
// with the default values initialized, and the ability to set a context for a request
func NewTransferChildCreditToParentParamsWithContext(ctx context.Context) *TransferChildCreditToParentParams {
	var ()
	return &TransferChildCreditToParentParams{

		Context: ctx,
	}
}

// NewTransferChildCreditToParentParamsWithHTTPClient creates a new TransferChildCreditToParentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTransferChildCreditToParentParamsWithHTTPClient(client *http.Client) *TransferChildCreditToParentParams {
	var ()
	return &TransferChildCreditToParentParams{
		HTTPClient: client,
	}
}

/*TransferChildCreditToParentParams contains all the parameters to send to the API endpoint
for the transfer child credit to parent operation typically these are written to a http.Request
*/
type TransferChildCreditToParentParams struct {

	/*XKillbillComment*/
	XKillbillComment *string
	/*XKillbillCreatedBy*/
	XKillbillCreatedBy string
	/*XKillbillReason*/
	XKillbillReason *string
	/*ChildAccountID*/
	ChildAccountID strfmt.UUID

	WithStackTrace        *bool // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the transfer child credit to parent params
func (o *TransferChildCreditToParentParams) WithTimeout(timeout time.Duration) *TransferChildCreditToParentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the transfer child credit to parent params
func (o *TransferChildCreditToParentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the transfer child credit to parent params
func (o *TransferChildCreditToParentParams) WithContext(ctx context.Context) *TransferChildCreditToParentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the transfer child credit to parent params
func (o *TransferChildCreditToParentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the transfer child credit to parent params
func (o *TransferChildCreditToParentParams) WithHTTPClient(client *http.Client) *TransferChildCreditToParentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the transfer child credit to parent params
func (o *TransferChildCreditToParentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the transfer child credit to parent params
func (o *TransferChildCreditToParentParams) WithXKillbillComment(xKillbillComment *string) *TransferChildCreditToParentParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the transfer child credit to parent params
func (o *TransferChildCreditToParentParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the transfer child credit to parent params
func (o *TransferChildCreditToParentParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *TransferChildCreditToParentParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the transfer child credit to parent params
func (o *TransferChildCreditToParentParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the transfer child credit to parent params
func (o *TransferChildCreditToParentParams) WithXKillbillReason(xKillbillReason *string) *TransferChildCreditToParentParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the transfer child credit to parent params
func (o *TransferChildCreditToParentParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithChildAccountID adds the childAccountID to the transfer child credit to parent params
func (o *TransferChildCreditToParentParams) WithChildAccountID(childAccountID strfmt.UUID) *TransferChildCreditToParentParams {
	o.SetChildAccountID(childAccountID)
	return o
}

// SetChildAccountID adds the childAccountId to the transfer child credit to parent params
func (o *TransferChildCreditToParentParams) SetChildAccountID(childAccountID strfmt.UUID) {
	o.ChildAccountID = childAccountID
}

// WriteToRequest writes these params to a swagger request
func (o *TransferChildCreditToParentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XKillbillComment != nil {

		// header param X-Killbill-Comment
		if err := r.SetHeaderParam("X-Killbill-Comment", *o.XKillbillComment); err != nil {
			return err
		}

	}

	// header param X-Killbill-CreatedBy
	if err := r.SetHeaderParam("X-Killbill-CreatedBy", o.XKillbillCreatedBy); err != nil {
		return err
	}

	if o.XKillbillReason != nil {

		// header param X-Killbill-Reason
		if err := r.SetHeaderParam("X-Killbill-Reason", *o.XKillbillReason); err != nil {
			return err
		}

	}

	// path param childAccountId
	if err := r.SetPathParam("childAccountId", o.ChildAccountID.String()); err != nil {
		return err
	}

	// header param withStackTrace
	if o.WithStackTrace != nil && *o.WithStackTrace {
		if err := r.SetQueryParam("withStackTrace", "true"); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
