package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"tenbounce/api"
)

// TODO(bruce): Error handling
func (c TenbounceClient) ListPoints() (api.ListPointsResponse, error) {
	var listPointsResponse api.ListPointsResponse

	var reqURL = fmt.Sprintf("%s/api/points", c.baseURL)
	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return api.ListPointsResponse{}, fmt.Errorf("new request: %w", err)
	}

	var cookie = http.Cookie{
		Name:  "TENBOUNCE_USER_ID",
		Value: c.cookie,
	}
	req.AddCookie(&cookie)

	var client = &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return api.ListPointsResponse{}, fmt.Errorf("client do: %w", err)
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return api.ListPointsResponse{}, nil
	}
	fmt.Println("resBody", string(resBody))

	if err = json.Unmarshal(resBody, &listPointsResponse); err != nil {
		return api.ListPointsResponse{}, fmt.Errorf("unmarshal list points response: %w", err)
	}

	return listPointsResponse, nil
}
