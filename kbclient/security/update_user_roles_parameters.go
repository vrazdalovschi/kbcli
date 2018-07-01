// Code generated by go-swagger; DO NOT EDIT.

package security

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

	kbmodel "github.com/killbill/kbcli/kbmodel"
)

// NewUpdateUserRolesParams creates a new UpdateUserRolesParams object
// with the default values initialized.
func NewUpdateUserRolesParams() *UpdateUserRolesParams {
	var ()
	return &UpdateUserRolesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateUserRolesParamsWithTimeout creates a new UpdateUserRolesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateUserRolesParamsWithTimeout(timeout time.Duration) *UpdateUserRolesParams {
	var ()
	return &UpdateUserRolesParams{

		timeout: timeout,
	}
}

// NewUpdateUserRolesParamsWithContext creates a new UpdateUserRolesParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateUserRolesParamsWithContext(ctx context.Context) *UpdateUserRolesParams {
	var ()
	return &UpdateUserRolesParams{

		Context: ctx,
	}
}

// NewUpdateUserRolesParamsWithHTTPClient creates a new UpdateUserRolesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateUserRolesParamsWithHTTPClient(client *http.Client) *UpdateUserRolesParams {
	var ()
	return &UpdateUserRolesParams{
		HTTPClient: client,
	}
}

/*UpdateUserRolesParams contains all the parameters to send to the API endpoint
for the update user roles operation typically these are written to a http.Request
*/
type UpdateUserRolesParams struct {

	/*XKillbillComment*/
	XKillbillComment *string
	/*XKillbillCreatedBy*/
	XKillbillCreatedBy string
	/*XKillbillReason*/
	XKillbillReason *string
	/*Body*/
	Body *kbmodel.UserRoles
	/*Username*/
	Username string

	WithStackTrace        *bool // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the update user roles params
func (o *UpdateUserRolesParams) WithTimeout(timeout time.Duration) *UpdateUserRolesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update user roles params
func (o *UpdateUserRolesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update user roles params
func (o *UpdateUserRolesParams) WithContext(ctx context.Context) *UpdateUserRolesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update user roles params
func (o *UpdateUserRolesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update user roles params
func (o *UpdateUserRolesParams) WithHTTPClient(client *http.Client) *UpdateUserRolesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update user roles params
func (o *UpdateUserRolesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the update user roles params
func (o *UpdateUserRolesParams) WithXKillbillComment(xKillbillComment *string) *UpdateUserRolesParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the update user roles params
func (o *UpdateUserRolesParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the update user roles params
func (o *UpdateUserRolesParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *UpdateUserRolesParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the update user roles params
func (o *UpdateUserRolesParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the update user roles params
func (o *UpdateUserRolesParams) WithXKillbillReason(xKillbillReason *string) *UpdateUserRolesParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the update user roles params
func (o *UpdateUserRolesParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithBody adds the body to the update user roles params
func (o *UpdateUserRolesParams) WithBody(body *kbmodel.UserRoles) *UpdateUserRolesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update user roles params
func (o *UpdateUserRolesParams) SetBody(body *kbmodel.UserRoles) {
	o.Body = body
}

// WithUsername adds the username to the update user roles params
func (o *UpdateUserRolesParams) WithUsername(username string) *UpdateUserRolesParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the update user roles params
func (o *UpdateUserRolesParams) SetUsername(username string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateUserRolesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param username
	if err := r.SetPathParam("username", o.Username); err != nil {
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
