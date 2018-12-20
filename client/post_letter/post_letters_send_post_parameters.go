// Code generated by go-swagger; DO NOT EDIT.

package post_letter

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

	models "github.com/danikarik/clicksend-go/models"
)

// NewPostLettersSendPostParams creates a new PostLettersSendPostParams object
// with the default values initialized.
func NewPostLettersSendPostParams() *PostLettersSendPostParams {
	var ()
	return &PostLettersSendPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLettersSendPostParamsWithTimeout creates a new PostLettersSendPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLettersSendPostParamsWithTimeout(timeout time.Duration) *PostLettersSendPostParams {
	var ()
	return &PostLettersSendPostParams{

		timeout: timeout,
	}
}

// NewPostLettersSendPostParamsWithContext creates a new PostLettersSendPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLettersSendPostParamsWithContext(ctx context.Context) *PostLettersSendPostParams {
	var ()
	return &PostLettersSendPostParams{

		Context: ctx,
	}
}

// NewPostLettersSendPostParamsWithHTTPClient creates a new PostLettersSendPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLettersSendPostParamsWithHTTPClient(client *http.Client) *PostLettersSendPostParams {
	var ()
	return &PostLettersSendPostParams{
		HTTPClient: client,
	}
}

/*PostLettersSendPostParams contains all the parameters to send to the API endpoint
for the post letters send post operation typically these are written to a http.Request
*/
type PostLettersSendPostParams struct {

	/*PostLetter
	  PostLetter model

	*/
	PostLetter models.PostLetter

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post letters send post params
func (o *PostLettersSendPostParams) WithTimeout(timeout time.Duration) *PostLettersSendPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post letters send post params
func (o *PostLettersSendPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post letters send post params
func (o *PostLettersSendPostParams) WithContext(ctx context.Context) *PostLettersSendPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post letters send post params
func (o *PostLettersSendPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post letters send post params
func (o *PostLettersSendPostParams) WithHTTPClient(client *http.Client) *PostLettersSendPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post letters send post params
func (o *PostLettersSendPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPostLetter adds the postLetter to the post letters send post params
func (o *PostLettersSendPostParams) WithPostLetter(postLetter models.PostLetter) *PostLettersSendPostParams {
	o.SetPostLetter(postLetter)
	return o
}

// SetPostLetter adds the postLetter to the post letters send post params
func (o *PostLettersSendPostParams) SetPostLetter(postLetter models.PostLetter) {
	o.PostLetter = postLetter
}

// WriteToRequest writes these params to a swagger request
func (o *PostLettersSendPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.PostLetter); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}