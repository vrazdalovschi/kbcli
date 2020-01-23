// Code generated by go-swagger; DO NOT EDIT.

package subscription

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

// NewGetSubscriptionAuditLogsWithHistoryParams creates a new GetSubscriptionAuditLogsWithHistoryParams object
// with the default values initialized.
func NewGetSubscriptionAuditLogsWithHistoryParams() *GetSubscriptionAuditLogsWithHistoryParams {
	var ()
	return &GetSubscriptionAuditLogsWithHistoryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSubscriptionAuditLogsWithHistoryParamsWithTimeout creates a new GetSubscriptionAuditLogsWithHistoryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSubscriptionAuditLogsWithHistoryParamsWithTimeout(timeout time.Duration) *GetSubscriptionAuditLogsWithHistoryParams {
	var ()
	return &GetSubscriptionAuditLogsWithHistoryParams{

		timeout: timeout,
	}
}

// NewGetSubscriptionAuditLogsWithHistoryParamsWithContext creates a new GetSubscriptionAuditLogsWithHistoryParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSubscriptionAuditLogsWithHistoryParamsWithContext(ctx context.Context) *GetSubscriptionAuditLogsWithHistoryParams {
	var ()
	return &GetSubscriptionAuditLogsWithHistoryParams{

		Context: ctx,
	}
}

// NewGetSubscriptionAuditLogsWithHistoryParamsWithHTTPClient creates a new GetSubscriptionAuditLogsWithHistoryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSubscriptionAuditLogsWithHistoryParamsWithHTTPClient(client *http.Client) *GetSubscriptionAuditLogsWithHistoryParams {
	var ()
	return &GetSubscriptionAuditLogsWithHistoryParams{
		HTTPClient: client,
	}
}

/*GetSubscriptionAuditLogsWithHistoryParams contains all the parameters to send to the API endpoint
for the get subscription audit logs with history operation typically these are written to a http.Request
*/
type GetSubscriptionAuditLogsWithHistoryParams struct {

	/*SubscriptionID*/
	SubscriptionID strfmt.UUID

	WithStackTrace        *bool // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the get subscription audit logs with history params
func (o *GetSubscriptionAuditLogsWithHistoryParams) WithTimeout(timeout time.Duration) *GetSubscriptionAuditLogsWithHistoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get subscription audit logs with history params
func (o *GetSubscriptionAuditLogsWithHistoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get subscription audit logs with history params
func (o *GetSubscriptionAuditLogsWithHistoryParams) WithContext(ctx context.Context) *GetSubscriptionAuditLogsWithHistoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get subscription audit logs with history params
func (o *GetSubscriptionAuditLogsWithHistoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get subscription audit logs with history params
func (o *GetSubscriptionAuditLogsWithHistoryParams) WithHTTPClient(client *http.Client) *GetSubscriptionAuditLogsWithHistoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get subscription audit logs with history params
func (o *GetSubscriptionAuditLogsWithHistoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSubscriptionID adds the subscriptionID to the get subscription audit logs with history params
func (o *GetSubscriptionAuditLogsWithHistoryParams) WithSubscriptionID(subscriptionID strfmt.UUID) *GetSubscriptionAuditLogsWithHistoryParams {
	o.SetSubscriptionID(subscriptionID)
	return o
}

// SetSubscriptionID adds the subscriptionId to the get subscription audit logs with history params
func (o *GetSubscriptionAuditLogsWithHistoryParams) SetSubscriptionID(subscriptionID strfmt.UUID) {
	o.SubscriptionID = subscriptionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetSubscriptionAuditLogsWithHistoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
