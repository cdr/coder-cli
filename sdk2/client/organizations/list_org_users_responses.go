// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"cdr.dev/coder-cli/sdk2/models"
)

// ListOrgUsersReader is a Reader for the ListOrgUsers structure.
type ListOrgUsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListOrgUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListOrgUsersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListOrgUsersOK creates a ListOrgUsersOK with default headers values
func NewListOrgUsersOK() *ListOrgUsersOK {
	return &ListOrgUsersOK{}
}

/* ListOrgUsersOK describes a response with status code 200, with default header values.

OK
*/
type ListOrgUsersOK struct {
	Payload []*models.OrganizationUser
}

func (o *ListOrgUsersOK) Error() string {
	return fmt.Sprintf("[GET /v0/orgs/{id}/members][%d] listOrgUsersOK  %+v", 200, o.Payload)
}
func (o *ListOrgUsersOK) GetPayload() []*models.OrganizationUser {
	return o.Payload
}

func (o *ListOrgUsersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
