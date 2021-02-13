// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"cdr.dev/coder-cli/sdk2/models"
)

// RegenUserSshkeyReader is a Reader for the RegenUserSshkey structure.
type RegenUserSshkeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RegenUserSshkeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRegenUserSshkeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRegenUserSshkeyOK creates a RegenUserSshkeyOK with default headers values
func NewRegenUserSshkeyOK() *RegenUserSshkeyOK {
	return &RegenUserSshkeyOK{}
}

/* RegenUserSshkeyOK describes a response with status code 200, with default header values.

OK
*/
type RegenUserSshkeyOK struct {
	Payload *models.SSHKeyPair
}

func (o *RegenUserSshkeyOK) Error() string {
	return fmt.Sprintf("[POST /v0/users/{id}/regen-ssh][%d] regenUserSshkeyOK  %+v", 200, o.Payload)
}
func (o *RegenUserSshkeyOK) GetPayload() *models.SSHKeyPair {
	return o.Payload
}

func (o *RegenUserSshkeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SSHKeyPair)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
