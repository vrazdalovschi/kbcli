// Code generated by go-swagger; DO NOT EDIT.

package subscription

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
)

// NewUncancelSubscriptionPlanParams creates a new UncancelSubscriptionPlanParams object
// with the default values initialized.
func NewUncancelSubscriptionPlanParams() *UncancelSubscriptionPlanParams {
	var ()
	return &UncancelSubscriptionPlanParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUncancelSubscriptionPlanParamsWithTimeout creates a new UncancelSubscriptionPlanParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUncancelSubscriptionPlanParamsWithTimeout(timeout time.Duration) *UncancelSubscriptionPlanParams {
	var ()
	return &UncancelSubscriptionPlanParams{

		timeout: timeout,
	}
}

// NewUncancelSubscriptionPlanParamsWithContext creates a new UncancelSubscriptionPlanParams object
// with the default values initialized, and the ability to set a context for a request
func NewUncancelSubscriptionPlanParamsWithContext(ctx context.Context) *UncancelSubscriptionPlanParams {
	var ()
	return &UncancelSubscriptionPlanParams{

		Context: ctx,
	}
}

// NewUncancelSubscriptionPlanParamsWithHTTPClient creates a new UncancelSubscriptionPlanParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUncancelSubscriptionPlanParamsWithHTTPClient(client *http.Client) *UncancelSubscriptionPlanParams {
	var ()
	return &UncancelSubscriptionPlanParams{
		HTTPClient: client,
	}
}

/*UncancelSubscriptionPlanParams contains all the parameters to send to the API endpoint
for the uncancel subscription plan operation typically these are written to a http.Request
*/
type UncancelSubscriptionPlanParams struct {

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
	/*PluginProperty*/
	PluginProperty []string
	/*SubscriptionID*/
	SubscriptionID strfmt.UUID

	WithStackTrace        *bool // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the uncancel subscription plan params
func (o *UncancelSubscriptionPlanParams) WithTimeout(timeout time.Duration) *UncancelSubscriptionPlanParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the uncancel subscription plan params
func (o *UncancelSubscriptionPlanParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the uncancel subscription plan params
func (o *UncancelSubscriptionPlanParams) WithContext(ctx context.Context) *UncancelSubscriptionPlanParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the uncancel subscription plan params
func (o *UncancelSubscriptionPlanParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the uncancel subscription plan params
func (o *UncancelSubscriptionPlanParams) WithHTTPClient(client *http.Client) *UncancelSubscriptionPlanParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the uncancel subscription plan params
func (o *UncancelSubscriptionPlanParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillAPIKey adds the xKillbillAPIKey to the uncancel subscription plan params
func (o *UncancelSubscriptionPlanParams) WithXKillbillAPIKey(xKillbillAPIKey string) *UncancelSubscriptionPlanParams {
	o.SetXKillbillAPIKey(xKillbillAPIKey)
	return o
}

// SetXKillbillAPIKey adds the xKillbillApiKey to the uncancel subscription plan params
func (o *UncancelSubscriptionPlanParams) SetXKillbillAPIKey(xKillbillAPIKey string) {
	o.XKillbillAPIKey = xKillbillAPIKey
}

// WithXKillbillAPISecret adds the xKillbillAPISecret to the uncancel subscription plan params
func (o *UncancelSubscriptionPlanParams) WithXKillbillAPISecret(xKillbillAPISecret string) *UncancelSubscriptionPlanParams {
	o.SetXKillbillAPISecret(xKillbillAPISecret)
	return o
}

// SetXKillbillAPISecret adds the xKillbillApiSecret to the uncancel subscription plan params
func (o *UncancelSubscriptionPlanParams) SetXKillbillAPISecret(xKillbillAPISecret string) {
	o.XKillbillAPISecret = xKillbillAPISecret
}

// WithXKillbillComment adds the xKillbillComment to the uncancel subscription plan params
func (o *UncancelSubscriptionPlanParams) WithXKillbillComment(xKillbillComment *string) *UncancelSubscriptionPlanParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the uncancel subscription plan params
func (o *UncancelSubscriptionPlanParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the uncancel subscription plan params
func (o *UncancelSubscriptionPlanParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *UncancelSubscriptionPlanParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the uncancel subscription plan params
func (o *UncancelSubscriptionPlanParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the uncancel subscription plan params
func (o *UncancelSubscriptionPlanParams) WithXKillbillReason(xKillbillReason *string) *UncancelSubscriptionPlanParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the uncancel subscription plan params
func (o *UncancelSubscriptionPlanParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithPluginProperty adds the pluginProperty to the uncancel subscription plan params
func (o *UncancelSubscriptionPlanParams) WithPluginProperty(pluginProperty []string) *UncancelSubscriptionPlanParams {
	o.SetPluginProperty(pluginProperty)
	return o
}

// SetPluginProperty adds the pluginProperty to the uncancel subscription plan params
func (o *UncancelSubscriptionPlanParams) SetPluginProperty(pluginProperty []string) {
	o.PluginProperty = pluginProperty
}

// WithSubscriptionID adds the subscriptionID to the uncancel subscription plan params
func (o *UncancelSubscriptionPlanParams) WithSubscriptionID(subscriptionID strfmt.UUID) *UncancelSubscriptionPlanParams {
	o.SetSubscriptionID(subscriptionID)
	return o
}

// SetSubscriptionID adds the subscriptionId to the uncancel subscription plan params
func (o *UncancelSubscriptionPlanParams) SetSubscriptionID(subscriptionID strfmt.UUID) {
	o.SubscriptionID = subscriptionID
}

// WriteToRequest writes these params to a swagger request
func (o *UncancelSubscriptionPlanParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	valuesPluginProperty := o.PluginProperty

	joinedPluginProperty := swag.JoinByFormat(valuesPluginProperty, "multi")
	// query array param pluginProperty
	if err := r.SetQueryParam("pluginProperty", joinedPluginProperty...); err != nil {
		return err
	}

	// path param subscriptionId
	if err := r.SetPathParam("subscriptionId", o.SubscriptionID.String()); err != nil {
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
