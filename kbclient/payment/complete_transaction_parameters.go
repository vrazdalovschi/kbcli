// Code generated by go-swagger; DO NOT EDIT.

package payment

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

	kbmodel "github.com/killbill/kbcli/v2/kbmodel"
)

// NewCompleteTransactionParams creates a new CompleteTransactionParams object
// with the default values initialized.
func NewCompleteTransactionParams() *CompleteTransactionParams {
	var ()
	return &CompleteTransactionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCompleteTransactionParamsWithTimeout creates a new CompleteTransactionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCompleteTransactionParamsWithTimeout(timeout time.Duration) *CompleteTransactionParams {
	var ()
	return &CompleteTransactionParams{

		timeout: timeout,
	}
}

// NewCompleteTransactionParamsWithContext creates a new CompleteTransactionParams object
// with the default values initialized, and the ability to set a context for a request
func NewCompleteTransactionParamsWithContext(ctx context.Context) *CompleteTransactionParams {
	var ()
	return &CompleteTransactionParams{

		Context: ctx,
	}
}

// NewCompleteTransactionParamsWithHTTPClient creates a new CompleteTransactionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCompleteTransactionParamsWithHTTPClient(client *http.Client) *CompleteTransactionParams {
	var ()
	return &CompleteTransactionParams{
		HTTPClient: client,
	}
}

/*CompleteTransactionParams contains all the parameters to send to the API endpoint
for the complete transaction operation typically these are written to a http.Request
*/
type CompleteTransactionParams struct {

	/*XKillbillComment*/
	XKillbillComment *string
	/*XKillbillCreatedBy*/
	XKillbillCreatedBy string
	/*XKillbillReason*/
	XKillbillReason *string
	/*Body*/
	Body *kbmodel.PaymentTransaction
	/*ControlPluginName*/
	ControlPluginName []string
	/*PaymentID*/
	PaymentID strfmt.UUID
	/*PluginProperty*/
	PluginProperty []string

	WithStackTrace        *bool // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the complete transaction params
func (o *CompleteTransactionParams) WithTimeout(timeout time.Duration) *CompleteTransactionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the complete transaction params
func (o *CompleteTransactionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the complete transaction params
func (o *CompleteTransactionParams) WithContext(ctx context.Context) *CompleteTransactionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the complete transaction params
func (o *CompleteTransactionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the complete transaction params
func (o *CompleteTransactionParams) WithHTTPClient(client *http.Client) *CompleteTransactionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the complete transaction params
func (o *CompleteTransactionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the complete transaction params
func (o *CompleteTransactionParams) WithXKillbillComment(xKillbillComment *string) *CompleteTransactionParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the complete transaction params
func (o *CompleteTransactionParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the complete transaction params
func (o *CompleteTransactionParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *CompleteTransactionParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the complete transaction params
func (o *CompleteTransactionParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the complete transaction params
func (o *CompleteTransactionParams) WithXKillbillReason(xKillbillReason *string) *CompleteTransactionParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the complete transaction params
func (o *CompleteTransactionParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithBody adds the body to the complete transaction params
func (o *CompleteTransactionParams) WithBody(body *kbmodel.PaymentTransaction) *CompleteTransactionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the complete transaction params
func (o *CompleteTransactionParams) SetBody(body *kbmodel.PaymentTransaction) {
	o.Body = body
}

// WithControlPluginName adds the controlPluginName to the complete transaction params
func (o *CompleteTransactionParams) WithControlPluginName(controlPluginName []string) *CompleteTransactionParams {
	o.SetControlPluginName(controlPluginName)
	return o
}

// SetControlPluginName adds the controlPluginName to the complete transaction params
func (o *CompleteTransactionParams) SetControlPluginName(controlPluginName []string) {
	o.ControlPluginName = controlPluginName
}

// WithPaymentID adds the paymentID to the complete transaction params
func (o *CompleteTransactionParams) WithPaymentID(paymentID strfmt.UUID) *CompleteTransactionParams {
	o.SetPaymentID(paymentID)
	return o
}

// SetPaymentID adds the paymentId to the complete transaction params
func (o *CompleteTransactionParams) SetPaymentID(paymentID strfmt.UUID) {
	o.PaymentID = paymentID
}

// WithPluginProperty adds the pluginProperty to the complete transaction params
func (o *CompleteTransactionParams) WithPluginProperty(pluginProperty []string) *CompleteTransactionParams {
	o.SetPluginProperty(pluginProperty)
	return o
}

// SetPluginProperty adds the pluginProperty to the complete transaction params
func (o *CompleteTransactionParams) SetPluginProperty(pluginProperty []string) {
	o.PluginProperty = pluginProperty
}

// WriteToRequest writes these params to a swagger request
func (o *CompleteTransactionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	valuesControlPluginName := o.ControlPluginName

	joinedControlPluginName := swag.JoinByFormat(valuesControlPluginName, "multi")
	// query array param controlPluginName
	if err := r.SetQueryParam("controlPluginName", joinedControlPluginName...); err != nil {
		return err
	}

	// path param paymentId
	if err := r.SetPathParam("paymentId", o.PaymentID.String()); err != nil {
		return err
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
