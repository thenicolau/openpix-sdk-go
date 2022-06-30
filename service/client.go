package service

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
)

type doFunc func(request *http.Request) (*http.Response, error)

type client struct {
	appID string
	hTTPClient *http.Client
	do	  doFunc
}

func Connect(appID string) *client {
	return &client{
		appID: appID,
		hTTPClient: http.DefaultClient,
	}
}

func (c *client) callApi(ctx context.Context, r *request) (data []byte, err error) {
	r.Url = "https://api.openpix.com.br/api"
	req, err := http.NewRequest(r.Method, fmt.Sprintf("%v%v", r.Url, r.Endpoint), r.Body)
	if err != nil {
		return []byte{}, err
	}

	req = req.WithContext(ctx)
	
	req.Header = r.Header

	f := c.do

	if f == nil {
		f = c.hTTPClient.Do
	}

	res, err := f(req)
	if err != nil {
		return []byte{}, err
	}

	defer func() {
		cerr := res.Body.Close()
		if err == nil && cerr != nil {
			err = cerr
		}
	}()

	data, err = ioutil.ReadAll(res.Body)
	fmt.Println(string(data))
	if err != nil {
		return []byte{}, err
	}


	


	return data, nil

}

func (c *client) GetOneCharge() *getOneCharge {
	return &getOneCharge{client: c}
}