// This file was auto-generated by Fern from our API Definition.

package dataset

import (
	context "context"
	fmt "fmt"
	coherego "github.com/cohere-ai/cohere-go"
	core "github.com/cohere-ai/cohere-go/core"
	http "net/http"
	url "net/url"
	time "time"
)

type Client struct {
	baseURL string
	caller  *core.Caller
	header  http.Header
}

func NewClient(opts ...core.ClientOption) *Client {
	options := core.NewClientOptions()
	for _, opt := range opts {
		opt(options)
	}
	return &Client{
		baseURL: options.BaseURL,
		caller:  core.NewCaller(options.HTTPClient),
		header:  options.ToHeader(),
	}
}

func (c *Client) Get(ctx context.Context, request *coherego.DatasetGetRequest) (*coherego.DatasetGetResponse, error) {
	baseURL := "https://api.cohere.ai"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "v1/dataset"

	queryParams := make(url.Values)
	if request.DatasetType != nil {
		queryParams.Add("datasetType", fmt.Sprintf("%v", *request.DatasetType))
	}
	if request.Before != nil {
		queryParams.Add("before", fmt.Sprintf("%v", request.Before.Format(time.RFC3339)))
	}
	if request.After != nil {
		queryParams.Add("after", fmt.Sprintf("%v", request.After.Format(time.RFC3339)))
	}
	if request.Limit != nil {
		queryParams.Add("limit", fmt.Sprintf("%v", *request.Limit))
	}
	if request.Offset != nil {
		queryParams.Add("offset", fmt.Sprintf("%v", *request.Offset))
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	var response *coherego.DatasetGetResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:      endpointURL,
			Method:   http.MethodGet,
			Headers:  c.header,
			Response: &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) UsageGet(ctx context.Context) (*coherego.DatasetUsageGetResponse, error) {
	baseURL := "https://api.cohere.ai"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "v1/dataset/usage"

	var response *coherego.DatasetUsageGetResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:      endpointURL,
			Method:   http.MethodGet,
			Headers:  c.header,
			Response: &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) IdGet(ctx context.Context, id string) (*coherego.DatasetIdGetResponse, error) {
	baseURL := "https://api.cohere.ai"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"v1/dataset/%v", id)

	var response *coherego.DatasetIdGetResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:      endpointURL,
			Method:   http.MethodGet,
			Headers:  c.header,
			Response: &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) IdDelete(ctx context.Context, id string) (map[string]interface{}, error) {
	baseURL := "https://api.cohere.ai"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"v1/dataset/%v", id)

	var response map[string]interface{}
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:      endpointURL,
			Method:   http.MethodDelete,
			Headers:  c.header,
			Response: &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}