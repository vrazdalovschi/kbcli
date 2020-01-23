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

	kbmodel "github.com/killbill/kbcli/kbmodel"
)

// NewRefundPaymentByExternalKeyParams creates a new RefundPaymentByExternalKeyParams object
// with the default values initialized.
func NewRefundPaymentByExternalKeyParams() *RefundPaymentByExternalKeyParams {
	var ()
	return &RefundPaymentByExternalKeyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRefundPaymentByExternalKeyParamsWithTimeout creates a new RefundPaymentByExternalKeyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRefundPaymentByExternalKeyParamsWithTimeout(timeout time.Duration) *RefundPaymentByExternalKeyParams {
	var ()
	return &RefundPaymentByExternalKeyParams{

		timeout: timeout,
	}
}

// NewRefundPaymentByExternalKeyParamsWithContext creates a new RefundPaymentByExternalKeyParams object
// with the default values initialized, and the ability to set a context for a request
func NewRefundPaymentByExternalKeyParamsWithContext(ctx context.Context) *RefundPaymentByExternalKeyParams {
	var ()
	return &RefundPaymentByExternalKeyParams{

		Context: ctx,
	}
}

// NewRefundPaymentByExternalKeyParamsWithHTTPClient creates a new RefundPaymentByExternalKeyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRefundPaymentByExternalKeyParamsWithHTTPClient(client *http.Client) *RefundPaymentByExternalKeyParams {
	var ()
	return &RefundPaymentByExternalKeyParams{
		HTTPClient: client,
	}
}

/*RefundPaymentByExternalKeyParams contains all the parameters to send to the API endpoint
for the refund payment by external key operation typically these are written to a http.Request
*/
type RefundPaymentByExternalKeyParams struct {

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
	/*PluginProperty*/
	PluginProperty []string

	WithStackTrace        *bool // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the refund payment by external key params
func (o *RefundPaymentByExternalKeyParams) WithTimeout(timeout time.Duration) *RefundPaymentByExternalKeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the refund payment by external key params
func (o *RefundPaymentByExternalKeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the refund payment by external key params
func (o *RefundPaymentByExternalKeyParams) WithContext(ctx context.Context) *RefundPaymentByExternalKeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the refund payment by external key params
func (o *RefundPaymentByExternalKeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the refund payment by external key params
func (o *RefundPaymentByExternalKeyParams) WithHTTPClient(client *http.Client) *RefundPaymentByExternalKeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the refund payment by external key params
func (o *RefundPaymentByExternalKeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the refund payment by external key params
func (o *RefundPaymentByExternalKeyParams) WithXKillbillComment(xKillbillComment *string) *RefundPaymentByExternalKeyParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the refund payment by external key params
func (o *RefundPaymentByExternalKeyParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the refund payment by external key params
func (o *RefundPaymentByExternalKeyParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *RefundPaymentByExternalKeyParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the refund payment by external key params
func (o *RefundPaymentByExternalKeyParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the refund payment by external key params
func (o *RefundPaymentByExternalKeyParams) WithXKillbillReason(xKillbillReason *string) *RefundPaymentByExternalKeyParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the refund payment by external key params
func (o *RefundPaymentByExternalKeyParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithBody adds the body to the refund payment by external key params
func (o *RefundPaymentByExternalKeyParams) WithBody(body *kbmodel.PaymentTransaction) *RefundPaymentByExternalKeyParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the refund payment by external key params
func (o *RefundPaymentByExternalKeyParams) SetBody(body *kbmodel.PaymentTransaction) {
	o.Body = body
}

// WithControlPluginName adds the controlPluginName to the refund payment by external key params
func (o *RefundPaymentByExternalKeyParams) WithControlPluginName(controlPluginName []string) *RefundPaymentByExternalKeyParams {
	o.SetControlPluginName(controlPluginName)
	return o
}

// SetControlPluginName adds the controlPluginName to the refund payment by external key params
func (o *RefundPaymentByExternalKeyParams) SetControlPluginName(controlPluginName []string) {
	o.ControlPluginName = controlPluginName
}

// WithPluginProperty adds the pluginProperty to the refund payment by external key params
func (o *RefundPaymentByExternalKeyParams) WithPluginProperty(pluginProperty []string) *RefundPaymentByExternalKeyParams {
	o.SetPluginProperty(pluginProperty)
	return o
}

// SetPluginProperty adds the pluginProperty to the refund payment by external key params
func (o *RefundPaymentByExternalKeyParams) SetPluginProperty(pluginProperty []string) {
	o.PluginProperty = pluginProperty
}

// WriteToRequest writes these params to a swagger request
func (o *RefundPaymentByExternalKeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
