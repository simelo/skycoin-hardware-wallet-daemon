// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/skycoin/hardware-wallet-daemon/src/models"
)

// PostIntermediateButtonReader is a Reader for the PostIntermediateButton structure.
type PostIntermediateButtonReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostIntermediateButtonReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostIntermediateButtonOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPostIntermediateButtonDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostIntermediateButtonOK creates a PostIntermediateButtonOK with default headers values
func NewPostIntermediateButtonOK() *PostIntermediateButtonOK {
	return &PostIntermediateButtonOK{}
}

/*PostIntermediateButtonOK handles this case with default header values.

success
*/
type PostIntermediateButtonOK struct {
	Payload *models.HttpsuccessResponse
}

func (o *PostIntermediateButtonOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HttpsuccessResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntermediateButtonDefault creates a PostIntermediateButtonDefault with default headers values
func NewPostIntermediateButtonDefault(code int) *PostIntermediateButtonDefault {
	return &PostIntermediateButtonDefault{
		_statusCode: code,
	}
}

/*PostIntermediateButtonDefault handles this case with default header values.

error
*/
type PostIntermediateButtonDefault struct {
	_statusCode int

	Payload *models.HTTPErrorResponse
}

// Code gets the status code for the post intermediate button default response
func (o *PostIntermediateButtonDefault) Code() int {
	return o._statusCode
}

func (o *PostIntermediateButtonDefault) Error() string {
	return o.Payload.Error.Message
}

func (o *PostIntermediateButtonDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
