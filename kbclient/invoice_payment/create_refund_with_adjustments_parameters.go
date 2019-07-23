// Code generated by go-swagger; DO NOT EDIT.

package invoice_payment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	kbmodel "github.com/killbill/kbcli/kbmodel"
)

// NewCreateRefundWithAdjustmentsParams creates a new CreateRefundWithAdjustmentsParams object
// with the default values initialized.
func NewCreateRefundWithAdjustmentsParams() *CreateRefundWithAdjustmentsParams {
	var (
		externalPaymentDefault = bool(false)
	)
	return &CreateRefundWithAdjustmentsParams{
		ExternalPayment: &externalPaymentDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateRefundWithAdjustmentsParamsWithTimeout creates a new CreateRefundWithAdjustmentsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateRefundWithAdjustmentsParamsWithTimeout(timeout time.Duration) *CreateRefundWithAdjustmentsParams {
	var (
		externalPaymentDefault = bool(false)
	)
	return &CreateRefundWithAdjustmentsParams{
		ExternalPayment: &externalPaymentDefault,

		timeout: timeout,
	}
}

// NewCreateRefundWithAdjustmentsParamsWithContext creates a new CreateRefundWithAdjustmentsParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateRefundWithAdjustmentsParamsWithContext(ctx context.Context) *CreateRefundWithAdjustmentsParams {
	var (
		externalPaymentDefault = bool(false)
	)
	return &CreateRefundWithAdjustmentsParams{
		ExternalPayment: &externalPaymentDefault,

		Context: ctx,
	}
}

// NewCreateRefundWithAdjustmentsParamsWithHTTPClient creates a new CreateRefundWithAdjustmentsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateRefundWithAdjustmentsParamsWithHTTPClient(client *http.Client) *CreateRefundWithAdjustmentsParams {
	var (
		externalPaymentDefault = bool(false)
	)
	return &CreateRefundWithAdjustmentsParams{
		ExternalPayment: &externalPaymentDefault,
		HTTPClient:      client,
	}
}

