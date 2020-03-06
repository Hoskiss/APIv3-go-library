// Code generated by go-swagger; DO NOT EDIT.

package reseller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/sendinblue/APIv3-go-library/models"
)

// GetSsoTokenReader is a Reader for the GetSsoToken structure.
type GetSsoTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSsoTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSsoTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetSsoTokenBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetSsoTokenForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetSsoTokenNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSsoTokenOK creates a GetSsoTokenOK with default headers values
func NewGetSsoTokenOK() *GetSsoTokenOK {
	return &GetSsoTokenOK{}
}

/*GetSsoTokenOK handles this case with default header values.

Session token
*/
type GetSsoTokenOK struct {
	Payload *models.GetSsoToken
}

func (o *GetSsoTokenOK) Error() string {
	return fmt.Sprintf("[GET /reseller/children/{childAuthKey}/auth][%d] getSsoTokenOK  %+v", 200, o.Payload)
}

func (o *GetSsoTokenOK) GetPayload() *models.GetSsoToken {
	return o.Payload
}

func (o *GetSsoTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetSsoToken)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSsoTokenBadRequest creates a GetSsoTokenBadRequest with default headers values
func NewGetSsoTokenBadRequest() *GetSsoTokenBadRequest {
	return &GetSsoTokenBadRequest{}
}

/*GetSsoTokenBadRequest handles this case with default header values.

bad request
*/
type GetSsoTokenBadRequest struct {
	Payload *models.ErrorModel
}

func (o *GetSsoTokenBadRequest) Error() string {
	return fmt.Sprintf("[GET /reseller/children/{childAuthKey}/auth][%d] getSsoTokenBadRequest  %+v", 400, o.Payload)
}

func (o *GetSsoTokenBadRequest) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *GetSsoTokenBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSsoTokenForbidden creates a GetSsoTokenForbidden with default headers values
func NewGetSsoTokenForbidden() *GetSsoTokenForbidden {
	return &GetSsoTokenForbidden{}
}

/*GetSsoTokenForbidden handles this case with default header values.

Current account is not a reseller
*/
type GetSsoTokenForbidden struct {
	Payload *models.ErrorModel
}

func (o *GetSsoTokenForbidden) Error() string {
	return fmt.Sprintf("[GET /reseller/children/{childAuthKey}/auth][%d] getSsoTokenForbidden  %+v", 403, o.Payload)
}

func (o *GetSsoTokenForbidden) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *GetSsoTokenForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSsoTokenNotFound creates a GetSsoTokenNotFound with default headers values
func NewGetSsoTokenNotFound() *GetSsoTokenNotFound {
	return &GetSsoTokenNotFound{}
}

/*GetSsoTokenNotFound handles this case with default header values.

Child auth key not found
*/
type GetSsoTokenNotFound struct {
	Payload *models.ErrorModel
}

func (o *GetSsoTokenNotFound) Error() string {
	return fmt.Sprintf("[GET /reseller/children/{childAuthKey}/auth][%d] getSsoTokenNotFound  %+v", 404, o.Payload)
}

func (o *GetSsoTokenNotFound) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *GetSsoTokenNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}