// Code generated by go-swagger; DO NOT EDIT.

package payment_method

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

	kbmodel "github.com/killbill/kbcli/v2/kbmodel"
)

// NewModifyPaymentMethodCustomFieldsParams creates a new ModifyPaymentMethodCustomFieldsParams object
// with the default values initialized.
func NewModifyPaymentMethodCustomFieldsParams() *ModifyPaymentMethodCustomFieldsParams {
	var ()
	return &ModifyPaymentMethodCustomFieldsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewModifyPaymentMethodCustomFieldsParamsWithTimeout creates a new ModifyPaymentMethodCustomFieldsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewModifyPaymentMethodCustomFieldsParamsWithTimeout(timeout time.Duration) *ModifyPaymentMethodCustomFieldsParams {
	var ()
	return &ModifyPaymentMethodCustomFieldsParams{

		timeout: timeout,
	}
}

// NewModifyPaymentMethodCustomFieldsParamsWithContext creates a new ModifyPaymentMethodCustomFieldsParams object
// with the default values initialized, and the ability to set a context for a request
func NewModifyPaymentMethodCustomFieldsParamsWithContext(ctx context.Context) *ModifyPaymentMethodCustomFieldsParams {
	var ()
	return &ModifyPaymentMethodCustomFieldsParams{

		Context: ctx,
	}
}

// NewModifyPaymentMethodCustomFieldsParamsWithHTTPClient creates a new ModifyPaymentMethodCustomFieldsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewModifyPaymentMethodCustomFieldsParamsWithHTTPClient(client *http.Client) *ModifyPaymentMethodCustomFieldsParams {
	var ()
	return &ModifyPaymentMethodCustomFieldsParams{
		HTTPClient: client,
	}
}

/*ModifyPaymentMethodCustomFieldsParams contains all the parameters to send to the API endpoint
for the modify payment method custom fields operation typically these are written to a http.Request
*/
type ModifyPaymentMethodCustomFieldsParams struct {

	/*XKillbillComment*/
	XKillbillComment *string
	/*XKillbillCreatedBy*/
	XKillbillCreatedBy string
	/*XKillbillReason*/
	XKillbillReason *string
	/*Body*/
	Body []*kbmodel.CustomField
	/*PaymentMethodID*/
	PaymentMethodID strfmt.UUID

	WithStackTrace        *bool // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the modify payment method custom fields params
func (o *ModifyPaymentMethodCustomFieldsParams) WithTimeout(timeout time.Duration) *ModifyPaymentMethodCustomFieldsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the modify payment method custom fields params
func (o *ModifyPaymentMethodCustomFieldsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the modify payment method custom fields params
func (o *ModifyPaymentMethodCustomFieldsParams) WithContext(ctx context.Context) *ModifyPaymentMethodCustomFieldsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the modify payment method custom fields params
func (o *ModifyPaymentMethodCustomFieldsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the modify payment method custom fields params
func (o *ModifyPaymentMethodCustomFieldsParams) WithHTTPClient(client *http.Client) *ModifyPaymentMethodCustomFieldsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the modify payment method custom fields params
func (o *ModifyPaymentMethodCustomFieldsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the modify payment method custom fields params
func (o *ModifyPaymentMethodCustomFieldsParams) WithXKillbillComment(xKillbillComment *string) *ModifyPaymentMethodCustomFieldsParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the modify payment method custom fields params
func (o *ModifyPaymentMethodCustomFieldsParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the modify payment method custom fields params
func (o *ModifyPaymentMethodCustomFieldsParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *ModifyPaymentMethodCustomFieldsParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the modify payment method custom fields params
func (o *ModifyPaymentMethodCustomFieldsParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the modify payment method custom fields params
func (o *ModifyPaymentMethodCustomFieldsParams) WithXKillbillReason(xKillbillReason *string) *ModifyPaymentMethodCustomFieldsParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the modify payment method custom fields params
func (o *ModifyPaymentMethodCustomFieldsParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithBody adds the body to the modify payment method custom fields params
func (o *ModifyPaymentMethodCustomFieldsParams) WithBody(body []*kbmodel.CustomField) *ModifyPaymentMethodCustomFieldsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the modify payment method custom fields params
func (o *ModifyPaymentMethodCustomFieldsParams) SetBody(body []*kbmodel.CustomField) {
	o.Body = body
}

// WithPaymentMethodID adds the paymentMethodID to the modify payment method custom fields params
func (o *ModifyPaymentMethodCustomFieldsParams) WithPaymentMethodID(paymentMethodID strfmt.UUID) *ModifyPaymentMethodCustomFieldsParams {
	o.SetPaymentMethodID(paymentMethodID)
	return o
}

// SetPaymentMethodID adds the paymentMethodId to the modify payment method custom fields params
func (o *ModifyPaymentMethodCustomFieldsParams) SetPaymentMethodID(paymentMethodID strfmt.UUID) {
	o.PaymentMethodID = paymentMethodID
}

// WriteToRequest writes these params to a swagger request
func (o *ModifyPaymentMethodCustomFieldsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param paymentMethodId
	if err := r.SetPathParam("paymentMethodId", o.PaymentMethodID.String()); err != nil {
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
