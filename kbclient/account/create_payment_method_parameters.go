// Code generated by go-swagger; DO NOT EDIT.

package account

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

// NewCreatePaymentMethodParams creates a new CreatePaymentMethodParams object
// with the default values initialized.
func NewCreatePaymentMethodParams() *CreatePaymentMethodParams {
	var (
		isDefaultDefault            = bool(false)
		payAllUnpaidInvoicesDefault = bool(false)
	)
	return &CreatePaymentMethodParams{
		IsDefault:            &isDefaultDefault,
		PayAllUnpaidInvoices: &payAllUnpaidInvoicesDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewCreatePaymentMethodParamsWithTimeout creates a new CreatePaymentMethodParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreatePaymentMethodParamsWithTimeout(timeout time.Duration) *CreatePaymentMethodParams {
	var (
		isDefaultDefault            = bool(false)
		payAllUnpaidInvoicesDefault = bool(false)
	)
	return &CreatePaymentMethodParams{
		IsDefault:            &isDefaultDefault,
		PayAllUnpaidInvoices: &payAllUnpaidInvoicesDefault,

		timeout: timeout,
	}
}

// NewCreatePaymentMethodParamsWithContext creates a new CreatePaymentMethodParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreatePaymentMethodParamsWithContext(ctx context.Context) *CreatePaymentMethodParams {
	var (
		isDefaultDefault            = bool(false)
		payAllUnpaidInvoicesDefault = bool(false)
	)
	return &CreatePaymentMethodParams{
		IsDefault:            &isDefaultDefault,
		PayAllUnpaidInvoices: &payAllUnpaidInvoicesDefault,

		Context: ctx,
	}
}

// NewCreatePaymentMethodParamsWithHTTPClient creates a new CreatePaymentMethodParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreatePaymentMethodParamsWithHTTPClient(client *http.Client) *CreatePaymentMethodParams {
	var (
		isDefaultDefault            = bool(false)
		payAllUnpaidInvoicesDefault = bool(false)
	)
	return &CreatePaymentMethodParams{
		IsDefault:            &isDefaultDefault,
		PayAllUnpaidInvoices: &payAllUnpaidInvoicesDefault,
		HTTPClient:           client,
	}
}

