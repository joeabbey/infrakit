package tags

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/spiegela/gorackhd/models"
)

// GetTagReader is a Reader for the GetTag structure.
type GetTagReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTagReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetTagOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetTagNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetTagDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetTagOK creates a GetTagOK with default headers values
func NewGetTagOK() *GetTagOK {
	return &GetTagOK{}
}

/*GetTagOK handles this case with default header values.

Successfully retrieved information about the specified tag
*/
type GetTagOK struct {
	Payload GetTagOKBody
}

func (o *GetTagOK) Error() string {
	return fmt.Sprintf("[GET /tags/{tagName}][%d] getTagOK  %+v", 200, o.Payload)
}

func (o *GetTagOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTagNotFound creates a GetTagNotFound with default headers values
func NewGetTagNotFound() *GetTagNotFound {
	return &GetTagNotFound{}
}

/*GetTagNotFound handles this case with default header values.

The tag name identifier was not found
*/
type GetTagNotFound struct {
	Payload *models.Error
}

func (o *GetTagNotFound) Error() string {
	return fmt.Sprintf("[GET /tags/{tagName}][%d] getTagNotFound  %+v", 404, o.Payload)
}

func (o *GetTagNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTagDefault creates a GetTagDefault with default headers values
func NewGetTagDefault(code int) *GetTagDefault {
	return &GetTagDefault{
		_statusCode: code,
	}
}

/*GetTagDefault handles this case with default header values.

Unexpected error
*/
type GetTagDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get tag default response
func (o *GetTagDefault) Code() int {
	return o._statusCode
}

func (o *GetTagDefault) Error() string {
	return fmt.Sprintf("[GET /tags/{tagName}][%d] getTag default  %+v", o._statusCode, o.Payload)
}

func (o *GetTagDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetTagOKBody get tag o k body
swagger:model GetTagOKBody
*/
type GetTagOKBody interface{}
