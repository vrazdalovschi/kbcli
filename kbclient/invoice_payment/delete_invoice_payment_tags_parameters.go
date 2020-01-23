// Code generated by go-swagger; DO NOT EDIT.

package invoice_payment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteInvoicePaymentTagsParams creates a new DeleteInvoicePaymentTagsParams object
// with the default values initialized.
func NewDeleteInvoicePaymentTagsParams() *DeleteInvoicePaymentTagsParams {
	var ()
	return &DeleteInvoicePaymentTagsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteInvoicePaymentTagsParamsWithTimeout creates a new DeleteInvoicePaymentTagsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteInvoicePaymentTagsParamsWithTimeout(timeout time.Duration) *DeleteInvoicePaymentTagsParams {
	var ()
	return &DeleteInvoicePaymentTagsParams{

		timeout: timeout,
	}
}

// NewDeleteInvoicePaymentTagsParamsWithContext creates a new DeleteInvoicePaymentTagsParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteInvoicePaymentTagsParamsWithContext(ctx context.Context) *DeleteInvoicePaymentTagsParams {
	var ()
	return &DeleteInvoicePaymentTagsParams{

		Context: ctx,
	}
}

// NewDeleteInvoicePaymentTagsParamsWithHTTPClient creates a new DeleteInvoicePaymentTagsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteInvoicePaymentTagsParamsWithHTTPClient(client *http.Client) *DeleteInvoicePaymentTagsParams {
	var ()
	return &DeleteInvoicePaymentTagsParams{
		HTTPClient: client,
	}
}

/*DeleteInvoicePaymentTagsParams contains all the parameters to send to the API endpoint
for the delete invoice payment tags operation typically these are written to a http.Request
*/
type DeleteInvoicePaymentTagsParams struct {

	/*XKillbillComment*/
	XKillbillComment *string
	/*XKillbillCreatedBy*/
	XKillbillCreatedBy string
	/*XKillbillReason*/
	XKillbillReason *string
	/*PaymentID*/
	PaymentID strfmt.UUID
	/*TagDef*/
	TagDef []strfmt.UUID

	WithStackTrace        *bool // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the delete invoice payment tags params
func (o *DeleteInvoicePaymentTagsParams) WithTimeout(timeout time.Duration) *DeleteInvoicePaymentTagsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete invoice payment tags params
func (o *DeleteInvoicePaymentTagsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete invoice payment tags params
func (o *DeleteInvoicePaymentTagsParams) WithContext(ctx context.Context) *DeleteInvoicePaymentTagsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete invoice payment tags params
func (o *DeleteInvoicePaymentTagsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete invoice payment tags params
func (o *DeleteInvoicePaymentTagsParams) WithHTTPClient(client *http.Client) *DeleteInvoicePaymentTagsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete invoice payment tags params
func (o *DeleteInvoicePaymentTagsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the delete invoice payment tags params
func (o *DeleteInvoicePaymentTagsParams) WithXKillbillComment(xKillbillComment *string) *DeleteInvoicePaymentTagsParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the delete invoice payment tags params
func (o *DeleteInvoicePaymentTagsParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the delete invoice payment tags params
func (o *DeleteInvoicePaymentTagsParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *DeleteInvoicePaymentTagsParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the delete invoice payment tags params
func (o *DeleteInvoicePaymentTagsParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the delete invoice payment tags params
func (o *DeleteInvoicePaymentTagsParams) WithXKillbillReason(xKillbillReason *string) *DeleteInvoicePaymentTagsParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the delete invoice payment tags params
func (o *DeleteInvoicePaymentTagsParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithPaymentID adds the paymentID to the delete invoice payment tags params
func (o *DeleteInvoicePaymentTagsParams) WithPaymentID(paymentID strfmt.UUID) *DeleteInvoicePaymentTagsParams {
	o.SetPaymentID(paymentID)
	return o
}

// SetPaymentID adds the paymentId to the delete invoice payment tags params
func (o *DeleteInvoicePaymentTagsParams) SetPaymentID(paymentID strfmt.UUID) {
	o.PaymentID = paymentID
}

// WithTagDef adds the tagDef to the delete invoice payment tags params
func (o *DeleteInvoicePaymentTagsParams) WithTagDef(tagDef []strfmt.UUID) *DeleteInvoicePaymentTagsParams {
	o.SetTagDef(tagDef)
	return o
}

// SetTagDef adds the tagDef to the delete invoice payment tags params
func (o *DeleteInvoicePaymentTagsParams) SetTagDef(tagDef []strfmt.UUID) {
	o.TagDef = tagDef
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteInvoicePaymentTagsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param paymentId
	if err := r.SetPathParam("paymentId", o.PaymentID.String()); err != nil {
		return err
	}

	var valuesTagDef []string
	for _, v := range o.TagDef {
		valuesTagDef = append(valuesTagDef, v.String())
	}

	joinedTagDef := swag.JoinByFormat(valuesTagDef, "multi")
	// query array param tagDef
	if err := r.SetQueryParam("tagDef", joinedTagDef...); err != nil {
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
