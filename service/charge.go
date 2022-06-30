package service

import (
	"context"
	"encoding/json"
	"time"
)

type getOneCharge struct {
	client *client
	params params

}
func (c *getOneCharge) DateStart(t time.Time) *getOneCharge {
	c.params.start = &t
	return c
}

func (c *getOneCharge) DateEnd(t time.Time) *getOneCharge{
	c.params.end = &t
	return c
}

func (c *getOneCharge) Status(s string) *getOneCharge {
	c.params.status = &s
	return c
}

func (c *getOneCharge) Do(ctx context.Context) (response *charge, err error) {

	request := &request{
		Method: "GET",
		Endpoint: "/openpix/v1/charge" ,
	}
	
	if c.params.start != nil {
		request.SetParam("start", *c.params.start)
	}
	if c.params.end != nil {
		request.SetParam("end", *c.params.end)
	}
	if c.params.status != nil {
		request.SetParam("status", *c.params.status)
	}

	data, err := c.client.callApi(ctx, request)
	if err != nil {
		return &charge{}, err
	}

	response = &charge{}
	
	err = json.Unmarshal(data, &response)

	if err != nil {
		return &charge{}, err
	}

	return response, err
}