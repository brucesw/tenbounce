package client

import "net/http"

type TenbounceClient struct {
	baseURL string
	cookie  string

	client *http.Client
}

func NewTenbounceClient(baseURL, cookie string) (TenbounceClient, error) {
	var client = &http.Client{}

	var tenbounceClient = TenbounceClient{
		baseURL: baseURL,
		cookie:  cookie,

		client: client,
	}

	return tenbounceClient, nil
}
