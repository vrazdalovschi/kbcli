// Code generated by go-swagger; DO NOT EDIT.

package tenant

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetPluginPaymentStateMachineConfigParams creates a new GetPluginPaymentStateMachineConfigParams object
// with the default values initialized.
func NewGetPluginPaymentStateMachineConfigParams() *GetPluginPaymentStateMachineConfigParams {
	var ()
	return &GetPluginPaymentStateMachineConfigParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPluginPaymentStateMachineConfigParamsWithTimeout creates a new GetPluginPaymentStateMachineConfigParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPluginPaymentStateMachineConfigParamsWithTimeout(timeout time.Duration) *GetPluginPaymentStateMachineConfigParams {
	var ()
	return &GetPluginPaymentStateMachineConfigParams{

		timeout: timeout,
	}
}

// NewGetPluginPaymentStateMachineConfigParamsWithContext creates a new GetPluginPaymentStateMachineConfigParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPluginPaymentStateMachineConfigParamsWithContext(ctx context.Context) *GetPluginPaymentStateMachineConfigParams {
	var ()
	return &GetPluginPaymentStateMachineConfigParams{

		Context: ctx,
	}
}

// NewGetPluginPaymentStateMachineConfigParamsWithHTTPClient creates a new GetPluginPaymentStateMachineConfigParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPluginPaymentStateMachineConfigParamsWithHTTPClient(client *http.Client) *GetPluginPaymentStateMachineConfigParams {
	var ()
	return &GetPluginPaymentStateMachineConfigParams{
		HTTPClient: client,
	}
}

/*GetPluginPaymentStateMachineConfigParams contains all the parameters to send to the API endpoint
for the get plugin payment state machine config operation typically these are written to a http.Request
*/
type GetPluginPaymentStateMachineConfigParams struct {

	/*XKillbillAPIKey*/
	XKillbillAPIKey string
	/*XKillbillAPISecret*/
	XKillbillAPISecret string
	/*PluginName*/
	PluginName string

	WithStackTrace        *bool // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the get plugin payment state machine config params
func (o *GetPluginPaymentStateMachineConfigParams) WithTimeout(timeout time.Duration) *GetPluginPaymentStateMachineConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get plugin payment state machine config params
func (o *GetPluginPaymentStateMachineConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get plugin payment state machine config params
func (o *GetPluginPaymentStateMachineConfigParams) WithContext(ctx context.Context) *GetPluginPaymentStateMachineConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get plugin payment state machine config params
func (o *GetPluginPaymentStateMachineConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get plugin payment state machine config params
func (o *GetPluginPaymentStateMachineConfigParams) WithHTTPClient(client *http.Client) *GetPluginPaymentStateMachineConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get plugin payment state machine config params
func (o *GetPluginPaymentStateMachineConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillAPIKey adds the xKillbillAPIKey to the get plugin payment state machine config params
func (o *GetPluginPaymentStateMachineConfigParams) WithXKillbillAPIKey(xKillbillAPIKey string) *GetPluginPaymentStateMachineConfigParams {
	o.SetXKillbillAPIKey(xKillbillAPIKey)
	return o
}

// SetXKillbillAPIKey adds the xKillbillApiKey to the get plugin payment state machine config params
func (o *GetPluginPaymentStateMachineConfigParams) SetXKillbillAPIKey(xKillbillAPIKey string) {
	o.XKillbillAPIKey = xKillbillAPIKey
}

// WithXKillbillAPISecret adds the xKillbillAPISecret to the get plugin payment state machine config params
func (o *GetPluginPaymentStateMachineConfigParams) WithXKillbillAPISecret(xKillbillAPISecret string) *GetPluginPaymentStateMachineConfigParams {
	o.SetXKillbillAPISecret(xKillbillAPISecret)
	return o
}

// SetXKillbillAPISecret adds the xKillbillApiSecret to the get plugin payment state machine config params
func (o *GetPluginPaymentStateMachineConfigParams) SetXKillbillAPISecret(xKillbillAPISecret string) {
	o.XKillbillAPISecret = xKillbillAPISecret
}

// WithPluginName adds the pluginName to the get plugin payment state machine config params
func (o *GetPluginPaymentStateMachineConfigParams) WithPluginName(pluginName string) *GetPluginPaymentStateMachineConfigParams {
	o.SetPluginName(pluginName)
	return o
}

// SetPluginName adds the pluginName to the get plugin payment state machine config params
func (o *GetPluginPaymentStateMachineConfigParams) SetPluginName(pluginName string) {
	o.PluginName = pluginName
}

// WriteToRequest writes these params to a swagger request
func (o *GetPluginPaymentStateMachineConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param pluginName
	if err := r.SetPathParam("pluginName", o.PluginName); err != nil {
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
