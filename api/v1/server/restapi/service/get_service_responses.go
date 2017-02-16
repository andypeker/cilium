package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/cilium/api/v1/models"
)

// GetServiceOKCode is the HTTP code returned for type GetServiceOK
const GetServiceOKCode int = 200

/*GetServiceOK Success

swagger:response getServiceOK
*/
type GetServiceOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Service `json:"body,omitempty"`
}

// NewGetServiceOK creates GetServiceOK with default headers values
func NewGetServiceOK() *GetServiceOK {
	return &GetServiceOK{}
}

// WithPayload adds the payload to the get service o k response
func (o *GetServiceOK) WithPayload(payload []*models.Service) *GetServiceOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get service o k response
func (o *GetServiceOK) SetPayload(payload []*models.Service) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetServiceOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make([]*models.Service, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}