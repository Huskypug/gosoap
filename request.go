package gosoap

import (
	"fmt"
)

// Request Soap Request
type Request struct {
	Method   string
	Params   SoapParams
	HttpHead map[string]string
}

func NewRequest(m string, p SoapParams) *Request {
	return &Request{
		Method: m,
		Params: p,
	}
}

func NewRequestWithHeader(m string, p SoapParams, h map[string]string) *Request {
	return &Request{
		Method:   m,
		Params:   p,
		HttpHead: h,
	}
}

// RequestStruct soap request interface
type RequestWithHeaderStruct interface {
	SoapBuildRequestWithHeader() *Request
}

// NewRequestByStruct create a new request using builder
func NewRequestWithHeaderByStruct(s RequestWithHeaderStruct) (*Request, error) {
	if s == nil {
		return nil, fmt.Errorf("'s' cannot be 'nil'")
	}

	return s.SoapBuildRequestWithHeader(), nil
}

// RequestStruct soap request interface
type RequestStruct interface {
	SoapBuildRequest() *Request
}

// NewRequestByStruct create a new request using builder
func NewRequestByStruct(s RequestStruct) (*Request, error) {
	if s == nil {
		return nil, fmt.Errorf("'s' cannot be 'nil'")
	}

	return s.SoapBuildRequest(), nil
}
