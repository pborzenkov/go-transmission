package transmission

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"
)

const (
	defaultRPCPath = "/transmission/rpc"

	headerCSRF = "X-Transmission-Session-Id"
)

// Client is a Transmission RPC client.
type Client struct {
	config

	url string

	mu        sync.Mutex
	sessionID string
}

// New returns new instance of a Client.
func New(host string, opts ...Option) (*Client, error) {
	uri, err := url.Parse(host)
	if err != nil {
		return nil, err
	}

	if uri.Path == "" {
		uri.Path = defaultRPCPath
	}

	c := &Client{
		url: uri.String(),
	}
	for _, opt := range opts {
		opt.apply(&c.config)
	}
	if c.HTTPClient == nil {
		c.HTTPClient = http.DefaultClient
	}

	return c, nil
}

func (c *Client) getSessionID() string {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.sessionID
}

func (c *Client) setSessionID(id string) {
	c.mu.Lock()
	c.sessionID = id
	c.mu.Unlock()
}

type rpcRequest struct {
	Method    string      `json:"method"`
	Arguments interface{} `json:"arguments,omitempty"`
}

type rpcResponse struct {
	Result    string      `json:"result"`
	Arguments interface{} `json:"arguments"`
}

func (c *Client) newRequest(ctx context.Context, data []byte) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, "POST", c.url, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(headerCSRF, c.getSessionID())
	if c.Username != "" || c.Password != "" {
		req.SetBasicAuth(c.Username, c.Password)
	}
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}

	return req, nil
}

func (c *Client) callRPC(ctx context.Context, method string, args interface{}, reply interface{}) error {
	reqData := new(bytes.Buffer)
	if err := json.NewEncoder(reqData).Encode(&rpcRequest{
		Method:    method,
		Arguments: args,
	}); err != nil {
		return err
	}

	var resp *http.Response
	for i := 0; i < 2; i++ {
		req, err := c.newRequest(ctx, reqData.Bytes())
		if err != nil {
			return err
		}
		resp, err = c.HTTPClient.Do(req)
		if err != nil {
			return err
		}
		if resp.StatusCode/100 == 2 {
			defer resp.Body.Close()
			break
		}

		_, _ = io.Copy(ioutil.Discard, io.LimitReader(resp.Body, 4096))
		resp.Body.Close()
		if resp.StatusCode != http.StatusConflict {
			return fmt.Errorf("transmission: HTTP request failed (%s)", http.StatusText(resp.StatusCode))
		}
		c.setSessionID(resp.Header.Get(headerCSRF))
	}
	if resp.StatusCode/100 != 2 {
		return fmt.Errorf("transmission: CSRF token not accepted")
	}

	response := &rpcResponse{
		Arguments: reply,
	}
	if err := json.NewDecoder(resp.Body).Decode(response); err != nil {
		return err
	}
	if response.Result != "success" {
		return fmt.Errorf("transmission: RPC call failed (%s)", response.Result)
	}

	return nil
}
