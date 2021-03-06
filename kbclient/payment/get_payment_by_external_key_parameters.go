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
)

// NewGetPaymentByExternalKeyParams creates a new GetPaymentByExternalKeyParams object
// with the default values initialized.
func NewGetPaymentByExternalKeyParams() *GetPaymentByExternalKeyParams {
	var (
		auditDefault          = string("NONE")
		withAttemptsDefault   = bool(false)
		withPluginInfoDefault = bool(false)
	)
	return &GetPaymentByExternalKeyParams{
		Audit:          &auditDefault,
		WithAttempts:   &withAttemptsDefault,
		WithPluginInfo: &withPluginInfoDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPaymentByExternalKeyParamsWithTimeout creates a new GetPaymentByExternalKeyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPaymentByExternalKeyParamsWithTimeout(timeout time.Duration) *GetPaymentByExternalKeyParams {
	var (
		auditDefault          = string("NONE")
		withAttemptsDefault   = bool(false)
		withPluginInfoDefault = bool(false)
	)
	return &GetPaymentByExternalKeyParams{
		Audit:          &auditDefault,
		WithAttempts:   &withAttemptsDefault,
		WithPluginInfo: &withPluginInfoDefault,

		timeout: timeout,
	}
}

// NewGetPaymentByExternalKeyParamsWithContext creates a new GetPaymentByExternalKeyParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPaymentByExternalKeyParamsWithContext(ctx context.Context) *GetPaymentByExternalKeyParams {
	var (
		auditDefault          = string("NONE")
		withAttemptsDefault   = bool(false)
		withPluginInfoDefault = bool(false)
	)
	return &GetPaymentByExternalKeyParams{
		Audit:          &auditDefault,
		WithAttempts:   &withAttemptsDefault,
		WithPluginInfo: &withPluginInfoDefault,

		Context: ctx,
	}
}

// NewGetPaymentByExternalKeyParamsWithHTTPClient creates a new GetPaymentByExternalKeyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPaymentByExternalKeyParamsWithHTTPClient(client *http.Client) *GetPaymentByExternalKeyParams {
	var (
		auditDefault          = string("NONE")
		withAttemptsDefault   = bool(false)
		withPluginInfoDefault = bool(false)
	)
	return &GetPaymentByExternalKeyParams{
		Audit:          &auditDefault,
		WithAttempts:   &withAttemptsDefault,
		WithPluginInfo: &withPluginInfoDefault,
		HTTPClient:     client,
	}
}

