package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"tenbounce/api"
)

func (c TenbounceClient) ListPoints(ctx context.Context) (api.ListPointsResponse, error) {
	var listPointsResponse api.ListPointsResponse

	var reqURL = fmt.Sprintf("%s/api/points", c.baseURL)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL, nil)
	if err != nil {
		return api.ListPointsResponse{}, fmt.Errorf("new request: %w", err)
	}

	// TODO(bruce): Abstract this somehow once more routes are supported
	var cookie = http.Cookie{
		Name:  api.CookieName_UserID,
		Value: c.cookie,
	}
	req.AddCookie(&cookie)

	res, err := c.client.Do(req)
	if err != nil {
		return api.ListPointsResponse{}, fmt.Errorf("client do request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode >= 300 {
		body, _ := io.ReadAll(res.Body)
		return api.ListPointsResponse{}, fmt.Errorf("http error response code %d: %s", res.StatusCode, body)
	}

	if err := json.NewDecoder(res.Body).Decode(&listPointsResponse); err != nil {
		return api.ListPointsResponse{}, fmt.Errorf("decode list points response: %w", err)
	}

	return listPointsResponse, nil
}
