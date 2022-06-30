package service

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type _params map[string]interface{}

type request struct {
	Method string
	Endpoint string
	Query url.Values
	Form url.Values
	Header http.Header
	Body io.Reader
	Url string
}

func (r *request) SetParam(key string, value interface{}) *request {
	if r.Query == nil {
		r.Query = url.Values{}
	}

	r.Query.Set(key, fmt.Sprintf("%v", value))
	return r
}

func (r *request) SetParams(p _params) *request {
	for key, value := range p{
		r.SetParam(key, value)
	}

	return r
}

