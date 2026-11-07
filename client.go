package aliyun

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/andywang2011/go-httplib"

	uuid "github.com/satori/go.uuid"
)

// jsonMarshalRequest may be replaced in tests to cover marshal error paths.
var jsonMarshalRequest = json.Marshal

// Client implements API.
type Client struct {
	cfg        Config
	httpClient *http.Client
}

// NewClient builds a Green client from cfg.
func NewClient(cfg Config) (API, error) {
	if cfg.Host == "" {
		cfg.Host = DefaultGreenHost
	}
	if cfg.Timeout <= 0 {
		cfg.Timeout = 10 * time.Second
	}
	if err := cfg.Validate(); err != nil {
		return nil, err
	}
	return &Client{
		cfg:        cfg,
		httpClient: httplib.NewClient(httplib.WithTimeout(cfg.Timeout)),
	}, nil
}

// NewGreenAPI is Wire-friendly: it returns the validation error instead of a nil API
// when cfg is invalid.
func NewGreenAPI(cfg Config) (API, error) {
	return NewClient(cfg)
}

func (c *Client) CheckText(ctx context.Context, content string) (*GreenResp, error) {
	path := "/green/text/scan"
	task := TextTask{
		DataId:  uuid.NewV4().String(),
		Content: content,
	}
	data := ScanRequest{
		Scenes: []string{"antispam"},
		Tasks:  []TextTask{task},
	}
	return c.doScan(ctx, path, data)
}

func (c *Client) CheckImage(ctx context.Context, imageURL string) (*GreenResp, error) {
	path := "/green/image/scan"
	task := ImageTask{
		DataId: uuid.NewV4().String(),
		Url:    imageURL,
	}
	data := ScanRequest{
		Scenes: []string{"antispam"},
		Tasks:  []ImageTask{task},
	}
	return c.doScan(ctx, path, data)
}

func (c *Client) doScan(ctx context.Context, path string, requestData ScanRequest) (*GreenResp, error) {
	body, err := c.doRequest(ctx, path, requestData)
	if err != nil {
		return nil, err
	}
	var greenResp GreenResp
	if err = json.Unmarshal(body, &greenResp); err != nil {
		return nil, fmt.Errorf("aliyun green unmarshal: %w", err)
	}
	return &greenResp, nil
}

func (c *Client) doRequest(ctx context.Context, path string, requestData ScanRequest) ([]byte, error) {
	reqDataByte, err := jsonMarshalRequest(requestData)
	if err != nil {
		return nil, fmt.Errorf("aliyun green marshal: %w", err)
	}
	bodyStr := string(reqDataByte)

	req, err := http.NewRequestWithContext(ctx, greenMethod, c.cfg.Host+path, strings.NewReader(bodyStr))
	if err != nil {
		return nil, err
	}
	for k, v := range buildRequestHeader(bodyStr, path, c.cfg.AccessKeyID, c.cfg.AccessKeySecret) {
		req.Header.Set(k, v)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return nil, &GreenError{
			HTTPStatus: resp.StatusCode,
			RawBody:    string(body),
			Msg:        string(body),
		}
	}
	return body, nil
}