/*CreatePaymentMethodParams contains all the parameters to send to the API endpoint
for the create payment method operation typically these are written to a http.Request
*/
type CreatePaymentMethodParams struct {

	/*XKillbillAPIKey*/
	XKillbillAPIKey string
	/*XKillbillAPISecret*/
	XKillbillAPISecret string
	/*XKillbillComment*/
	XKillbillComment *string
	/*XKillbillCreatedBy*/
	XKillbillCreatedBy string
	/*XKillbillReason*/
	XKillbillReason *string
	/*AccountID*/
	AccountID strfmt.UUID
	/*Body*/
	Body *kbmodel.PaymentMethod
	/*ControlPluginName*/
	ControlPluginName []string
	/*IsDefault*/
	IsDefault *bool
	/*PayAllUnpaidInvoices*/
	PayAllUnpaidInvoices *bool
	/*PluginProperty*/
	PluginProperty []string

	WithStackTrace        *bool // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the create payment method params
func (o *CreatePaymentMethodParams) WithTimeout(timeout time.Duration) *CreatePaymentMethodParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create payment method params
func (o *CreatePaymentMethodParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create payment method params
func (o *CreatePaymentMethodParams) WithContext(ctx context.Context) *CreatePaymentMethodParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create payment method params
func (o *CreatePaymentMethodParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create payment method params
func (o *CreatePaymentMethodParams) WithHTTPClient(client *http.Client) *CreatePaymentMethodParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create payment method params
func (o *CreatePaymentMethodParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillAPIKey adds the xKillbillAPIKey to the create payment method params
func (o *CreatePaymentMethodParams) WithXKillbillAPIKey(xKillbillAPIKey string) *CreatePaymentMethodParams {
	o.SetXKillbillAPIKey(xKillbillAPIKey)
	return o
}

// SetXKillbillAPIKey adds the xKillbillApiKey to the create payment method params
func (o *CreatePaymentMethodParams) SetXKillbillAPIKey(xKillbillAPIKey string) {
	o.XKillbillAPIKey = xKillbillAPIKey
}

// WithXKillbillAPISecret adds the xKillbillAPISecret to the create payment method params
func (o *CreatePaymentMethodParams) WithXKillbillAPISecret(xKillbillAPISecret string) *CreatePaymentMethodParams {
	o.SetXKillbillAPISecret(xKillbillAPISecret)
	return o
}

// SetXKillbillAPISecret adds the xKillbillApiSecret to the create payment method params
func (o *CreatePaymentMethodParams) SetXKillbillAPISecret(xKillbillAPISecret string) {
	o.XKillbillAPISecret = xKillbillAPISecret
}

// WithXKillbillComment adds the xKillbillComment to the create payment method params
func (o *CreatePaymentMethodParams) WithXKillbillComment(xKillbillComment *string) *CreatePaymentMethodParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the create payment method params
func (o *CreatePaymentMethodParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the create payment method params
func (o *CreatePaymentMethodParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *CreatePaymentMethodParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the create payment method params
func (o *CreatePaymentMethodParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the create payment method params
func (o *CreatePaymentMethodParams) WithXKillbillReason(xKillbillReason *string) *CreatePaymentMethodParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the create payment method params
func (o *CreatePaymentMethodParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithAccountID adds the accountID to the create payment method params
func (o *CreatePaymentMethodParams) WithAccountID(accountID strfmt.UUID) *CreatePaymentMethodParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the create payment method params
func (o *CreatePaymentMethodParams) SetAccountID(accountID strfmt.UUID) {
	o.AccountID = accountID
}

// WithBody adds the body to the create payment method params
func (o *CreatePaymentMethodParams) WithBody(body *kbmodel.PaymentMethod) *CreatePaymentMethodParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create payment method params
func (o *CreatePaymentMethodParams) SetBody(body *kbmodel.PaymentMethod) {
	o.Body = body
}

// WithControlPluginName adds the controlPluginName to the create payment method params
func (o *CreatePaymentMethodParams) WithControlPluginName(controlPluginName []string) *CreatePaymentMethodParams {
	o.SetControlPluginName(controlPluginName)
	return o
}

// SetControlPluginName adds the controlPluginName to the create payment method params
func (o *CreatePaymentMethodParams) SetControlPluginName(controlPluginName []string) {
	o.ControlPluginName = controlPluginName
}

// WithIsDefault adds the isDefault to the create payment method params
func (o *CreatePaymentMethodParams) WithIsDefault(isDefault *bool) *CreatePaymentMethodParams {
	o.SetIsDefault(isDefault)
	return o
}

// SetIsDefault adds the isDefault to the create payment method params
func (o *CreatePaymentMethodParams) SetIsDefault(isDefault *bool) {
	o.IsDefault = isDefault
}

// WithPayAllUnpaidInvoices adds the payAllUnpaidInvoices to the create payment method params
func (o *CreatePaymentMethodParams) WithPayAllUnpaidInvoices(payAllUnpaidInvoices *bool) *CreatePaymentMethodParams {
	o.SetPayAllUnpaidInvoices(payAllUnpaidInvoices)
	return o
}

// SetPayAllUnpaidInvoices adds the payAllUnpaidInvoices to the create payment method params
func (o *CreatePaymentMethodParams) SetPayAllUnpaidInvoices(payAllUnpaidInvoices *bool) {
	o.PayAllUnpaidInvoices = payAllUnpaidInvoices
}

// WithPluginProperty adds the pluginProperty to the create payment method params
func (o *CreatePaymentMethodParams) WithPluginProperty(pluginProperty []string) *CreatePaymentMethodParams {
	o.SetPluginProperty(pluginProperty)
	return o
}

// SetPluginProperty adds the pluginProperty to the create payment method params
func (o *CreatePaymentMethodParams) SetPluginProperty(pluginProperty []string) {
	o.PluginProperty = pluginProperty
}

// WriteToRequest writes these params to a swagger request
func (o *CreatePaymentMethodParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param X-Killbill-ApiKey
	if err := r.SetHeaderParam("X-Killbill-ApiKey", o.XKillbillAPIKey); err != nil {
		return err
	}

	// header param X-Killbill-ApiSecret
	if err := r.SetHeaderParam("X-Killbill-ApiSecret", o.XKillbillAPISecret); err != nil {
		return err
	}

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

	// path param accountId
	if err := r.SetPathParam("accountId", o.AccountID.String()); err != nil {
		return err
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

	if o.IsDefault != nil {

		// query param isDefault
		var qrIsDefault bool
		if o.IsDefault != nil {
			qrIsDefault = *o.IsDefault
		}
		qIsDefault := swag.FormatBool(qrIsDefault)
		if qIsDefault != "" {
			if err := r.SetQueryParam("isDefault", qIsDefault); err != nil {
				return err
			}
		}

	}

	if o.PayAllUnpaidInvoices != nil {

		// query param payAllUnpaidInvoices
		var qrPayAllUnpaidInvoices bool
		if o.PayAllUnpaidInvoices != nil {
			qrPayAllUnpaidInvoices = *o.PayAllUnpaidInvoices
		}
		qPayAllUnpaidInvoices := swag.FormatBool(qrPayAllUnpaidInvoices)
		if qPayAllUnpaidInvoices != "" {
			if err := r.SetQueryParam("payAllUnpaidInvoices", qPayAllUnpaidInvoices); err != nil {
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
