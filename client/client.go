package client

type TenbounceClient struct {
	baseURL string
	cookie  string
}

func NewTenbounceClient(baseURL, cookie string) (TenbounceClient, error) {
	var tenbounceClient = TenbounceClient{
		baseURL: baseURL,
		cookie:  cookie,
	}

	return tenbounceClient, nil
}

