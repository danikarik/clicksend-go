// Code generated by go-swagger; DO NOT EDIT.

package post_postcard

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

	models "github.com/danikarik/clicksend-go/models"
)

// NewPostPostcardsPricePostParams creates a new PostPostcardsPricePostParams object
// with the default values initialized.
func NewPostPostcardsPricePostParams() *PostPostcardsPricePostParams {
	var ()
	return &PostPostcardsPricePostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostPostcardsPricePostParamsWithTimeout creates a new PostPostcardsPricePostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostPostcardsPricePostParamsWithTimeout(timeout time.Duration) *PostPostcardsPricePostParams {
	var ()
	return &PostPostcardsPricePostParams{

		timeout: timeout,
	}
}

// NewPostPostcardsPricePostParamsWithContext creates a new PostPostcardsPricePostParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostPostcardsPricePostParamsWithContext(ctx context.Context) *PostPostcardsPricePostParams {
	var ()
	return &PostPostcardsPricePostParams{

		Context: ctx,
	}
}

// NewPostPostcardsPricePostParamsWithHTTPClient creates a new PostPostcardsPricePostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostPostcardsPricePostParamsWithHTTPClient(client *http.Client) *PostPostcardsPricePostParams {
	var ()
	return &PostPostcardsPricePostParams{
		HTTPClient: client,
	}
}

/*PostPostcardsPricePostParams contains all the parameters to send to the API endpoint
for the post postcards price post operation typically these are written to a http.Request
*/
type PostPostcardsPricePostParams struct {

	/*PostPostcards
	  PostPostcard model

	*/
	PostPostcards models.PostPostcard

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post postcards price post params
func (o *PostPostcardsPricePostParams) WithTimeout(timeout time.Duration) *PostPostcardsPricePostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post postcards price post params
func (o *PostPostcardsPricePostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post postcards price post params
func (o *PostPostcardsPricePostParams) WithContext(ctx context.Context) *PostPostcardsPricePostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post postcards price post params
func (o *PostPostcardsPricePostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post postcards price post params
func (o *PostPostcardsPricePostParams) WithHTTPClient(client *http.Client) *PostPostcardsPricePostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post postcards price post params
func (o *PostPostcardsPricePostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPostPostcards adds the postPostcards to the post postcards price post params
func (o *PostPostcardsPricePostParams) WithPostPostcards(postPostcards models.PostPostcard) *PostPostcardsPricePostParams {
	o.SetPostPostcards(postPostcards)
	return o
}

// SetPostPostcards adds the postPostcards to the post postcards price post params
func (o *PostPostcardsPricePostParams) SetPostPostcards(postPostcards models.PostPostcard) {
	o.PostPostcards = postPostcards
}

// WriteToRequest writes these params to a swagger request
func (o *PostPostcardsPricePostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.PostPostcards); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
