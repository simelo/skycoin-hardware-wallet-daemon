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

// PostGenerateMnemonicReader is a Reader for the PostGenerateMnemonic structure.
type PostGenerateMnemonicReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostGenerateMnemonicReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostGenerateMnemonicOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPostGenerateMnemonicDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostGenerateMnemonicOK creates a PostGenerateMnemonicOK with default headers values
func NewPostGenerateMnemonicOK() *PostGenerateMnemonicOK {
	return &PostGenerateMnemonicOK{}
}

/*PostGenerateMnemonicOK handles this case with default header values.

successful operation
*/
type PostGenerateMnemonicOK struct {
	Payload *models.HttpsuccessResponse
}

func (o *PostGenerateMnemonicOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HttpsuccessResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGenerateMnemonicDefault creates a PostGenerateMnemonicDefault with default headers values
func NewPostGenerateMnemonicDefault(code int) *PostGenerateMnemonicDefault {
	return &PostGenerateMnemonicDefault{
		_statusCode: code,
	}
}

/*PostGenerateMnemonicDefault handles this case with default header values.

error
*/
type PostGenerateMnemonicDefault struct {
	_statusCode int

	Payload *models.HTTPErrorResponse
}

// Code gets the status code for the post generate mnemonic default response
func (o *PostGenerateMnemonicDefault) Code() int {
	return o._statusCode
}

func (o *PostGenerateMnemonicDefault) Error() string {
	return o.Payload.Error.Message
}

func (o *PostGenerateMnemonicDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
