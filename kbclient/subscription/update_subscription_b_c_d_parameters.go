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

	kbmodel "github.com/killbill/kbcli/kbmodel"
)

// NewUpdateSubscriptionBCDParams creates a new UpdateSubscriptionBCDParams object
// with the default values initialized.
func NewUpdateSubscriptionBCDParams() *UpdateSubscriptionBCDParams {
	var (
		forceNewBcdWithPastEffectiveDateDefault = bool(false)
	)
	return &UpdateSubscriptionBCDParams{
		ForceNewBcdWithPastEffectiveDate: &forceNewBcdWithPastEffectiveDateDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateSubscriptionBCDParamsWithTimeout creates a new UpdateSubscriptionBCDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateSubscriptionBCDParamsWithTimeout(timeout time.Duration) *UpdateSubscriptionBCDParams {
	var (
		forceNewBcdWithPastEffectiveDateDefault = bool(false)
	)
	return &UpdateSubscriptionBCDParams{
		ForceNewBcdWithPastEffectiveDate: &forceNewBcdWithPastEffectiveDateDefault,

		timeout: timeout,
	}
}

// NewUpdateSubscriptionBCDParamsWithContext creates a new UpdateSubscriptionBCDParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateSubscriptionBCDParamsWithContext(ctx context.Context) *UpdateSubscriptionBCDParams {
	var (
		forceNewBcdWithPastEffectiveDateDefault = bool(false)
	)
	return &UpdateSubscriptionBCDParams{
		ForceNewBcdWithPastEffectiveDate: &forceNewBcdWithPastEffectiveDateDefault,

		Context: ctx,
	}
}

// NewUpdateSubscriptionBCDParamsWithHTTPClient creates a new UpdateSubscriptionBCDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateSubscriptionBCDParamsWithHTTPClient(client *http.Client) *UpdateSubscriptionBCDParams {
	var (
		forceNewBcdWithPastEffectiveDateDefault = bool(false)
	)
	return &UpdateSubscriptionBCDParams{
		ForceNewBcdWithPastEffectiveDate: &forceNewBcdWithPastEffectiveDateDefault,
		HTTPClient:                       client,
	}
}