/*CreateRefundWithAdjustmentsParams contains all the parameters to send to the API endpoint
for the create refund with adjustments operation typically these are written to a http.Request
*/
type CreateRefundWithAdjustmentsParams struct {

	/*XKillbillComment*/
	XKillbillComment *string
	/*XKillbillCreatedBy*/
	XKillbillCreatedBy string
	/*XKillbillReason*/
	XKillbillReason *string
	/*Body*/
	Body *kbmodel.InvoicePaymentTransaction
	/*ExternalPayment*/
	ExternalPayment *bool
	/*PaymentID*/
	PaymentID strfmt.UUID
	/*PaymentMethodID*/
	PaymentMethodID *strfmt.UUID
	/*PluginProperty*/
	PluginProperty []string

	WithStackTrace        *bool // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the create refund with adjustments params
func (o *CreateRefundWithAdjustmentsParams) WithTimeout(timeout time.Duration) *CreateRefundWithAdjustmentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create refund with adjustments params
func (o *CreateRefundWithAdjustmentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create refund with adjustments params
func (o *CreateRefundWithAdjustmentsParams) WithContext(ctx context.Context) *CreateRefundWithAdjustmentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create refund with adjustments params
func (o *CreateRefundWithAdjustmentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create refund with adjustments params
func (o *CreateRefundWithAdjustmentsParams) WithHTTPClient(client *http.Client) *CreateRefundWithAdjustmentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create refund with adjustments params
func (o *CreateRefundWithAdjustmentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the create refund with adjustments params
func (o *CreateRefundWithAdjustmentsParams) WithXKillbillComment(xKillbillComment *string) *CreateRefundWithAdjustmentsParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the create refund with adjustments params
func (o *CreateRefundWithAdjustmentsParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the create refund with adjustments params
func (o *CreateRefundWithAdjustmentsParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *CreateRefundWithAdjustmentsParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the create refund with adjustments params
func (o *CreateRefundWithAdjustmentsParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the create refund with adjustments params
func (o *CreateRefundWithAdjustmentsParams) WithXKillbillReason(xKillbillReason *string) *CreateRefundWithAdjustmentsParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the create refund with adjustments params
func (o *CreateRefundWithAdjustmentsParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithBody adds the body to the create refund with adjustments params
func (o *CreateRefundWithAdjustmentsParams) WithBody(body *kbmodel.InvoicePaymentTransaction) *CreateRefundWithAdjustmentsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create refund with adjustments params
func (o *CreateRefundWithAdjustmentsParams) SetBody(body *kbmodel.InvoicePaymentTransaction) {
	o.Body = body
}

// WithExternalPayment adds the externalPayment to the create refund with adjustments params
func (o *CreateRefundWithAdjustmentsParams) WithExternalPayment(externalPayment *bool) *CreateRefundWithAdjustmentsParams {
	o.SetExternalPayment(externalPayment)
	return o
}

// SetExternalPayment adds the externalPayment to the create refund with adjustments params
func (o *CreateRefundWithAdjustmentsParams) SetExternalPayment(externalPayment *bool) {
	o.ExternalPayment = externalPayment
}

// WithPaymentID adds the paymentID to the create refund with adjustments params
func (o *CreateRefundWithAdjustmentsParams) WithPaymentID(paymentID strfmt.UUID) *CreateRefundWithAdjustmentsParams {
	o.SetPaymentID(paymentID)
	return o
}

// SetPaymentID adds the paymentId to the create refund with adjustments params
func (o *CreateRefundWithAdjustmentsParams) SetPaymentID(paymentID strfmt.UUID) {
	o.PaymentID = paymentID
}

// WithPaymentMethodID adds the paymentMethodID to the create refund with adjustments params
func (o *CreateRefundWithAdjustmentsParams) WithPaymentMethodID(paymentMethodID *strfmt.UUID) *CreateRefundWithAdjustmentsParams {
	o.SetPaymentMethodID(paymentMethodID)
	return o
}

// SetPaymentMethodID adds the paymentMethodId to the create refund with adjustments params
func (o *CreateRefundWithAdjustmentsParams) SetPaymentMethodID(paymentMethodID *strfmt.UUID) {
	o.PaymentMethodID = paymentMethodID
}

// WithPluginProperty adds the pluginProperty to the create refund with adjustments params
func (o *CreateRefundWithAdjustmentsParams) WithPluginProperty(pluginProperty []string) *CreateRefundWithAdjustmentsParams {
	o.SetPluginProperty(pluginProperty)
	return o
}

// SetPluginProperty adds the pluginProperty to the create refund with adjustments params
func (o *CreateRefundWithAdjustmentsParams) SetPluginProperty(pluginProperty []string) {
	o.PluginProperty = pluginProperty
}

// WriteToRequest writes these params to a swagger request
func (o *CreateRefundWithAdjustmentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.ExternalPayment != nil {

		// query param externalPayment
		var qrExternalPayment bool
		if o.ExternalPayment != nil {
			qrExternalPayment = *o.ExternalPayment
		}
		qExternalPayment := swag.FormatBool(qrExternalPayment)
		if qExternalPayment != "" {
			if err := r.SetQueryParam("externalPayment", qExternalPayment); err != nil {
				return err
			}
		}

	}

	// path param paymentId
	if err := r.SetPathParam("paymentId", o.PaymentID.String()); err != nil {
		return err
	}

	if o.PaymentMethodID != nil {

		// query param paymentMethodId
		var qrPaymentMethodID strfmt.UUID
		if o.PaymentMethodID != nil {
			qrPaymentMethodID = *o.PaymentMethodID
		}
		qPaymentMethodID := qrPaymentMethodID.String()
		if qPaymentMethodID != "" {
			if err := r.SetQueryParam("paymentMethodId", qPaymentMethodID); err != nil {
				return err
			}
		}

	}

	valuesPluginProperty := o.PluginProperty

	joinedPluginProperty := swag.JoinByFormat(valuesPluginProperty, "multi")
	// query array param pluginProperty
	if err := r.SetQueryParam("pluginProperty", joinedPluginProperty...); err != nil {
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