/*GetPaymentByExternalKeyParams contains all the parameters to send to the API endpoint
for the get payment by external key operation typically these are written to a http.Request
*/
type GetPaymentByExternalKeyParams struct {

	/*Audit*/
	Audit *string
	/*ExternalKey*/
	ExternalKey string
	/*PluginProperty*/
	PluginProperty []string
	/*WithAttempts*/
	WithAttempts *bool
	/*WithPluginInfo*/
	WithPluginInfo *bool

	WithStackTrace        *bool // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the get payment by external key params
func (o *GetPaymentByExternalKeyParams) WithTimeout(timeout time.Duration) *GetPaymentByExternalKeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get payment by external key params
func (o *GetPaymentByExternalKeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get payment by external key params
func (o *GetPaymentByExternalKeyParams) WithContext(ctx context.Context) *GetPaymentByExternalKeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get payment by external key params
func (o *GetPaymentByExternalKeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get payment by external key params
func (o *GetPaymentByExternalKeyParams) WithHTTPClient(client *http.Client) *GetPaymentByExternalKeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get payment by external key params
func (o *GetPaymentByExternalKeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAudit adds the audit to the get payment by external key params
func (o *GetPaymentByExternalKeyParams) WithAudit(audit *string) *GetPaymentByExternalKeyParams {
	o.SetAudit(audit)
	return o
}

// SetAudit adds the audit to the get payment by external key params
func (o *GetPaymentByExternalKeyParams) SetAudit(audit *string) {
	o.Audit = audit
}

// WithExternalKey adds the externalKey to the get payment by external key params
func (o *GetPaymentByExternalKeyParams) WithExternalKey(externalKey string) *GetPaymentByExternalKeyParams {
	o.SetExternalKey(externalKey)
	return o
}

// SetExternalKey adds the externalKey to the get payment by external key params
func (o *GetPaymentByExternalKeyParams) SetExternalKey(externalKey string) {
	o.ExternalKey = externalKey
}

// WithPluginProperty adds the pluginProperty to the get payment by external key params
func (o *GetPaymentByExternalKeyParams) WithPluginProperty(pluginProperty []string) *GetPaymentByExternalKeyParams {
	o.SetPluginProperty(pluginProperty)
	return o
}

// SetPluginProperty adds the pluginProperty to the get payment by external key params
func (o *GetPaymentByExternalKeyParams) SetPluginProperty(pluginProperty []string) {
	o.PluginProperty = pluginProperty
}

// WithWithAttempts adds the withAttempts to the get payment by external key params
func (o *GetPaymentByExternalKeyParams) WithWithAttempts(withAttempts *bool) *GetPaymentByExternalKeyParams {
	o.SetWithAttempts(withAttempts)
	return o
}

// SetWithAttempts adds the withAttempts to the get payment by external key params
func (o *GetPaymentByExternalKeyParams) SetWithAttempts(withAttempts *bool) {
	o.WithAttempts = withAttempts
}

// WithWithPluginInfo adds the withPluginInfo to the get payment by external key params
func (o *GetPaymentByExternalKeyParams) WithWithPluginInfo(withPluginInfo *bool) *GetPaymentByExternalKeyParams {
	o.SetWithPluginInfo(withPluginInfo)
	return o
}

// SetWithPluginInfo adds the withPluginInfo to the get payment by external key params
func (o *GetPaymentByExternalKeyParams) SetWithPluginInfo(withPluginInfo *bool) {
	o.WithPluginInfo = withPluginInfo
}

// WriteToRequest writes these params to a swagger request
func (o *GetPaymentByExternalKeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Audit != nil {

		// query param audit
		var qrAudit string
		if o.Audit != nil {
			qrAudit = *o.Audit
		}
		qAudit := qrAudit
		if qAudit != "" {
			if err := r.SetQueryParam("audit", qAudit); err != nil {
				return err
			}
		}

	}

	// query param externalKey
	qrExternalKey := o.ExternalKey
	qExternalKey := qrExternalKey
	if qExternalKey != "" {
		if err := r.SetQueryParam("externalKey", qExternalKey); err != nil {
			return err
		}
	}

	valuesPluginProperty := o.PluginProperty

	joinedPluginProperty := swag.JoinByFormat(valuesPluginProperty, "multi")
	// query array param pluginProperty
	if err := r.SetQueryParam("pluginProperty", joinedPluginProperty...); err != nil {
		return err
	}

	if o.WithAttempts != nil {

		// query param withAttempts
		var qrWithAttempts bool
		if o.WithAttempts != nil {
			qrWithAttempts = *o.WithAttempts
		}
		qWithAttempts := swag.FormatBool(qrWithAttempts)
		if qWithAttempts != "" {
			if err := r.SetQueryParam("withAttempts", qWithAttempts); err != nil {
				return err
			}
		}

	}

	if o.WithPluginInfo != nil {

		// query param withPluginInfo
		var qrWithPluginInfo bool
		if o.WithPluginInfo != nil {
			qrWithPluginInfo = *o.WithPluginInfo
		}
		qWithPluginInfo := swag.FormatBool(qrWithPluginInfo)
		if qWithPluginInfo != "" {
			if err := r.SetQueryParam("withPluginInfo", qWithPluginInfo); err != nil {
				return err
			}
		}

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