/*UpdateSubscriptionBCDParams contains all the parameters to send to the API endpoint
for the update subscription b c d operation typically these are written to a http.Request
*/
type UpdateSubscriptionBCDParams struct {

	/*XKillbillComment*/
	XKillbillComment *string
	/*XKillbillCreatedBy*/
	XKillbillCreatedBy string
	/*XKillbillReason*/
	XKillbillReason *string
	/*Body*/
	Body *kbmodel.Subscription
	/*EffectiveFromDate*/
	EffectiveFromDate *strfmt.Date
	/*ForceNewBcdWithPastEffectiveDate*/
	ForceNewBcdWithPastEffectiveDate *bool
	/*SubscriptionID*/
	SubscriptionID strfmt.UUID

	WithStackTrace        *bool // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the update subscription b c d params
func (o *UpdateSubscriptionBCDParams) WithTimeout(timeout time.Duration) *UpdateSubscriptionBCDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update subscription b c d params
func (o *UpdateSubscriptionBCDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update subscription b c d params
func (o *UpdateSubscriptionBCDParams) WithContext(ctx context.Context) *UpdateSubscriptionBCDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update subscription b c d params
func (o *UpdateSubscriptionBCDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update subscription b c d params
func (o *UpdateSubscriptionBCDParams) WithHTTPClient(client *http.Client) *UpdateSubscriptionBCDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update subscription b c d params
func (o *UpdateSubscriptionBCDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the update subscription b c d params
func (o *UpdateSubscriptionBCDParams) WithXKillbillComment(xKillbillComment *string) *UpdateSubscriptionBCDParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the update subscription b c d params
func (o *UpdateSubscriptionBCDParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the update subscription b c d params
func (o *UpdateSubscriptionBCDParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *UpdateSubscriptionBCDParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the update subscription b c d params
func (o *UpdateSubscriptionBCDParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the update subscription b c d params
func (o *UpdateSubscriptionBCDParams) WithXKillbillReason(xKillbillReason *string) *UpdateSubscriptionBCDParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the update subscription b c d params
func (o *UpdateSubscriptionBCDParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithBody adds the body to the update subscription b c d params
func (o *UpdateSubscriptionBCDParams) WithBody(body *kbmodel.Subscription) *UpdateSubscriptionBCDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update subscription b c d params
func (o *UpdateSubscriptionBCDParams) SetBody(body *kbmodel.Subscription) {
	o.Body = body
}

// WithEffectiveFromDate adds the effectiveFromDate to the update subscription b c d params
func (o *UpdateSubscriptionBCDParams) WithEffectiveFromDate(effectiveFromDate *strfmt.Date) *UpdateSubscriptionBCDParams {
	o.SetEffectiveFromDate(effectiveFromDate)
	return o
}

// SetEffectiveFromDate adds the effectiveFromDate to the update subscription b c d params
func (o *UpdateSubscriptionBCDParams) SetEffectiveFromDate(effectiveFromDate *strfmt.Date) {
	o.EffectiveFromDate = effectiveFromDate
}

// WithForceNewBcdWithPastEffectiveDate adds the forceNewBcdWithPastEffectiveDate to the update subscription b c d params
func (o *UpdateSubscriptionBCDParams) WithForceNewBcdWithPastEffectiveDate(forceNewBcdWithPastEffectiveDate *bool) *UpdateSubscriptionBCDParams {
	o.SetForceNewBcdWithPastEffectiveDate(forceNewBcdWithPastEffectiveDate)
	return o
}

// SetForceNewBcdWithPastEffectiveDate adds the forceNewBcdWithPastEffectiveDate to the update subscription b c d params
func (o *UpdateSubscriptionBCDParams) SetForceNewBcdWithPastEffectiveDate(forceNewBcdWithPastEffectiveDate *bool) {
	o.ForceNewBcdWithPastEffectiveDate = forceNewBcdWithPastEffectiveDate
}

// WithSubscriptionID adds the subscriptionID to the update subscription b c d params
func (o *UpdateSubscriptionBCDParams) WithSubscriptionID(subscriptionID strfmt.UUID) *UpdateSubscriptionBCDParams {
	o.SetSubscriptionID(subscriptionID)
	return o
}

// SetSubscriptionID adds the subscriptionId to the update subscription b c d params
func (o *UpdateSubscriptionBCDParams) SetSubscriptionID(subscriptionID strfmt.UUID) {
	o.SubscriptionID = subscriptionID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateSubscriptionBCDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.EffectiveFromDate != nil {

		// query param effectiveFromDate
		var qrEffectiveFromDate strfmt.Date
		if o.EffectiveFromDate != nil {
			qrEffectiveFromDate = *o.EffectiveFromDate
		}
		qEffectiveFromDate := qrEffectiveFromDate.String()
		if qEffectiveFromDate != "" {
			if err := r.SetQueryParam("effectiveFromDate", qEffectiveFromDate); err != nil {
				return err
			}
		}

	}

	if o.ForceNewBcdWithPastEffectiveDate != nil {

		// query param forceNewBcdWithPastEffectiveDate
		var qrForceNewBcdWithPastEffectiveDate bool
		if o.ForceNewBcdWithPastEffectiveDate != nil {
			qrForceNewBcdWithPastEffectiveDate = *o.ForceNewBcdWithPastEffectiveDate
		}
		qForceNewBcdWithPastEffectiveDate := swag.FormatBool(qrForceNewBcdWithPastEffectiveDate)
		if qForceNewBcdWithPastEffectiveDate != "" {
			if err := r.SetQueryParam("forceNewBcdWithPastEffectiveDate", qForceNewBcdWithPastEffectiveDate); err != nil {
				return err
			}
		}

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